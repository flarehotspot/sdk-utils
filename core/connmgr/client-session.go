package connmgr

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/db/models"
	connmgr "github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/api/models"
)

type ClientSession struct {
	mu        sync.RWMutex
	db        *db.Database
	mdls      *models.Models
	id        int64
	devId     int64
	t         sdkmodels.SessionType
	timeSecs  uint
	dataMb    float64
	timeCons  uint
	dataCons  float64
	startedAt *time.Time
	expDays   *uint
	downMbits int
	upMbits   int
	useGlobal bool
	createdAt time.Time
}

func NewClientSession(dtb *db.Database, mdls *models.Models, s sdkmodels.ISession) connmgr.IClientSession {
	cs := &ClientSession{db: dtb, mdls: mdls}
	cs.load(s)
	return cs
}

func (cs *ClientSession) Id() int64 {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.id
}

func (cs *ClientSession) DeviceId() int64 {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.devId
}

func (cs *ClientSession) Type() sdkmodels.SessionType {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.t
}

func (cs *ClientSession) TimeSecs() (sec uint) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.timeSecs
}

func (cs *ClientSession) DataMb() (mbytes float64) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.dataMb
}

func (cs *ClientSession) TimeConsumption() (sec uint) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.timeCons
}

func (cs *ClientSession) DataConsumption() (mbytes float64) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.dataCons
}

func (cs *ClientSession) StartedAt() *time.Time {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.startedAt
}

func (cs *ClientSession) ExpDays() *uint {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.expDays
}

func (cs *ClientSession) ExpiresAt() *time.Time {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.expiresAt()
}

func (cs *ClientSession) DownMbits() int {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.downMbits
}

func (cs *ClientSession) UpMbits() int {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.upMbits
}

func (cs *ClientSession) UseGlobal() bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.useGlobal
}

func (cs *ClientSession) SetTimeSecs(sec uint) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.timeSecs = sec
}

func (cs *ClientSession) SetDataMb(mbytes float64) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.dataMb = mbytes
}

func (cs *ClientSession) SetTimeCons(sec uint) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.timeCons = sec
}

func (cs *ClientSession) SetDataCons(mbytes float64) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.dataCons = mbytes
}

func (cs *ClientSession) SetStartedAt(started *time.Time) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.startedAt = started
}

func (cs *ClientSession) SetExpDays(exp *uint) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.expDays = exp
}

func (cs *ClientSession) SetDownMbits(mbits int) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.downMbits = mbits
}

func (cs *ClientSession) SetUpMbits(mbits int) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.upMbits = mbits
}

func (cs *ClientSession) SetUseGlobals(g bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.useGlobal = g
}

func (cs *ClientSession) IncTimeCons(sec uint) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.timeCons = cs.timeCons + sec
}

func (cs *ClientSession) IncDataCons(mbytes float64) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.dataCons = cs.dataCons + mbytes
}

func (cs *ClientSession) Update(timeSecs uint, dataMb float64, timeCons uint, dataCons float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	err := cs.mdls.Session().Update(context.Background(), cs.id, cs.devId, cs.t.ToUint8(), timeSecs, dataMb, timeCons, dataCons, started, exp, downMbit, upMbit, g)
	if err != nil {
		return err
	}

	cs.timeSecs = timeSecs
	cs.dataMb = dataMb
	cs.timeCons = timeCons
	cs.dataCons = dataCons
	cs.downMbits = downMbit
	cs.upMbits = upMbit

	return nil
}

func (cs *ClientSession) Save(ctx context.Context) error {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	id := cs.id
	devId := cs.devId
	t := cs.t.ToUint8()
	timeSecs := cs.timeSecs
	dataMb := cs.dataMb
	timeCons := cs.timeCons
	dataCons := cs.dataCons
	started := cs.startedAt
	exp := cs.expDays
	d := cs.downMbits
	u := cs.upMbits
	g := cs.useGlobal

	err := cs.mdls.Session().Update(ctx, id, devId, t, timeSecs, dataMb, timeCons, dataCons, started, exp, d, u, g)
	if err != nil {
		log.Println("Session save error: ", err)
	}

	return err
}

func (cs *ClientSession) SessionModel() sdkmodels.ISession {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return models.BuildSession(
		cs.id,
		cs.devId,
		cs.t.ToUint8(),
		cs.timeSecs,
		cs.dataMb,
		cs.timeCons,
		cs.dataCons,
		cs.startedAt,
		cs.expDays,
		cs.expiresAt(),
		cs.downMbits,
		cs.upMbits,
		cs.useGlobal,
	)
}

func (cs *ClientSession) Reload(ctx context.Context) error {
	s, err := cs.mdls.Session().Find(ctx, cs.id)
	if err != nil {
		return err
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.load(s)

	return nil
}

func (cs *ClientSession) expiresAt() *time.Time {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	started := cs.startedAt
	exp := cs.expDays
	if started != nil && exp != nil {
		exp := cs.startedAt.Add(time.Hour * 24 * time.Duration(*exp))
		return &exp
	}
	return nil
}

func (cs *ClientSession) load(s sdkmodels.ISession) {
	cs.id = s.Id()
	cs.devId = s.DeviceId()
	cs.t = sdkmodels.SessionType(s.SessionType())
	cs.timeSecs = s.TimeSecs()
	cs.dataMb = s.DataMbyte()
	cs.timeCons = s.TimeConsumed()
	cs.dataCons = s.DataConsumed()
	cs.downMbits = s.DownMbits()
	cs.upMbits = s.UpMbits()
	cs.useGlobal = s.UseGlobal()
	cs.expDays = s.ExpDays()
	cs.startedAt = s.StartedAt()
}
