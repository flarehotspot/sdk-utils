package models

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/flarehotspot/core/db"
)

type WalletModel struct {
	db        *db.Database
	models    *Models
	id        int64
	balance   float64
	createdAt time.Time
}

func NewWalletModel(dtb *db.Database, mdls *Models) *WalletModel {
	return &WalletModel{
		db:     dtb,
		models: mdls,
	}
}

func (self *WalletModel) CreateTx(tx *sql.Tx, ctx context.Context, devId int64, bal float64) (*Wallet, error) {
	query := "INSERT INTO wallets (device_id, balance) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, query, devId, bal)
	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return self.FindTx(tx, ctx, lastId)
}

func (self *WalletModel) FindTx(tx *sql.Tx, ctx context.Context, id int64) (*Wallet, error) {
	wallet := NewWallet(self.db, self.models)
	query := "SELECT id, device_id, balance, created_at FROM wallets WHERE id = ? LIMIT 1"
	err := tx.QueryRowContext(ctx, query, id).
		Scan(&wallet.id, &wallet.deviceId, &wallet.balance, &wallet.createdAt)

	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (self *WalletModel) FindByDeviceTx(tx *sql.Tx, ctx context.Context, devId int64) (*Wallet, error) {
	wallet := NewWallet(self.db, self.models)
	query := "SELECT id, device_id, balance, created_at FROM wallets WHERE device_id = ? LIMIT 1"
	err := tx.QueryRowContext(ctx, query, devId).
		Scan(&wallet.id, &wallet.deviceId, &wallet.balance, &wallet.createdAt)

	if err != nil {
		log.Println("Error finding wallet for device id "+string(rune(devId)), err.Error())
		return nil, err
	}

	return wallet, nil
}

func (self *WalletModel) UpdateTx(tx *sql.Tx, ctx context.Context, id int64, bal float64) error {
	query := "UPDATE wallets SET balance = ? WHERE id = ? LIMIT 1"
	_, err := tx.ExecContext(ctx, query, bal, id)
	return err
}

func (self *WalletModel) Find(ctx context.Context, id int64) (*Wallet, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	wallet, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, err
	}

	return wallet, tx.Commit()
}

func (self *WalletModel) FindByDevice(ctx context.Context, devId int64) (*Wallet, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	wallet, err := self.FindByDeviceTx(tx, ctx, devId)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	return wallet, err
}

func (self *WalletModel) Update(ctx context.Context, id int64, bal float64) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.UpdateTx(tx, ctx, id, bal)
	if err != nil {
		return err
	}

	return tx.Commit()
}
