package sdkmdls

import "time"

// IWalletTrns is used to query the wallet_transactions table.
type IWalletTrns interface {

	// Returns the wallet transaction id.
	Id() int64

	// Returns the ID of the wallet that the transaction is associated with.
	WalletId() int64

	// Returns the amount of the transaction.
	Amount() float64

	// Returns the new balance of the wallet after the transaction.
	NewBalance() float64

	// Returns the description of the transaction.
	Description() string

	// Returns the time when the transaction was created.
	CreatedAt() time.Time
}
