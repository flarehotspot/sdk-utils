package models

import (
	"context"
	"fmt"
	"log"

	"core/internal/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type WalletTrnsModel struct {
	db     *db.Database
	models *Models
}

func NewWalletTrnsModel(dtb *db.Database, mdls *Models) *WalletTrnsModel {
	return &WalletTrnsModel{dtb, mdls}
}

func (self *WalletTrnsModel) CreateTx(tx pgx.Tx, ctx context.Context, wltId uuid.UUID, amount float64, newBal float64, desc string) (*WalletTrns, error) {
	query := "INSERT INTO wallet_transactions (wallet_id, amount, new_balance, description) VALUES($1, $2, $3, $4)"

	var lastId int
	err := tx.QueryRow(ctx, query, wltId, amount, newBal, desc).Scan(&lastId)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			log.Printf("Rollback failed: %v", rbErr)
			return nil, err
		}
		log.Printf("SQL Execution failed: %v", err)

		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return self.FindTx(tx, ctx, int64(lastId))
}

func (self *WalletTrnsModel) FindTx(tx pgx.Tx, ctx context.Context, id int64) (*WalletTrns, error) {
	trns := NewWalletTrns(self.db, self.models)
	query := "SELECT id, wallet_id, amount, new_balance, description, created_at FROM wallet_transactions WHERE id = $1 LIMIT 1"
	err := tx.QueryRow(ctx, query, id).
		Scan(&trns.id, &trns.walletId, &trns.amount, &trns.newBalance, &trns.description, &trns.createdAt)
	if err != nil {
		log.Println("Error scanning row:", err)
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return trns, err
}

func (self *WalletTrnsModel) Create(ctx context.Context, wltId uuid.UUID, amount float64, newBal float64, desc string) (*WalletTrns, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	trns, err := self.CreateTx(tx, ctx, wltId, amount, newBal, desc)
	if err != nil {
		return nil, err
	}

	return trns, nil
}

func (self *WalletTrnsModel) Find(ctx context.Context, id int64) (*WalletTrns, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	trns, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, err
	}

	return trns, nil
}
