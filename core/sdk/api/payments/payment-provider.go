package sdkpayments

import (
	connmgr "github.com/flarehotspot/flarehotspot/core/sdk/api/connmgr"
)

// PaymentProvider represents a payment provider.
// A payment provider can have many payment options.
type PaymentProvider interface {

	// Returns name of the payment provider.
	Name() string

	// Returns a list of available payment options.
	PaymentOpts(clnt connmgr.ClientDevice) []PaymentOpt
}
