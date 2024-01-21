package models

import (
	"context"
	"database/sql"
	"log"

	"github.com/flarehotspot/core/db"
	models "github.com/flarehotspot/core/sdk/api/models"
)

type PurchaseItemModel struct {
	db   *db.Database
	mdls *Models
}

func NewPurchaseItemModel(dtb *db.Database, mdls *Models) *PurchaseItemModel {
	return &PurchaseItemModel{dtb, mdls}
}

func (self *PurchaseItemModel) CreateTx(tx *sql.Tx, ctx context.Context, purchaseId int64, sku string, name string, desc string, price float64) (models.IPurchaseItem, error) {
	query := "INSERT INTO purchase_items (purchase_id, sku, name, description, price) VALUES(?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, purchaseId, sku, name, desc, price)
	if err != nil {
		log.Println("SQL Exec Error: ", err)
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println("SQL Lastid() Error: ", err)
		return nil, err
	}

	return self.FindTx(tx, ctx, lastId)
}

func (self *PurchaseItemModel) FindTx(tx *sql.Tx, ctx context.Context, id int64) (models.IPurchaseItem, error) {
	item := NewPurchaseItem(self.db)
	query := "SELECT id, purchase_id, sku, name, description, price, created_at FROM purchase_items WHERE id = ? LIMIT 1"
	err := tx.QueryRowContext(ctx, query, id).
		Scan(&item.id, &item.purchaseId, &item.sku, &item.name, &item.description, &item.price, &item.createdAt)

	return item, err
}

func (self *PurchaseItemModel) FindByPurchaseTx(tx *sql.Tx, ctx context.Context, purchaseId int64) ([]models.IPurchaseItem, error) {
	query := "SELECT id, purchase_id, sku, name, description, price, created_at FROM purchase_items WHERE purchase_id = ?"
	rows, err := tx.QueryContext(ctx, query, purchaseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []models.IPurchaseItem{}
	for rows.Next() {
		item := NewPurchaseItem(self.db)
		err = rows.Scan(&item.id, &item.purchaseId, &item.sku, &item.name, &item.description, &item.price, &item.createdAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (self *PurchaseItemModel) Create(ctx context.Context, purchaseId int64, sku string, name string, desc string, price float64) (models.IPurchaseItem, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	item, err := self.CreateTx(tx, ctx, purchaseId, sku, name, desc, price)

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (self *PurchaseItemModel) Find(ctx context.Context, id int64) (models.IPurchaseItem, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	item, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, err
	}

	return item, tx.Commit()
}

func (self *PurchaseItemModel) FindByPurchase(ctx context.Context, purchaseId int64) ([]models.IPurchaseItem, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	items, err := self.FindByPurchaseTx(tx, ctx, purchaseId)
	if err != nil {
		return nil, err
	}

	return items, tx.Commit()
}
