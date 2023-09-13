package payments

import (
	"net/http"

	"github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/api/models"
)

// IPaymentOpt represents a payment option.
type IPaymentOpt interface {

	// Returns the name of the payment option.
	Name() string

	// Handles the payment request.
	PaymentHandler(w http.ResponseWriter, r *http.Request, client connmgr.IClientDevice, purchase models.IPurchase)
}
