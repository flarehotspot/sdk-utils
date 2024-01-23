package sdkpayments

import (
	connmgr "github.com/flarehotspot/core/sdk/api/connmgr"
)

// IPaymentProvider represents a payment provider.
// A payment provider can have many payment options.
type IPaymentProvider interface {

	// Returns name of the payment provider.
	Name() string

	// Returns a list of available payment options.
	PaymentOpts(clnt connmgr.IClientDevice) []IPaymentOpt
}
