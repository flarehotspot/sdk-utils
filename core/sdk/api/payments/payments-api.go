package sdkpayments

import (
	"net/http"
)

// PaymentsApi is used to handle customer payments.
type PaymentsApi interface {

	// Registers a new payment provider.
	// The provider's payment options will become available for the customers.
	NewPaymentProvider(PaymentProvider)

	// Creates a purchase request and prompts the user for payment.
	// It sends HTTP response and must be put as last line in the handler function.
	Checkout(w http.ResponseWriter, r *http.Request, req PurchaseRequest)

    // Returns the pending purchase for the client device.
	GetPendingPurchase(r *http.Request) (Purchase, error)
}
