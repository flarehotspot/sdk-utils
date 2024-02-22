package sdkinappur

import "net/http"

// InAppPurchasesApi is used to perform purchases and subscriptions.
type InAppPurchasesApi interface {
	// Verify if user has already purchased the item.
	VerifyPurchase(InAppCheckoutItem) error

	// Verify if user has already subscribed to the item.
	VerifySubscription(InAppSubscriptionItem) error

	// This will redirect the user to the purchase page to perform the transaction.
	PurchaseGuardMiddleware(InAppCheckoutItem) (middleware func(next http.Handler) http.Handler)

	// This will redirect the user to the subscription page to perform the transaction.
	SubscriptionGuardMiddleware(InAppSubscriptionItem) (middleware func(next http.Handler) http.Handler)
}
