package connmgr

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/flarehotspot/flarehotspot/core/db/models"
	"github.com/flarehotspot/flarehotspot/core/network"
	connmgr "github.com/flarehotspot/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/flarehotspot/core/sdk/api/network"
	jobque "github.com/flarehotspot/flarehotspot/core/utils/job-que"
	"github.com/flarehotspot/flarehotspot/core/utils/tc"
)

var sessionQ *jobque.JobQues = jobque.NewJobQues()

type RunningSession struct {
	mu         sync.RWMutex
	mac        string
	ip         string
	lan        *network.NetworkLan
	tcClassId  *tc.TcClassId
	tcFilter   *tc.TcFilter
	timeTicker *time.Ticker
	tickerDone chan bool
	session    connmgr.ClientSession
	callbacks  []chan error
}

func NewRunningSession(mac string, ip string, s connmgr.ClientSession, classid *tc.TcClassId) (*RunningSession, error) {
	lan, err := network.FindByIp(ip)
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

func (self *RunningSession) GetSession() connmgr.ClientSession {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.session
}

func (self *RunningSession) Lan() *network.NetworkLan {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.lan
}

func (self *RunningSession) Done() <-chan error {
	self.mu.Lock()
	defer self.mu.Unlock()

	ch := make(chan error)
	self.callbacks = append(self.callbacks, ch)
	return ch
}

func (self *RunningSession) Start(ctx context.Context, s connmgr.ClientSession) error {
	_, err := sessionQ.Exec(func() (any, error) {
		self.mu.Lock()
		defer self.mu.Unlock()

		self.session = s

		started := time.Now()
		s.SetStartedAt(&started)
		err := s.Save(ctx)
		if err != nil {
			return nil, err
		}

		if self.tcClassId == nil {
			err := self.prepTc()
			if err != nil {
				return nil, err
			}
		}

		self.initTimeTicker()
		log.Println("Session Tickeself has started...")

		return nil, nil
	})

	return err
}

func (self *RunningSession) Change(cs connmgr.ClientSession) error {
	_, err := sessionQ.Exec(func() (any, error) {
		self.mu.Lock()
		defer self.mu.Unlock()

		self.session = cs

		downMbits, upMbits := cs.DownMbits(), cs.UpMbits()
		if cs.UseGlobal() {
			lan, err := network.FindByIp(self.ip)
			if err != nil {
				return nil, err
			}

			d, u := lan.Bandwidth()
			downMbits, upMbits = int(d), int(u)
		}

		err := self.lan.ChangeClass(self.tcClassId.Uint(), downMbits, upMbits)
		if err != nil {
			return nil, err
		}

		if self.timeTicker != nil {
			self.initTimeTicker()
		}

		return nil, nil
	})

	return err
}

func (self *RunningSession) Stop(ctx context.Context) error {
	_, err := sessionQ.Exec(func() (any, error) {
		self.mu.Lock()
		defer self.mu.Unlock()

		err := self.save(ctx)
		self.cleanUpTickeself()
		self.runCallbacks(err)

		return nil, nil
	})

	return err
}

func (self *RunningSession) CleanupTc() error {
	errCh := make(chan error)

	go func() {
		self.mu.Lock()
		defer self.mu.Unlock()

		if self.tcClassId != nil {
			log.Println("Clean up TC...")
			classid := self.tcClassId.Uint()

			err := self.lan.DelFilter(self.ip, classid)
			if err != nil {
				errCh <- err
				return
			}

			err = self.lan.DelClass(classid)
			self.tcClassId = nil

			errCh <- err
			return
		}

		log.Println("Done cleaning TC.")
		errCh <- nil
	}()

	return <-errCh
}

func (self *RunningSession) UpdateData(stats *sdknet.TrafficData) {
	self.mu.Lock()
	defer self.mu.Unlock()

	download, dlok := stats.Download[self.ip]
	upload, ulok := stats.Upload[self.mac]

	if dlok && ulok {
		dataconMb := float64(download.Bytes+upload.Bytes) / (1 * 1000 * 1000)
		log.Println("CONSUMTION MB: ", dataconMb)
		self.session.IncDataCons(dataconMb)

		if self.isConsumed() {
			log.Println("Session data is consumed!!!")
			go self.Stop(context.Background())
		}
	}
}

func (self *RunningSession) initTimeTicker() {
	tickerCh := make(chan bool)
	ticker := time.NewTicker(time.Second)

	self.timeTicker = ticker
	self.tickerDone = tickerCh

	go func() {
		self.mu.RLock()
		s := self.session
		self.mu.RUnlock()

		for {
			select {
			case <-tickerCh:
				return
			case <-ticker.C:
				go func() {
					self.mu.RLock()
					defer self.mu.RUnlock()

					s.IncTimeCons(1)

					log.Println("time tick...")

					// save every 15s
					if s.TimeConsumption()%15 == 0 {
						err := self.save(context.Background())
						if err != nil {
							log.Println(err)
							go self.Stop(context.Background())
							return
						}
					}

					if self.isConsumed() {
						log.Println("Session time is consumed!!!")
						go self.Stop(context.Background())
					}
				}()
			}
		}
	}()
}

func (self *RunningSession) prepTc() error {
	classid := tc.GetAvailableId()
	defer classid.Cancel()

	lan := self.lan
	s := self.session

	err := lan.CreateClass(classid.Uint(), s.DownMbits(), s.UpMbits())
	if err != nil {
		return err
	}

	err = lan.CreateFilter(self.ip, classid.Uint())
	if err != nil {
		lan.DelClass(classid.Uint())
		return err
	}

	classid.Commit()
	self.tcClassId = &classid

	return nil
}

func (self *RunningSession) cleanUpTickeself() {
	log.Println("Cleaning up tickeself...")
	if self.timeTicker != nil {
		self.timeTicker.Stop()
		self.timeTicker = nil
		self.tickerDone <- true
		self.tickerDone = nil
	}
	log.Println("Done cleaning tickeself.")
}

func (self *RunningSession) save(ctx context.Context) error {
	if self.session != nil {
		if err := self.session.Save(ctx); err != nil {
			return err
		}

		if err := self.session.Reload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (self *RunningSession) runCallbacks(err error) {
	for _, cb := range self.callbacks {
		cb <- err
	}
	self.callbacks = []chan error{}
	log.Println("Done running callbacks.")
}

func (self *RunningSession) expired() (ok bool) {
	expiresAt := self.session.ExpiresAt()
	if expiresAt != nil {
		return !time.Now().Before(*expiresAt)
	}
	return false
}

func (self *RunningSession) isConsumed() bool {
	s := self.session
	t := s.Type()

	if t == models.SessionTypeTime || t == models.SessionTypeTimeOrData {
		isTimeConsumed := s.TimeConsumption() >= s.TimeSecs()
		return isTimeConsumed || self.expired()
	}

	if t == models.SessionTypeData || t == models.SessionTypeTimeOrData {
		isDataConsumed := s.DataConsumption() >= s.DataMb()
		return isDataConsumed || self.expired()
	}

	return false
}
