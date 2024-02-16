package connmgr

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/flarehotspot/flarehotspot/core/db"
	"github.com/flarehotspot/flarehotspot/core/db/models"
	"github.com/flarehotspot/flarehotspot/core/network"
	connmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdknet "github.com/flarehotspot/sdk/api/network"
	slices "github.com/flarehotspot/sdk/utils/slices"
	sse "github.com/flarehotspot/sdk/utils/sse"
	"github.com/flarehotspot/flarehotspot/core/utils/nftables"
)

const (
	EventConnected    string = "client:connected"
	EventDisconnected string = "client:disconnected"
)

func NewSessionsMgr(dtb *db.Database, mdl *models.Models) *SessionsMgr {
	return &SessionsMgr{
		mu:       sync.RWMutex{},
		mdl:      mdl,
		sessions: []*RunningSession{},
	}
}

type SessionsMgr struct {
	mu       sync.RWMutex
	mdl      *models.Models
	sessions []*RunningSession
}

func (self *SessionsMgr) ListenTraffic(trfk *network.TrafficMgr) {
	go func() {
		for data := range trfk.Listen() {
			go func(data *sdknet.TrafficData) {
				self.mu.RLock()
				defer self.mu.RUnlock()

				for _, s := range self.sessions {
					s.UpdateData(data)
				}
			}(&data)
		}
	}()
}

func (self *SessionsMgr) ReloadSessions(ctx context.Context, iface string) error {
	errCh := make(chan error)
	go func() {
		self.mu.RLock()
		defer self.mu.RUnlock()

		for _, rs := range self.sessions {
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

func (self *SessionsMgr) StopSessions(ctx context.Context, iface string, reason string) {
	done := make(chan bool)
	go func() {
		self.mu.Lock()
		defer self.mu.Unlock()
		defer func() {
			done <- true
		}()

		for _, rs := range self.sessions {
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

func (self *SessionsMgr) Connect(clnt connmgr.ClientDevice) error {
	errCh := make(chan error)

	go func() {
		if _, ok := self.CurrSession(clnt); ok {
			errCh <- errors.New("Device is already connected.")
			return
		}

		if !clnt.HasSession(context.Background()) {
			errCh <- errors.New("Device has no available sessions.")
			return
		}

		if !nftables.IsConnected(clnt.MacAddr()) {
			if err := nftables.Connect(clnt.IpAddr(), clnt.MacAddr()); err != nil {
				errCh <- err
				return
			}
		}

		go self.loopSessions(clnt)

		data := map[string]any{"message": "You are now connected to internet."}
		self.SocketEmit(clnt, "client:connected", data)
		errCh <- nil
	}()

	return <-errCh
}

func (self *SessionsMgr) Disconnect(clnt connmgr.ClientDevice, notify error) error {
	log.Println("Calling endsession()...")
	err := self.endSession(clnt)
	if err != nil {
		notify = err
	}

	if notify != nil {
		data := map[string]any{"message": notify.Error()}
		self.SocketEmit(clnt, EventDisconnected, data)
	}

	return err
}

func (self *SessionsMgr) IsConnected(clnt connmgr.ClientDevice) (connected bool) {
	return nftables.IsConnected(clnt.MacAddr())
}

func (self *SessionsMgr) CurrSession(clnt connmgr.ClientDevice) (cs connmgr.ClientSession, ok bool) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	for _, rs := range self.sessions {
		if rs.GetSession().DeviceId() == clnt.Id() {
			return rs.session, true
		}
	}

	return nil, false
}

func (self *SessionsMgr) SocketEmit(clnt connmgr.ClientDevice, t string, d map[string]any) {
	sse.Emit(clnt.MacAddr(), t, d)
}

func (self *SessionsMgr) loopSessions(clnt connmgr.ClientDevice) {
	for nftables.IsConnected(clnt.MacAddr()) {
		errCh := make(chan error)

		go func() {
			cs, err := clnt.ValidSession(context.Background())
			if err != nil {
				errCh <- err
				return
			}

			self.mu.RLock()
			rs, ok := self.getRunningSession(clnt)
			self.mu.RUnlock()

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

				self.mu.Lock()
				self.sessions = append(self.sessions, rs)
				self.mu.Unlock()
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
			self.Disconnect(clnt, err)
			break
		}
	}
}

func (self *SessionsMgr) getRunningSession(clnt connmgr.ClientDevice) (rs *RunningSession, ok bool) {
	for _, rs := range self.sessions {
		if rs.GetSession().DeviceId() == clnt.Id() {
			return rs, true
		}
	}
	return nil, false
}

func (self *SessionsMgr) endSession(clnt connmgr.ClientDevice) error {
	errCh := make(chan error)

	go func() {
		if nftables.IsConnected(clnt.MacAddr()) {
			err := nftables.Disconnect(clnt.IpAddr(), clnt.MacAddr())
			if err != nil {
				errCh <- err
				return
			}
		}

		self.mu.RLock()
		rs, ok := self.getRunningSession(clnt)
		self.mu.RUnlock()

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

		self.mu.Lock()
		self.sessions = slices.Filter(self.sessions, func(item *RunningSession) bool {
			return item.GetSession().DeviceId() != clnt.Id()
		})
		self.mu.Unlock()

		errCh <- nil
	}()

	return <-errCh
}

func (self *SessionsMgr) CreateSession(
	ctx context.Context,
	devId int64,
	t uint8,
	timeSecs uint,
	dataMbytes float64,
	expDays *uint,
	downMbits int,
	upMbits int,
	useGlobal bool,
) error {
	_, err := self.mdl.Session().Create(ctx, devId, t, timeSecs, dataMbytes, expDays, downMbits, upMbits, useGlobal)
	return err
}
