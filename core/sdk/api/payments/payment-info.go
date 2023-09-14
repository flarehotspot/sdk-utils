package payments

import (
	"github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/api/models"
)

// PaymentInfo represents a payment information.
type PaymentInfo struct {

	// Client is the client device.
	Client connmgr.IClientDevice

	// Purchase is the purchase record.
	Purchase models.IPurchase

	// Payments are the payments made for the purchase.
	Payments []models.IPayment
}
