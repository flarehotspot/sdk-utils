package sdkmodels

import (
	"context"
	"database/sql"
)

// IWalletTrnsModel is used to query the wallet_transactions table in the database.
type IWalletTrnsModel interface {
	// Creates a new wallet transaction record using a database transaction.
	CreateTx(tx *sql.Tx, ctx context.Context, wltId int64, amount float64, newBal float64, desc string) (IWalletTrns, error)

	// Finds a wallet transaction record using a database transaction.
	FindTx(tx *sql.Tx, ctx context.Context, id int64) (IWalletTrns, error)

	// Creates a new wallet transaction record by the given id.
	Create(ctx context.Context, wltId int64, amount float64, newBal float64, desc string) (IWalletTrns, error)

	// Finds a wallet transaction record by the given id.
	Find(ctx context.Context, id int64) (IWalletTrns, error)
}
