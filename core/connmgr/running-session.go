package connmgr

import (
	"context"
	"log"
	"sync"
	"time"

	coreNet "github.com/flarehotspot/core/network"
	"github.com/flarehotspot/core/utils/tc"
	jobque "github.com/flarehotspot/core/utils/job-que"
	"github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/api/models"
	"github.com/flarehotspot/core/sdk/api/network"
)

var sessionQ *jobque.JobQues = jobque.NewJobQues()

type RunningSession struct {
	mu         sync.RWMutex
	mac        string
	ip         string
	lan        *coreNet.NetworkLan
	tcClassId  *tc.TcClassId
	tcFilter   *tc.TcFilter
	timeTicker *time.Ticker
	tickerDone chan bool
	session    connmgr.IClientSession
	callbacks  []chan error
}

func NewRunningSession(mac string, ip string, s connmgr.IClientSession, classid *tc.TcClassId) (*RunningSession, error) {
	lan, err := coreNet.FindByIp(ip)
	if err != nil {
		return nil, err
	}

	rs := RunningSession{
		tcClassId: classid,
		session:   s,
		mac:       mac,
		ip:        ip,
		lan:       lan,
		callbacks: []chan error{},
	}

	return &rs, nil
}

func (rs *RunningSession) GetSession() connmgr.IClientSession {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	return rs.session
}

func (rs *RunningSession) Lan() *coreNet.NetworkLan {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	return rs.lan
}

func (rs *RunningSession) Done() <-chan error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	ch := make(chan error)
	rs.callbacks = append(rs.callbacks, ch)
	return ch
}

func (rs *RunningSession) Start(ctx context.Context, s connmgr.IClientSession) error {
	_, err := sessionQ.Exec(func() (interface{}, error) {
		rs.mu.Lock()
		defer rs.mu.Unlock()

		rs.session = s

		started := time.Now()
		s.SetStartedAt(&started)
		err := s.Save(ctx)
		if err != nil {
			return nil, err
		}

		if rs.tcClassId == nil {
			err := rs.prepTc()
			if err != nil {
				return nil, err
			}
		}

		rs.initTimeTicker()
		log.Println("Session Tickers has started...")

		return nil, nil
	})

	return err
}

func (rs *RunningSession) Change(cs connmgr.IClientSession) error {
	_, err := sessionQ.Exec(func() (interface{}, error) {
		rs.mu.Lock()
		defer rs.mu.Unlock()

		rs.session = cs

		downMbits, upMbits := cs.DownMbits(), cs.UpMbits()
		if cs.UseGlobal() {
			lan, err := coreNet.FindByIp(rs.ip)
			if err != nil {
				return nil, err
			}

			d, u := lan.Bandwidth()
			downMbits, upMbits = int(d), int(u)
		}

		err := rs.lan.ChangeClass(rs.tcClassId.Uint(), downMbits, upMbits)
		if err != nil {
			return nil, err
		}

		if rs.timeTicker != nil {
			rs.initTimeTicker()
		}

		return nil, nil
	})

	return err
}

func (rs *RunningSession) Stop(ctx context.Context) error {
	_, err := sessionQ.Exec(func() (interface{}, error) {
		rs.mu.Lock()
		defer rs.mu.Unlock()

		err := rs.save(ctx)
		rs.cleanUpTickers()
		rs.runCallbacks(err)

		return nil, nil
	})

	return err
}

func (rs *RunningSession) CleanupTc() error {
	errCh := make(chan error)

	go func() {
		rs.mu.Lock()
		defer rs.mu.Unlock()

		if rs.tcClassId != nil {
			log.Println("Clean up TC...")
			classid := rs.tcClassId.Uint()

			err := rs.lan.DelFilter(rs.ip, classid)
			if err != nil {
				errCh <- err
				return
			}

			err = rs.lan.DelClass(classid)
			rs.tcClassId = nil

			errCh <- err
			return
		}

		log.Println("Done cleaning TC.")
		errCh <- nil
	}()

	return <-errCh
}

func (rs *RunningSession) UpdateData(stats *network.TrafficData) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	download, dlok := stats.Download[rs.ip]
	upload, ulok := stats.Upload[rs.mac]

	if dlok && ulok {
		dataconMb := float64(download.Bytes+upload.Bytes) / (1 * 1000 * 1000)
		log.Println("CONSUMTION MB: ", dataconMb)
		rs.session.IncDataCons(dataconMb)

		if rs.isConsumed() {
			log.Println("Session data is consumed!!!")
			go rs.Stop(context.Background())
		}
	}
}

func (rs *RunningSession) initTimeTicker() {
	tickerCh := make(chan bool)
	ticker := time.NewTicker(time.Second)

	rs.timeTicker = ticker
	rs.tickerDone = tickerCh

	go func() {
		rs.mu.RLock()
		s := rs.session
		rs.mu.RUnlock()

		for {
			select {
			case <-tickerCh:
				return
			case <-ticker.C:
				go func() {
					rs.mu.RLock()
					defer rs.mu.RUnlock()

					s.IncTimeCons(1)

					log.Println("time tick...")

					// save every 15s
					if s.TimeConsumption()%15 == 0 {
						err := rs.save(context.Background())
						if err != nil {
							log.Println(err)
							go rs.Stop(context.Background())
							return
						}
					}

					if rs.isConsumed() {
						log.Println("Session time is consumed!!!")
						go rs.Stop(context.Background())
					}
				}()
			}
		}
	}()
}

func (rs *RunningSession) prepTc() error {
	classid := tc.GetAvailableId()
	defer classid.Cancel()

	lan := rs.lan
	s := rs.session

	err := lan.CreateClass(classid.Uint(), s.DownMbits(), s.UpMbits())
	if err != nil {
		return err
	}

	err = lan.CreateFilter(rs.ip, classid.Uint())
	if err != nil {
		lan.DelClass(classid.Uint())
		return err
	}

	classid.Commit()
	rs.tcClassId = &classid

	return nil
}

func (rs *RunningSession) cleanUpTickers() {
	log.Println("Cleaning up tickers...")
	if rs.timeTicker != nil {
		rs.timeTicker.Stop()
		rs.timeTicker = nil
		rs.tickerDone <- true
		rs.tickerDone = nil
	}
	log.Println("Done cleaning tickers.")
}

func (rs *RunningSession) save(ctx context.Context) error {
	if rs.session != nil {
		if err := rs.session.Save(ctx); err != nil {
			return err
		}

		if err := rs.session.Reload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (rs *RunningSession) runCallbacks(err error) {
	for _, cb := range rs.callbacks {
		cb <- err
	}
	rs.callbacks = []chan error{}
	log.Println("Done running callbacks.")
}

func (rs *RunningSession) expired() (ok bool) {
	expiresAt := rs.session.ExpiresAt()
	if expiresAt != nil {
		return !time.Now().Before(*expiresAt)
	}
	return false
}

func (rs *RunningSession) isConsumed() bool {
	s := rs.session
	t := s.Type()

	if t == models.SessionTypeTime || t == models.SessionTypeTimeOrData {
		isTimeConsumed := s.TimeConsumption() >= s.TimeSecs()
		return isTimeConsumed || rs.expired()
	}

	if t == models.SessionTypeData || t == models.SessionTypeTimeOrData {
		isDataConsumed := s.DataConsumption() >= s.DataMb()
		return isDataConsumed || rs.expired()
	}

	return false
}
