package sdkmdls

import (
	"context"
	"database/sql"
)

// IDeviceModel is used to query devices table in the database.
type IDeviceModel interface {

	// Creates a new device record using a transaction.
	CreateTx(tx *sql.Tx, ctx context.Context, mac string, ip string, hostname string) (IDevice, error)

	// Find a device record from devices table using a transaction with the given ID.
	FindTx(tx *sql.Tx, ctx context.Context, id int64) (IDevice, error)

	// Find a device record from devices table using a transaction with the given MAC address.
	FindByMacTx(tx *sql.Tx, ctx context.Context, mac string) (IDevice, error)

	// Update a device record using a trancation.
	UpdateTx(tx *sql.Tx, ctx context.Context, id int64, hostname string, mac string, ip string) error

	// Create a new database record in devices table.
	Create(ctx context.Context, mac string, ip string, hostname string) (IDevice, error)

	// Find a device record with the given ID.
	Find(ctx context.Context, id int64) (IDevice, error)

	// Find a device record with the given MAC address.
	FindByMac(ctx context.Context, mac string) (IDevice, error)

	// Update a device record.
	Update(ctx context.Context, id int64, hostname string, mac string, ip string) error
}
