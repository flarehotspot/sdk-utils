package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/flarehotspot/flarehotspot/core/db"
)

type Device struct {
	db        *db.Database
	models    *Models
	id        int64
	macAddr   string
	ipAddr    string
	hostname  string
	createdAt time.Time
}

func NewDevice(d *db.Database, m *Models) *Device {
	return &Device{db: d, models: m}
}

func BuildDevice(id int64, mac string, ip string, hostname string) *Device {
	return &Device{
		id:       id,
		macAddr:  mac,
		ipAddr:   ip,
		hostname: hostname,
	}
}

func (self *Device) Id() int64 {
	return self.id
}

func (self *Device) Hostname() string {
	return self.hostname
}

func (self *Device) IpAddress() string {
	return self.ipAddr
}

func (self *Device) MacAddress() string {
	return self.macAddr
}

func (self *Device) ReloadTx(tx *sql.Tx, ctx context.Context) error {
	d, err := self.models.deviceModel.FindTx(tx, ctx, self.id)
	if err != nil {
		return err
	}
	self.hostname = d.Hostname()
	self.ipAddr = d.IpAddress()
	self.macAddr = d.MacAddress()
	return nil
}

func (self *Device) WalletTx(tx *sql.Tx, ctx context.Context) (*Wallet, error) {
	wallet, err := self.models.walletModel.findByDeviceTx(tx, ctx, self.id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		wallet, err = self.models.walletModel.CreateTx(tx, ctx, self.id, 0)
	}
	return wallet, err
}

func (self *Device) UpdateTx(tx *sql.Tx, ctx context.Context, mac string, ip string, hostname string) error {
	err := self.models.deviceModel.UpdateTx(tx, ctx, self.id, mac, ip, hostname)
	if err != nil {
		return err
	}

	self.hostname = hostname
	self.macAddr = mac
	self.ipAddr = ip
	return nil
}

func (self *Device) NextSessionTx(tx *sql.Tx, ctx context.Context) (*Session, error) {
	return self.models.sessionModel.AvlForDevTx(tx, ctx, self.id)
}

func (self *Device) SessionsTx(tx *sql.Tx, ctx context.Context) ([]*Session, error) {
	return self.models.sessionModel.SessionsForDevTx(tx, ctx, self.id)
}

func (self *Device) Reload(ctx context.Context) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.ReloadTx(tx, ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (self *Device) Update(ctx context.Context, mac string, ip string, hostname string) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.UpdateTx(tx, ctx, mac, ip, hostname)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (self *Device) Wallet(ctx context.Context) (*Wallet, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	wallet, err := self.WalletTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	return wallet, tx.Commit()
}

func (self *Device) NextSession(ctx context.Context) (*Session, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	s, err := self.NextSessionTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	return s, tx.Commit()
}

func (self *Device) Sessions(ctx context.Context) ([]*Session, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	sessions, err := self.SessionsTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	return sessions, tx.Commit()
}
