package models

import (
	"context"
	"database/sql"
)

// IWalletModel is used to query wallets table in the database.
type IWalletModel interface {

	// Finds the wallet by the device id using a database transaction.
	FindByDeviceTx(tx *sql.Tx, ctx context.Context, devId int64) (IWallet, error)

	// Finds the wallet by the device id.
	FindByDevice(ctx context.Context, devId int64) (IWallet, error)
}
