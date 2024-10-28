package plugins

import (
	inappur "sdk/api/inappur"
	"net/http"
)

func NewInAppPurchaseApi(plugin *PluginApi) *InAppPurchaseApi {
	return &InAppPurchaseApi{plugin}
}

type InAppPurchaseApi struct {
	plugin *PluginApi
}

func (self *InAppPurchaseApi) VerifyPurchase(inappur.InAppCheckoutItem) error {
	return nil
}

func (self *InAppPurchaseApi) VerifySubscription(inappur.InAppSubscriptionItem) error {
	return nil
}

func (self *InAppPurchaseApi) PurchaseGuardMiddleware(inappur.InAppCheckoutItem) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}

func (self *InAppPurchaseApi) SubscriptionGuardMiddleware(inappur.InAppSubscriptionItem) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}
