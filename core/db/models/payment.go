package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/flarehotspot/core/db"
)

type Payment struct {
	db         *db.Database
	models     *Models
	id         int64
	purchaseId int64
	amount     float64
	method     string
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

func (self *Payment) Method() string {
	return self.method
}

func (self *Payment) CreatedAt() time.Time {
	return self.createdAt
}

func (self *Payment) UpdateTx(tx *sql.Tx, ctx context.Context, amt float64) error {
	err := self.models.paymentModel.UpdateTx(tx, ctx, self.id, amt)
	if err != nil {
		return err
	}

	self.amount = amt
	return nil
}

func (self *Payment) Update(ctx context.Context, amt float64) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.UpdateTx(tx, ctx, amt)
	if err != nil {
		return err
	}

	return tx.Commit()
}
