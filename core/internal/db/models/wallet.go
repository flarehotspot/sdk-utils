package models

import (
	"context"
	"fmt"
	"time"

	"core/internal/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Wallet struct {
	db        *db.Database
	models    *Models
	id        uuid.UUID
	deviceId  uuid.UUID
	balance   float64
	createdAt time.Time
}

func NewWallet(dtb *db.Database, m *Models) *Wallet {
	return &Wallet{
		db:     dtb,
		models: m,
	}
}

func (self *Wallet) Id() uuid.UUID {
	return self.id
}

func (self *Wallet) DeviceId() uuid.UUID {
	return self.deviceId
}

func (self *Wallet) Balance() float64 {
	return self.balance
}

func (self *Wallet) CreatedAt() time.Time {
	return self.createdAt
}

func (self *Wallet) IncBalanceTx(tx pgx.Tx, ctx context.Context, bal float64) error {
	newbal := self.balance + bal
	err := self.UpdateTx(tx, ctx, newbal)
	if err != nil {
		return err
	}

	self.balance = newbal
	return nil
}

func (self *Wallet) UpdateTx(tx pgx.Tx, ctx context.Context, bal float64) error {
	err := self.models.walletModel.UpdateTx(tx, ctx, self.id, bal)
	if err != nil {
		return err
	}
	self.balance = bal
	return nil
}

func (self *Wallet) AvailableBalTx(tx pgx.Tx, ctx context.Context) (float64, error) {
	pending, err := self.models.purchaseModel.PendingPurchaseTx(tx, ctx, self.deviceId)
	if err != nil {
		return 0, nil
	}

	dbt := pending.WalletDebit()
	return self.balance - dbt, nil
}

func (self *Wallet) IncBalance(ctx context.Context, bal float64) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.IncBalanceTx(tx, ctx, bal)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

func (self *Wallet) Update(ctx context.Context, bal float64) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.UpdateTx(tx, ctx, bal)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

func (self *Wallet) AvailableBal(ctx context.Context) (float64, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	bal, err := self.AvailableBalTx(tx, ctx)
	if err != nil {
		return 0, nil
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("could not commit transaction: %w", err)
	}

	return bal, nil
}
