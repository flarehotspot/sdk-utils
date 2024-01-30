package sdkmdls

import (
	"context"
	"database/sql"
	"time"
)

// IPayment represents a payment record in the database.
type IPayment interface {
	// Returns the payment ID.
	Id() int64

	// Returns the purchase ID.
	PurchaseId() int64

	// Returns the amount of the payment.
	Amount() float64

	// Returns the payment method.
	OptName() string

	// Returns when the payment was created.
	CreatedAt() time.Time

	// Updates the payment record in the database using a database transaction.
	UpdateTx(tx *sql.Tx, ctx context.Context, amt float64) error

	// Updates the payment record in the database.
	Update(ctx context.Context, amt float64) error
}
