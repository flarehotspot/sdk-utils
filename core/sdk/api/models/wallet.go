package sdkmdls

import (
	"context"
	"database/sql"
	"time"
)

// Represents a record in the wallets table.
type IWallet interface {

	// Returns the wallet ID.
	Id() int64

	// Returns the wallet device ID it belongs to.
	DeviceId() int64

	// Returns the wallet balance.
	Balance() float64

	// Returns the time when the wallet was created.
	CreatedAt() time.Time

	// Increments the wallet balance by the given amount using a database transaction.
	IncBalanceTx(tx *sql.Tx, ctx context.Context, bal float64) error

	// Updates the wallet balance by the given amount using a database transaction.
	UpdateTx(tx *sql.Tx, ctx context.Context, bal float64) error

	// Returns the available balance using a database transaction.
	AvailableBalTx(tx *sql.Tx, ctx context.Context) (float64, error)

	// Increments the wallet balance by the given amount.
	IncBalance(ctx context.Context, bal float64) error

	// Updates the wallet balance by the given amount.
	Update(ctx context.Context, bal float64) error

	// Returns the available balance.
	AvailableBal(ctx context.Context) (float64, error)
}
