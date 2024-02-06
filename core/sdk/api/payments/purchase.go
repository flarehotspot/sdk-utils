package sdkpayments

import ()

// Purchase represents a record in purchases table in the database.
type Purchase interface {
	Name() string
	FixedPrice() (float64, bool)
	TotalPayments() (float64, error)
	WalletDebit() float64
	WalletBalance() (float64, error)
	CreatePayment(amount float64, optname string) error
	PayWithWallet(amount float64) error
}
