package sdkpayments

import (
	"context"
	"net/http"

	models "github.com/flarehotspot/core/sdk/api/models"
)

// IPayments is used to handle customer payments.
type IPayments interface {

	// Creates a purchase request and prompts the user for payment.
	Checkout(w http.ResponseWriter, r *http.Request, purchreq PurchaseRequest)

	// Update the paid amount of the purchase request.
	PaymentReceived(ctx context.Context, token string, optname string, amount float64) error

	// Executes the callback URL of a purchase instance after the customer paid for the purchase item(s).
	// It informs the payment requestor that the customer has paid for the purchase.
	ExecCallback(w http.ResponseWriter, r *http.Request, purchase models.IPurchase)

	// Confirms the purchase request. All purchase transactions will be commited to the database.
	// It includes deduction of wallet balance (if available).
	ConfirmPurchase(w http.ResponseWriter, r *http.Request, purchase models.IPurchase)

	// Cancel the purchase request. All purchase transactions will be rolled back.
	CancelPurchase(w http.ResponseWriter, r *http.Request, purchase models.IPurchase)

	// Parses the payment info from the http request.
	// This handles the parsing of http request when callback URL was exected.
	ParsePaymentInfo(r *http.Request) (*PaymentInfo, error)

	// Registers a new payment provider.
	// The provider's payment options will become available for the customers.
	NewPaymentProvider(IPaymentProvider)
}
