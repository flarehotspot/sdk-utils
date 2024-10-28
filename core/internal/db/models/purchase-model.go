package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"core/internal/db"
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

func (self *PurchaseModel) CreateTx(tx *sql.Tx, ctx context.Context, deviceId int64, sku string, name string, desc string, price float64, vprice bool, pkg string, routename string) (*Purchase, error) {
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
    ) VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, query, deviceId, sku, name, desc, price, vprice, pkg, routename)
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

func (self *PurchaseModel) FindTx(tx *sql.Tx, ctx context.Context, id int64) (*Purchase, error) {
	p := NewPurchase(self.db, self.models)
	attrs := strings.Join(self.attrs, ", ")
	query := "SELECT " + attrs + " FROM purchases WHERE id = ? LIMIT 1"
	err := tx.QueryRowContext(ctx, query, id).
		Scan(&p.id, &p.deviceId, &p.sku, &p.name, &p.description, &p.price, &p.anyPrice, &p.callbackPluginPkg, &p.callbackVueRouteName, &p.walletDebit, &p.walletTxId, &p.confirmedAt, &p.cancelledAt, &p.cancelledReason, &p.createdAt)

	return p, err
}

func (self *PurchaseModel) FindByDeviceIdTx(tx *sql.Tx, ctx context.Context, deviceId int64) (*Purchase, error) {
	p := NewPurchase(self.db, self.models)
	attrs := strings.Join(self.attrs, ", ")
	query := fmt.Sprintf(`
  SELECT %s
  FROM purchases
  WHERE device_id = ?
  LIMIT 1
  `, attrs)

	err := tx.QueryRowContext(ctx, query, deviceId).
		Scan(&p.id, &p.deviceId, &p.sku, &p.name, &p.description, &p.price, &p.anyPrice, &p.callbackPluginPkg, &p.callbackVueRouteName, &p.walletDebit, &p.walletTxId, &p.confirmedAt, &p.cancelledAt, &p.cancelledReason, &p.createdAt)

	return p, err
}

func (self *PurchaseModel) UpdateTx(tx *sql.Tx, ctx context.Context, id int64, dbt float64, txid *int64, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error {
	query := "UPDATE purchases SET wallet_debit = ?, wallet_tx_id = ?, cancelled_at = ?, confirmed_at = ?, cancelled_reason = ? WHERE id = ? LIMIT 1"
	_, err := tx.ExecContext(ctx, query, dbt, txid, cancelledAt, confirmedAt, reason, id)
	return err
}

func (self *PurchaseModel) PendingPurchaseTx(tx *sql.Tx, ctx context.Context, deviceId int64) (*Purchase, error) {
	p := NewPurchase(self.db, self.models)
	attrs := strings.Join(self.attrs, ", ")
	query := fmt.Sprintf(`
  SELECT %s
  FROM purchases
  WHERE confirmed_at IS NULL
  AND cancelled_at IS NULL
  AND device_id = ?
  LIMIT 1
`, attrs)
	err := tx.QueryRowContext(ctx, query, deviceId).
		Scan(&p.id, &p.deviceId, &p.sku, &p.name, &p.description, &p.price, &p.anyPrice, &p.callbackPluginPkg, &p.callbackVueRouteName, &p.walletDebit, &p.walletTxId, &p.confirmedAt, &p.cancelledAt, &p.cancelledReason, &p.createdAt)

	return p, err
}

func (self *PurchaseModel) Create(ctx context.Context, deviceId int64, sku string, name string, desc string, price float64, vprice bool, pkg string, routename string) (*Purchase, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	d, err := self.CreateTx(tx, ctx, deviceId, sku, name, desc, price, vprice, pkg, routename)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	return d, err
}

func (self *PurchaseModel) Find(ctx context.Context, id int64) (*Purchase, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	p, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	return p, err
}

func (self *PurchaseModel) PendingPurchase(ctx context.Context, deviceId int64) (*Purchase, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	d, err := self.PendingPurchaseTx(tx, ctx, deviceId)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	return d, err
}

func (self *PurchaseModel) FindByDeviceId(ctx context.Context, deviceId int64) (*Purchase, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	purchase, err := self.FindByDeviceIdTx(tx, ctx, deviceId)
	if err != nil {
		return nil, err
	}

	return purchase, tx.Commit()
}

func (self *PurchaseModel) Update(ctx context.Context, id int64, dbt float64, txid *int64, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.UpdateTx(tx, ctx, id, dbt, txid, cancelledAt, confirmedAt, reason)
	if err != nil {
		return err
	}

	return tx.Commit()
}
