package sdkpayments

import "net/http"

// Purchase represents a record in purchases table in the database.
type Purchase interface {
	Name() string
	FixedPrice() (float64, bool)
	CreatePayment(amount float64, optname string) error
	PayWithWallet(amount float64) error
	State() (PurchaseState, error)
	Execute(w http.ResponseWriter)
	Confirm() error
	Cancel() error
}
