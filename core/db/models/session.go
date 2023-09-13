package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/flarehotspot/core/db"
)

type Session struct {
	db          *db.Database
	models      *Models
	id          int64
	deviceId    int64
	sessionType uint8
	timeSecs    uint
	dataMb      float64
	timeCons    uint
	dataCons    float64
	startedAt   *time.Time
	expDays     *uint
	expiresAt   *time.Time
	downMbits   int
	upMbits     int
	useGlobal   bool
	createdAt   time.Time
}

func NewSession(dtb *db.Database, mdls *Models) *Session {
	return &Session{
		db:     dtb,
		models: mdls,
	}
}

func BuildSession(id int64, devId int64, t uint8, timeSecs uint, dataMb float64, timeCons uint, dataCons float64, startedAt *time.Time, expDays *uint, expiresAt *time.Time, dmbits int, umbits int, g bool) *Session {
	return &Session{
		id:          id,
		deviceId:    devId,
		sessionType: t,
		timeSecs:    timeSecs,
		dataMb:      dataMb,
		timeCons:    timeCons,
		dataCons:    dataCons,
		startedAt:   startedAt,
		expDays:     expDays,
		expiresAt:   expiresAt,
		downMbits:   dmbits,
		upMbits:     umbits,
		useGlobal:   g,
	}
}

func (self *Session) Id() int64 {
	return self.id
}

func (self *Session) DeviceId() int64 {
	return self.deviceId
}

func (self *Session) SessionType() uint8 {
	return self.sessionType
}

func (self *Session) TimeSecs() uint {
	return self.timeSecs
}

func (self *Session) DataMbyte() float64 {
	return self.dataMb
}

func (self *Session) TimeConsumed() uint {
	return self.timeCons
}

func (self *Session) DataConsumed() float64 {
	return self.dataCons
}

func (self *Session) StartedAt() *time.Time {
	return self.startedAt
}

func (self *Session) ExpDays() *uint {
	return self.expDays
}

func (self *Session) CretedAt() time.Time {
	return self.createdAt
}

func (self *Session) ExpiresAt() *time.Time {
	if self.startedAt != nil && self.expDays != nil {
		exp := self.startedAt.Add(time.Hour * 24 * time.Duration(*self.expDays))
		return &exp
	}
	return nil
}

func (self *Session) DownMbits() int {
	return self.downMbits
}

func (self *Session) UpMbits() int {
	return self.upMbits
}

func (self *Session) UseGlobal() bool {
	return self.useGlobal
}

func (self *Session) CreatedAt() time.Time {
	return self.createdAt
}

func (self *Session) UpdateTx(tx *sql.Tx, ctx context.Context, devId int64, t uint8, minutes uint, mb float64, timecon uint, datacon float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error {
	err := self.models.sessionModel.UpdateTx(tx, ctx, self.id, devId, t, minutes, mb, timecon, datacon, started, exp, downMbit, upMbit, g)
	if err != nil {
		return err
	}

	self.deviceId = devId
	self.sessionType = t
	self.timeSecs = minutes
	self.dataMb = mb
	self.timeCons = timecon
	self.dataCons = datacon
	self.startedAt = started
	self.downMbits = downMbit
	self.upMbits = upMbit
	return nil
}

func (self *Session) SaveTx(tx *sql.Tx, ctx context.Context) error {
	return self.UpdateTx(tx, ctx, self.deviceId, self.sessionType, self.timeSecs, self.dataMb, self.timeCons, self.dataCons, self.startedAt, self.expDays, self.downMbits, self.upMbits, self.useGlobal)
}

func (self *Session) Update(ctx context.Context, devId int64, t uint8, time uint, mb float64, timecon uint, datacon float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.UpdateTx(tx, ctx, devId, t, time, mb, timecon, datacon, started, exp, downMbit, upMbit, g)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (self *Session) Save(ctx context.Context) error {
	return self.Update(ctx, self.deviceId, self.sessionType, self.timeSecs, self.dataMb, self.timeCons, self.dataCons, self.startedAt, self.expDays, self.downMbits, self.upMbits, self.useGlobal)
}
