package sdkmodels

import (
	"context"
	"database/sql"
)

// IDevice is a device record from the database.
type IDevice interface {
	// Returns the device ID.
	Id() int64

	// Returns the device IP address.
	IpAddress() string

	// Returns the device MAC address.
	MacAddress() string

	// Returns the device hostname.
	Hostname() string

	// Reloads the device data in the database using a transaction.
	ReloadTx(tx *sql.Tx, ctx context.Context) error

	// Updates the device data in the database using a transaction.
	UpdateTx(tx *sql.Tx, ctx context.Context, mac string, ip string, hostname string) error

	// Returns the device wallet using a transaction.
	WalletTx(tx *sql.Tx, ctx context.Context) (IWallet, error)

	// Returns the device sessions using a transaction.
	SessionsTx(tx *sql.Tx, ctx context.Context) ([]ISession, error)

	// Reloads the device data in the database.
	Reload(ctx context.Context) error

	// Updates the device data in the database.
	Update(ctx context.Context, mac string, ip string, hostname string) error

	// Returns the device wallet.
	Wallet(ctx context.Context) (IWallet, error)

	// Returns the device sessions.
	Sessions(ctx context.Context) ([]ISession, error)
}
