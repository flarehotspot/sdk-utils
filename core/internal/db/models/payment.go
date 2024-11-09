package models

import (
	"context"
	"fmt"
	"time"

	"core/internal/db"

	"github.com/jackc/pgx/v5"
)

type Payment struct {
	db         *db.Database
	models     *Models
	id         int64
	purchaseId int64
	amount     float64
	optname    string
	createdAt  time.Time
}

func NewPayment(dtb *db.Database, mdls *Models) *Payment {
	return &Payment{
		db:     dtb,
		models: mdls,
	}
}

func (self *Payment) Id() int64 {
	return self.id
}

func (self *Payment) PurchaseId() int64 {
	return self.purchaseId
}

func (self *Payment) Amount() float64 {
	return self.amount
}

func (self *Payment) OptName() string {
	return self.optname
}

func (self *Payment) CreatedAt() time.Time {
	return self.createdAt
}

func (self *Payment) UpdateTx(tx pgx.Tx, ctx context.Context, amt float64) error {
	err := self.models.paymentModel.UpdateTx(tx, ctx, self.id, amt)
	if err != nil {
		return fmt.Errorf("could not update payment model: %w", err)
	}

	self.amount = amt
	return nil
}

func (self *Payment) Update(ctx context.Context, amt float64) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.UpdateTx(tx, ctx, amt)
	if err != nil {
		return fmt.Errorf("could not update payment: %w", err)
	}

	return tx.Commit(ctx)
}
