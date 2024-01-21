package sdkmodels

import (
	"context"
	"database/sql"
)

// IPaymentModel is used to query payments table in the database.
type IPaymentModel interface {
	// Creates a payment record for a purchase using a database transaction.
	CreateTx(tx *sql.Tx, ctx context.Context, purchaseId int64, amt float64, mtd string) (IPayment, error)

	// Finds all payments for a given purchase using a database transaction.
	FindAllByPurchaseTx(tx *sql.Tx, ctx context.Context, purchaseId int64) ([]IPayment, error)

	// Creates a payment record for a purchase.
	Create(ctx context.Context, purid int64, amt float64, mtd string) (IPayment, error)

	// Finds all payments for a given purchase.
	FindAllByPurchase(ctx context.Context, purchaseId int64) ([]IPayment, error)
}
