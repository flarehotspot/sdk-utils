package models

import (
	"context"
	"database/sql"
	"time"
)

// IPurchaseModel is used to query the purchases table in the database.
type IPurchaseModel interface {
	// Creates a new purchase in the purchases table using a database transaction.
	CreateTx(tx *sql.Tx, ctx context.Context, deviceId int64, vprice bool, cburl string) (IPurchase, error)

	// Finds a purchase in the purchases table by its id using a database transaction.
	FindTx(tx *sql.Tx, ctx context.Context, id int64) (IPurchase, error)

	// Updates a purchase in the purchases table using a database transaction.
	UpdateTx(tx *sql.Tx, ctx context.Context, id int64, dbt float64, txid *int64, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error

	// Finds a purchase with a pending payment in the purchases table by its device id using a database transaction.
	PendingPurchaseTx(tx *sql.Tx, ctx context.Context, deviceId int64) (IPurchase, error)

	// Finds a purchase in the purchases table by its token using a database transaction.
	FindByTokenTx(tx *sql.Tx, ctx context.Context, token string) (IPurchase, error)

	// Creates a new purchase in the purchases table.
	Create(ctx context.Context, deviceId int64, vprice bool, cburl string) (IPurchase, error)

	// Finds a purchase in the purchases table by its id.
	Find(ctx context.Context, id int64) (IPurchase, error)

	// Updates a purchase in the purchases table.
	Update(ctx context.Context, id int64, dbt float64, txid *int64, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error

	// Finds a purchase with a pending payment in the purchases table by its device id.
	PendingPurchase(ctx context.Context, deviceId int64) (IPurchase, error)

	// Finds a purchase in the purchases table by its token.
	FindByToken(ctx context.Context, token string) (IPurchase, error)
}
