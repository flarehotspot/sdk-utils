package sdkmodels

import (
	"context"
	"database/sql"
)

// IPurchaseItemModel is used to query purchase_items table in the database.
type IPurchaseItemModel interface {

	// Creates a new purchase item with the given purchase id, sku, name, description, and price using a database transaction.
	CreateTx(tx *sql.Tx, ctx context.Context, purchaseId int64, sku string, name string, desc string, price float64) (IPurchaseItem, error)

	// Finds a purchase item with the given id using a database transaction.
	FindTx(tx *sql.Tx, ctx context.Context, id int64) (IPurchaseItem, error)

	// Finds all purchase items with the given purchase id using a database transaction.
	FindByPurchaseTx(tx *sql.Tx, ctx context.Context, purchaseId int64) ([]IPurchaseItem, error)

	// Creates a new purchase item with the given purchase id, sku, name, description, and price.
	Create(ctx context.Context, purchaseId int64, sku string, name string, desc string, price float64) (IPurchaseItem, error)

	// Finds a purchase item with the given id.
	Find(ctx context.Context, id int64) (IPurchaseItem, error)

	// Finds all purchase items with the given purchase id.
	FindByPurchase(ctx context.Context, purchaseId int64) ([]IPurchaseItem, error)
}
