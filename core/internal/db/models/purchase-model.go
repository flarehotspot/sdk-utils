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

type PurchaseModel struct {
	db     *db.Database
	models *Models
	attrs  []string
}

func NewPurchaseModel(dtb *db.Database, mdls *Models) *PurchaseModel {
	attrs := []string{"id", "device_id", "sku", "name", "description", "price", "any_price", "callback_plugin", "callback_vue_route_name", "wallet_debit", "wallet_tx_id", "confirmed_at", "cancelled_at", "cancelled_reason", "created_at"}
	return &PurchaseModel{dtb, mdls, attrs}
}

func (self *PurchaseModel) CreateTx(tx pgx.Tx, ctx context.Context, deviceId uuid.UUID, sku string, name string, desc string, price float64, vprice bool, pkg string, routename string) (*Purchase, error) {
	query := `
    INSERT INTO purchases (
        device_id,
        sku,
        name,
        description,
        price,
        any_price,
        callback_plugin,
        callback_vue_route_name
    ) VALUES($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id`
	var lastInsertId int

	err := tx.QueryRow(ctx, query, deviceId, sku, name, desc, price, vprice, pkg, routename).Scan(&lastInsertId)
	if err != nil {
		log.Println("SQL Exec Error: ", err)
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		log.Printf("SQL transaction commit failed: %v", err)
		return nil, err
	}

	return self.FindTx(tx, ctx, int64(lastInsertId))
}

func (self *PurchaseModel) FindTx(tx pgx.Tx, ctx context.Context, id int64) (*Purchase, error) {
	p := NewPurchase(self.db, self.models)

	attrs := strings.Join(self.attrs, ", ")
	query := "SELECT " + attrs + " FROM purchases WHERE id = $1 LIMIT 1"

	err := tx.QueryRow(ctx, query, id).
		Scan(&p.id, &p.deviceId, &p.sku, &p.name, &p.description, &p.price, &p.anyPrice, &p.callbackPluginPkg, &p.callbackVueRouteName, &p.walletDebit, &p.walletTxId, &p.confirmedAt, &p.cancelledAt, &p.cancelledReason, &p.createdAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("No purchase found with id %d", id)
			return nil, nil
		}
		log.Printf("Error finding purchase with id %d: %v", id, err)
	}

	return p, err
}

func (self *PurchaseModel) FindByDeviceIdTx(tx pgx.Tx, ctx context.Context, deviceId int64) (*Purchase, error) {
	p := NewPurchase(self.db, self.models)

	attrs := strings.Join(self.attrs, ", ")
	query := fmt.Sprintf(`
  SELECT %s
  FROM purchases
  WHERE device_id = $1
  LIMIT 1
  `, attrs)

	err := tx.QueryRow(ctx, query, deviceId).
		Scan(&p.id, &p.deviceId, &p.sku, &p.name, &p.description, &p.price, &p.anyPrice, &p.callbackPluginPkg, &p.callbackVueRouteName, &p.walletDebit, &p.walletTxId, &p.confirmedAt, &p.cancelledAt, &p.cancelledReason, &p.createdAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("No purchase found with device id %d", deviceId)
			return nil, nil
		}
		log.Printf("Error finding purchase with device id %d: %v", deviceId, err)
	}

	return p, err
}

func (self *PurchaseModel) UpdateTx(tx pgx.Tx, ctx context.Context, id uuid.UUID, dbt float64, txid *uuid.UUID, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error {
	query := "UPDATE purchases SET wallet_debit = $1, wallet_tx_id = $2, cancelled_at = $3, confirmed_at = $4, cancelled_reason = $5 WHERE id = $6 LIMIT 1"
	cmdTag, err := tx.Exec(ctx, query, dbt, txid, cancelledAt, confirmedAt, reason, id)

	if cmdTag.RowsAffected() == 0 {
		log.Printf("No purchase found with id %d; update operation skipped", id)
		return fmt.Errorf("purchase with id %d not found", id)
	}
	return err
}

func (self *PurchaseModel) PendingPurchaseTx(tx pgx.Tx, ctx context.Context, deviceId uuid.UUID) (*Purchase, error) {
	p := NewPurchase(self.db, self.models)
	attrs := strings.Join(self.attrs, ", ")

	query := fmt.Sprintf(`
  SELECT %s
  FROM purchases
  WHERE confirmed_at IS NULL
  AND cancelled_at IS NULL
  AND device_id = $1
  LIMIT 1
`, attrs)

	err := tx.QueryRow(ctx, query, deviceId).
		Scan(&p.id, &p.deviceId, &p.sku, &p.name, &p.description, &p.price, &p.anyPrice, &p.callbackPluginPkg, &p.callbackVueRouteName, &p.walletDebit, &p.walletTxId, &p.confirmedAt, &p.cancelledAt, &p.cancelledReason, &p.createdAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("No purchase found with device id %d", deviceId)
			return nil, nil
		}
		log.Printf("Error finding purchase with device id %d: %v", deviceId, err)
	}

	return p, err
}

func (self *PurchaseModel) Create(ctx context.Context, deviceId uuid.UUID, sku string, name string, desc string, price float64, vprice bool, pkg string, routename string) (*Purchase, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	d, err := self.CreateTx(tx, ctx, deviceId, sku, name, desc, price, vprice, pkg, routename)
	if err != nil {
		return nil, fmt.Errorf("could not create purchase: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return d, nil
}

func (self *PurchaseModel) Find(ctx context.Context, id int64) (*Purchase, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	p, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return p, nil
}

func (self *PurchaseModel) PendingPurchase(ctx context.Context, deviceId uuid.UUID) (*Purchase, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	d, err := self.PendingPurchaseTx(tx, ctx, deviceId)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return d, nil
}

func (self *PurchaseModel) FindByDeviceId(ctx context.Context, deviceId int64) (*Purchase, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	purchase, err := self.FindByDeviceIdTx(tx, ctx, deviceId)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return purchase, nil
}

func (self *PurchaseModel) Update(ctx context.Context, id uuid.UUID, dbt float64, txid *uuid.UUID, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.UpdateTx(tx, ctx, id, dbt, txid, cancelledAt, confirmedAt, reason)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}
