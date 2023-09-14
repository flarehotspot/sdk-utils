package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/flarehotspot/core/db"
)

type WalletTrns struct {
	db          *db.Database
	models      *Models
	id          int64
	walletId    int64
	amount      float64
	newBalance  float64
	description string
	createdAt   time.Time
}

func NewWalletTrns(dtb *db.Database, mdls *Models) *WalletTrns {
	return &WalletTrns{
		db:     dtb,
		models: mdls,
	}
}

func (self *WalletTrns) Id() int64 {
	return self.id
}

func (self *WalletTrns) WalletId() int64 {
	return self.walletId
}

func (self *WalletTrns) Amount() float64 {
	return self.amount
}

func (self *WalletTrns) NewBalance() float64 {
	return self.newBalance
}

func (self *WalletTrns) Description() string {
	return self.description
}

func (self *WalletTrns) CreatedAt() time.Time {
	return self.createdAt
}

func (self *WalletTrns) UpdateTx(tx *sql.Tx, ctx context.Context, walletId int64, amount float64, newbal float64, desc string) error {
	query := "UPDATE wallet_transactions SET wallet_id = ?, amount = ?, new_balance = ?, description = ? WHERE id = ? LIMIT 1"
	_, err := tx.ExecContext(ctx, query, walletId, amount, newbal, desc, self.id)
	if err != nil {
		return err
	}
	self.walletId = walletId
	self.amount = amount
	self.newBalance = newbal
	self.description = desc
	return nil
}
