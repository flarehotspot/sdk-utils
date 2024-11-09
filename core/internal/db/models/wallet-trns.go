package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"core/internal/db"

	"github.com/jackc/pgx/v5"
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

func (self *WalletTrns) UpdateTx(tx pgx.Tx, ctx context.Context, walletId int64, amount float64, newbal float64, desc string) error {
	query := "UPDATE wallet_transactions SET wallet_id = $1, amount = $2, new_balance = $3, description = $4 WHERE id = $5 LIMIT 1"

	cmdTag, err := tx.Exec(ctx, query, walletId, amount, newbal, desc, self.id)
	if err != nil {
		log.Println("SQL Exec Error while updating wallet transaction ID %d: %v", walletId, err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		log.Println("No wallet transaction found with id %d; update operation skipped", walletId)
		return fmt.Errorf("wallet with id %d not found", walletId)
	}

	self.walletId = walletId
	self.amount = amount
	self.newBalance = newbal
	self.description = desc

	log.Printf("Succcessfully updated wallet transaction with id %d", walletId)
	return nil
}
