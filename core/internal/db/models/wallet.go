package models

import (
	"context"
	"database/sql"
	"time"

	"core/internal/db"
)

type Wallet struct {
	db        *db.Database
	models    *Models
	id        int64
	deviceId  int64
	balance   float64
	createdAt time.Time
}

func NewWallet(dtb *db.Database, m *Models) *Wallet {
	return &Wallet{
		db:     dtb,
		models: m,
	}
}

func (self *Wallet) Id() int64 {
	return self.id
}

func (self *Wallet) DeviceId() int64 {
	return self.deviceId
}

func (self *Wallet) Balance() float64 {
	return self.balance
}

func (self *Wallet) CreatedAt() time.Time {
	return self.createdAt
}

func (self *Wallet) IncBalanceTx(tx *sql.Tx, ctx context.Context, bal float64) error {
	newbal := self.balance + bal
	err := self.UpdateTx(tx, ctx, newbal)
	if err != nil {
		return err
	}

	self.balance = newbal
	return nil
}

func (self *Wallet) UpdateTx(tx *sql.Tx, ctx context.Context, bal float64) error {
	err := self.models.walletModel.UpdateTx(tx, ctx, self.id, bal)
	if err != nil {
		return err
	}
	self.balance = bal
	return nil
}

func (self *Wallet) AvailableBalTx(tx *sql.Tx, ctx context.Context) (float64, error) {
	pending, err := self.models.purchaseModel.PendingPurchaseTx(tx, ctx, self.deviceId)
	if err != nil {
		return 0, nil
	}

	dbt := pending.WalletDebit()
	return self.balance - dbt, nil
}

func (self *Wallet) IncBalance(ctx context.Context, bal float64) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.IncBalanceTx(tx, ctx, bal)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (self *Wallet) Update(ctx context.Context, bal float64) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.UpdateTx(tx, ctx, bal)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (self *Wallet) AvailableBal(ctx context.Context) (float64, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	bal, err := self.AvailableBalTx(tx, ctx)
	if err != nil {
		return 0, nil
	}

	return bal, tx.Commit()
}
