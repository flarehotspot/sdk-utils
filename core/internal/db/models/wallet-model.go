package models

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"core/internal/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type WalletModel struct {
	db        *db.Database
	models    *Models
	attrs     []string
	id        uuid.UUID
	balance   float64
	createdAt time.Time
}

func NewWalletModel(dtb *db.Database, mdls *Models) *WalletModel {
	attrs := []string{"id", "device_id", "balance", "created_at"}
	return &WalletModel{
		db:     dtb,
		models: mdls,
		attrs:  attrs,
	}
}

func (self *WalletModel) CreateTx(tx pgx.Tx, ctx context.Context, devId uuid.UUID, bal float64) (*Wallet, error) {
	defer func() {
		if err := tx.Rollback(ctx); err != nil && err != pgx.ErrTxClosed {
			log.Printf("Rollback failed: %v", err)
		}
	}()

	query := "INSERT INTO wallets (device_id, balance) VALUES ($1, $2) RETURNING id"

	var lastInsertId int
	err := tx.QueryRow(ctx, query, devId, bal).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return self.FindTx(tx, ctx, int64(lastInsertId))
}

func (self *WalletModel) FindTx(tx pgx.Tx, ctx context.Context, id int64) (*Wallet, error) {
	wallet := NewWallet(self.db, self.models)

	query := fmt.Sprintf("SELECT %s FROM wallets WHERE id = $1 LIMIT 1", strings.Join(self.attrs, ", "))
	err := tx.QueryRow(ctx, query, id).
		Scan(&wallet.id, &wallet.deviceId, &wallet.balance, &wallet.createdAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("No wallet found with id %d", id)
			return nil, nil
		}
		log.Printf("Error finding wallet with id %d: %v", id, err)
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return wallet, nil
}

func (self *WalletModel) UpdateTx(tx pgx.Tx, ctx context.Context, id uuid.UUID, bal float64) error {
	query := "UPDATE wallets SET balance = $1 WHERE id = $2 LIMIT 1"
	cmdTag, err := tx.Exec(ctx, query, bal, id)
	if err != nil {
		log.Printf("SQL Exec Error while updating wallet ID %d: %v", id, err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		log.Printf("No wallet found with id %d; update operation skipped", id)
		return fmt.Errorf("wallet with id %d not found", id)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	log.Printf("Successfully updated wallet with id %d", id)
	return nil
}

func (self *WalletModel) Find(ctx context.Context, id int64) (*Wallet, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	wallet, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (self *WalletModel) Update(ctx context.Context, id uuid.UUID, bal float64) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.UpdateTx(tx, ctx, id, bal)
	if err != nil {
		return err
	}

	return nil
}

func (self *WalletModel) findByDeviceTx(tx pgx.Tx, ctx context.Context, devId uuid.UUID) (*Wallet, error) {
	wallet := NewWallet(self.db, self.models)
	query := fmt.Sprintf("SELECT %s FROM wallets WHERE device_id = $1 LIMIT 1", strings.Join(self.attrs, ", "))
	err := tx.QueryRow(ctx, query, devId).
		Scan(&wallet.id, &wallet.deviceId, &wallet.balance, &wallet.createdAt)

	if err != nil {
		log.Println("Error finding wallet for device id "+devId.String(), err.Error())
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return wallet, nil
}

func (self *WalletModel) findByDevice(ctx context.Context, devId uuid.UUID) (*Wallet, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	wallet, err := self.findByDeviceTx(tx, ctx, devId)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
