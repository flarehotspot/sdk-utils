package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"core/internal/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Device struct {
	db        *db.Database
	models    *Models
	id        uuid.UUID
	macAddr   string
	ipAddr    string
	hostname  string
	createdAt time.Time
}

func NewDevice(d *db.Database, m *Models) *Device {
	return &Device{db: d, models: m}
}

func BuildDevice(id uuid.UUID, mac string, ip string, hostname string) *Device {
	return &Device{
		id:       id,
		macAddr:  mac,
		ipAddr:   ip,
		hostname: hostname,
	}
}

func (self *Device) Id() uuid.UUID {
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

func (self *Device) ReloadTx(tx pgx.Tx, ctx context.Context) error {
	d, err := self.models.deviceModel.FindTx(tx, ctx, self.id)
	if err != nil {
		return err
	}
	self.hostname = d.Hostname()
	self.ipAddr = d.IpAddress()
	self.macAddr = d.MacAddress()
	return nil
}

func (self *Device) WalletTx(tx pgx.Tx, ctx context.Context) (*Wallet, error) {
	wallet, err := self.models.walletModel.findByDeviceTx(tx, ctx, self.id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		wallet, err = self.models.walletModel.CreateTx(tx, ctx, self.id, 0)
	}
	return wallet, err
}

func (self *Device) UpdateTx(tx pgx.Tx, ctx context.Context, mac string, ip string, hostname string) error {
	err := self.models.deviceModel.UpdateTx(tx, ctx, self.id, mac, ip, hostname)
	if err != nil {
		return err
	}

	self.hostname = hostname
	self.macAddr = mac
	self.ipAddr = ip
	return nil
}

func (self *Device) NextSessionTx(tx pgx.Tx, ctx context.Context) (*Session, error) {
	return self.models.sessionModel.AvlForDevTx(tx, ctx, self.id)
}

func (self *Device) SessionsTx(tx pgx.Tx, ctx context.Context) ([]*Session, error) {
	return self.models.sessionModel.SessionsForDevTx(tx, ctx, self.id)
}

func (self *Device) Reload(ctx context.Context) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = self.ReloadTx(tx, ctx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (self *Device) Update(ctx context.Context, mac string, ip string, hostname string) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = self.UpdateTx(tx, ctx, mac, ip, hostname)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (self *Device) Wallet(ctx context.Context) (*Wallet, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	wallet, err := self.WalletTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	return wallet, tx.Commit(ctx)
}

func (self *Device) NextSession(ctx context.Context) (*Session, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	s, err := self.NextSessionTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	return s, tx.Commit(ctx)
}

func (self *Device) Sessions(ctx context.Context) ([]*Session, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	sessions, err := self.SessionsTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	return sessions, tx.Commit(ctx)
}

func (self *Device) Clone() *Device {
	return &Device{
		db:       self.db,
		models:   self.models,
		id:       self.id,
		macAddr:  self.macAddr,
		ipAddr:   self.ipAddr,
		hostname: self.hostname,
	}
}
