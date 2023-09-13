package models

import (
	"context"
	"database/sql"
	"time"
)

// PurchaseStat represents the status of a purchase.
type PurchaseStat struct {

	// This is the total payment made for the purchase.
	// It includes payment made using wallet balance.
	PaymentTotal float64

	// This will be the amount deducted from the wallet balance.
	WalletDebit float64

	// This will be the wallet ballance after the purchase.
	WalletBal float64

	// This will be the available wallet balance after the purchase.
	WalletAvailBal float64
}

// IPurchase represents a record in purchases table in the database.
type IPurchase interface {
	// Returns the purchase id.
	Id() int64

	// Returns the device id the purchase belongs to.
	DeviceId() int64

	// Returns the purchase token.
	Token() string

	// Returns true if the purchase has variable price.
	VarPrice() bool

	// Returns the amount deducted from the wallet balance.
	WalletDebit() float64

	// Returns the wallet transaction id.
	WalletTxId() *int64

	// Returns the time when the purchase was confirmed.
	ConfirmedAt() *time.Time

	// Returns the time when the purchase was cancelled.
	CancelledAt() *time.Time

	// Returns the time when the purchase was created.
	CreatedAt() time.Time

	// Returns the callback URL after the purchase is confirmed or cancelled.
	CallbackUrl() string

	// Returns true if the purchase is confirmed.
	IsConfirmed() bool

	// Returns true if the purchase is cancelled.
	IsCancelled() bool

	// Returns true if the purchase is already processed.
	IsProcessed() bool

	// Updates the purchase record in the database using a database transaction.
	UpdateTx(tx *sql.Tx, ctx context.Context, dbt float64, txid *int64, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error

	// Confirms the purchase using a database transaction.
	ConfirmTx(tx *sql.Tx, ctx context.Context) error

	// Cancels the purchase using a database transaction.
	CancelTx(tx *sql.Tx, ctx context.Context) error

	// Returns the purchase items using a database transaction.
	PurchaseItemsTx(tx *sql.Tx, ctx context.Context) ([]IPurchaseItem, error)

	// Adds a payment record for this purchase using a database transaction.
	AddPaymentTx(tx *sql.Tx, ctx context.Context, amount float64, desc string) (IPayment, error)

	// Returns the payments made for this purchase using a database transaction.
	PaymentsTx(tx *sql.Tx, ctx context.Context) ([]IPayment, error)

	// Returns the total amount of payments made for this purchase using a database transaction.
	PaymentsTotalTx(tx *sql.Tx, ctx context.Context) (float64, error)

	// Returns the status of this purchase using a database transaction.
	StatTx(tx *sql.Tx, ctx context.Context) (*PurchaseStat, error)

	// Updates the purchase record in the database.
	Update(ctx context.Context, dbt float64, txid *int64, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error

	// Confirms the purchase.
	Confirm(ctx context.Context) error

	// Cancels the purchase.
	Cancel(ctx context.Context) error

	// Returns the purchase items.
	PurchaseItems(ctx context.Context) ([]IPurchaseItem, error)

	// Adds a payment record for this purchase.
	AddPayment(ctx context.Context, amount float64, desc string) (IPayment, error)

	// Returns the payments made for this purchase.
	Payments(ctx context.Context) ([]IPayment, error)

	// Returns the total amount of payments made for this purchase.
	PaymentsTotal(ctx context.Context) (float64, error)

	// Returns the status of this purchase.
	Stat(ctx context.Context) (*PurchaseStat, error)
}
