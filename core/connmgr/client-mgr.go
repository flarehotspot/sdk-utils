package connmgr

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/network"
	connmgr "github.com/flarehotspot/core/sdk/api/connmgr"
	sdknet "github.com/flarehotspot/core/sdk/api/network"
	slices "github.com/flarehotspot/core/sdk/utils/slices"
	sse "github.com/flarehotspot/core/sdk/utils/sse"
	"github.com/flarehotspot/core/utils/nftables"
)

const (
	EventConnected    string = "client:connected"
	EventDisconnected string = "client:disconnected"
)

type ClientMgr struct {
	mu       sync.RWMutex
	sessions []*RunningSession
}

func NewClientMgr(dtb *db.Database, mdls *models.Models) *ClientMgr {
	return &ClientMgr{}
}

func (cmgr *ClientMgr) ListenTraffic(trfk *network.TrafficMgr) {
	go func() {
		for data := range trfk.Listen() {
			go func(data *sdknet.TrafficData) {
				cmgr.mu.RLock()
				defer cmgr.mu.RUnlock()

				for _, s := range cmgr.sessions {
					s.UpdateData(data)
				}
			}(&data)
		}
	}()
}

func (cmgr *ClientMgr) ReloadSessions(ctx context.Context, iface string) error {
	errCh := make(chan error)
	go func() {
		cmgr.mu.RLock()
		defer cmgr.mu.RUnlock()

		for _, rs := range cmgr.sessions {
			lan := rs.Lan()

			if lan.Name() == iface {
				cs := rs.GetSession()
				err := cs.Reload(ctx)
				if err != nil {
					errCh <- err
					break
				}

				err = rs.Change(cs)
				if err != nil {
					errCh <- err
					break
				}
			}
		}

		errCh <- nil
	}()
	return <-errCh
}

func (cmgr *ClientMgr) StopSessions(ctx context.Context, iface string, reason string) {
	done := make(chan bool)
	go func() {
		cmgr.mu.Lock()
		defer cmgr.mu.Unlock()
		defer func() {
			done <- true
		}()

		for _, rs := range cmgr.sessions {
			err := nftables.Disconnect(rs.mac, reason)
			if err != nil {
				log.Println(err)
			}

			lan, err := network.FindByIp(rs.ip)
			if err != nil {
				log.Println(err)
			}

			if lan.Name() == iface {
				rs.Stop(context.Background())
			}
		}
	}()
	<-done
}

func (cmgr *ClientMgr) Connect(clnt connmgr.IClientDevice) error {
	errCh := make(chan error)

	go func() {
		if !clnt.HasSession(context.Background()) {
			errCh <- errors.New("No available sessions.")
			return
		}

		if !nftables.IsConnected(clnt.MacAddr()) {
			if err := nftables.Connect(clnt.IpAddr(), clnt.MacAddr()); err != nil {
				errCh <- err
				return
			}
		}

		go cmgr.loopSessions(clnt)

		data := map[string]any{"message": "You are now connected to internet."}
		cmgr.SocketEmit(clnt, "client:connected", data)
		errCh <- nil
	}()

	return <-errCh
}

func (cmgr *ClientMgr) Disconnect(clnt connmgr.IClientDevice, notify error) error {
	log.Println("Calling endsession()...")
	err := cmgr.endSession(clnt)
	if err != nil {
		notify = err
	}

	if notify != nil {
		data := map[string]any{"message": notify.Error()}
		cmgr.SocketEmit(clnt, EventDisconnected, data)
	}

	return err
}

func (cmgr *ClientMgr) IsConnected(clnt connmgr.IClientDevice) (connected bool) {
	return nftables.IsConnected(clnt.MacAddr())
}

func (cmgr *ClientMgr) CurrSession(clnt connmgr.IClientDevice) (cs connmgr.IClientSession, ok bool) {
	cmgr.mu.RLock()
	defer cmgr.mu.RUnlock()

	for _, rs := range cmgr.sessions {
		if rs.GetSession().DeviceId() == clnt.Id() {
			return rs.session, true
		}
	}

	return nil, false
}

func (cmgr *ClientMgr) SocketEmit(clnt connmgr.IClientDevice, t string, d map[string]any) {
	sse.Emit(clnt.MacAddr(), t, d)
}

func (cmgr *ClientMgr) loopSessions(clnt connmgr.IClientDevice) {
	for nftables.IsConnected(clnt.MacAddr()) {
		errCh := make(chan error)

		go func() {
			cs, err := clnt.ValidSession(context.Background())
			if err != nil {
				errCh <- err
				return
			}

			cmgr.mu.RLock()
			rs, ok := cmgr.getRunningSession(clnt)
			cmgr.mu.RUnlock()

			if !ok {
				rs, err = NewRunningSession(clnt.MacAddr(), clnt.IpAddr(), nil, nil)
				if err != nil {
					errCh <- err
					return
				}

				err = rs.Start(context.Background(), cs)
				log.Println("Start session error: ", err)
				if err != nil {
					errCh <- err
					return
				}

				cmgr.mu.Lock()
				cmgr.sessions = append(cmgr.sessions, rs)
				cmgr.mu.Unlock()
			} else {
				err = rs.Change(cs)
				if err != nil {
					errCh <- err
					return
				}
			}

			err = <-rs.Done()
			log.Println("Running session is done: ", err)

			errCh <- err
		}()

		err := <-errCh
		log.Println("Session done!!! ", err)

		if err != nil {
			log.Println("Error in session loop: ", err)
			cmgr.Disconnect(clnt, err)
			break
		}
	}
}

func (cmgr *ClientMgr) getRunningSession(clnt connmgr.IClientDevice) (rs *RunningSession, ok bool) {
	for _, rs := range cmgr.sessions {
		if rs.GetSession().DeviceId() == clnt.Id() {
			return rs, true
		}
	}
	return nil, false
}

func (cmgr *ClientMgr) endSession(clnt connmgr.IClientDevice) error {
	errCh := make(chan error)

	go func() {
		if nftables.IsConnected(clnt.MacAddr()) {
			err := nftables.Disconnect(clnt.IpAddr(), clnt.MacAddr())
			if err != nil {
				errCh <- err
				return
			}
		}

		cmgr.mu.RLock()
		rs, ok := cmgr.getRunningSession(clnt)
		cmgr.mu.RUnlock()

		if ok {
			err := rs.Stop(context.Background())
			if err != nil {
				errCh <- err
				return
			}

			err = rs.CleanupTc()
			if err != nil {
				errCh <- err
				return
			}
		}

		cmgr.mu.Lock()
		cmgr.sessions = slices.Filter(cmgr.sessions, func(item *RunningSession) bool {
			return item.GetSession().DeviceId() != clnt.Id()
		})
		cmgr.mu.Unlock()

		errCh <- nil
	}()

	return <-errCh
}
