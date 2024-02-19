package models

import (
	"context"
	"database/sql"
	"log"

	"github.com/flarehotspot/core/internal/db"
)

type WalletTrnsModel struct {
	db     *db.Database
	models *Models
}

func NewWalletTrnsModel(dtb *db.Database, mdls *Models) *WalletTrnsModel {
	return &WalletTrnsModel{dtb, mdls}
}

func (self *WalletTrnsModel) CreateTx(tx *sql.Tx, ctx context.Context, wltId int64, amount float64, newBal float64, desc string) (*WalletTrns, error) {
	query := "INSERT INTO wallet_transactions (wallet_id, amount, new_balance, description) VALUES(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, wltId, amount, newBal, desc)
	if err != nil {
		log.Println("SQL Exec Error: ", err)
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println("SQL Exec Error: ", err)
		return nil, err
	}

	return self.FindTx(tx, ctx, lastId)
}

func (self *WalletTrnsModel) FindTx(tx *sql.Tx, ctx context.Context, id int64) (*WalletTrns, error) {
	trns := NewWalletTrns(self.db, self.models)
	query := "SELECT id, wallet_id, amount, new_balance, description, created_at FROM wallet_transactions WHERE id = ? LIMIT 1"
	err := tx.QueryRowContext(ctx, query, id).
		Scan(&trns.id, &trns.walletId, &trns.amount, &trns.newBalance, &trns.description, &trns.createdAt)

	return trns, err
}

func (self *WalletTrnsModel) Create(ctx context.Context, wltId int64, amount float64, newBal float64, desc string) (*WalletTrns, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	trns, err := self.CreateTx(tx, ctx, wltId, amount, newBal, desc)
	if err != nil {
		return nil, err
	}

	return trns, tx.Commit()
}

func (self *WalletTrnsModel) Find(ctx context.Context, id int64) (*WalletTrns, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	trns, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, err
	}

	return trns, tx.Commit()
}
