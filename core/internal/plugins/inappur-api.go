package plugins

import (
	inappur "github.com/flarehotspot/core/sdk/api/inappur"
	"net/http"
)

type InAppPurchaseApi struct {
	plugin *PluginApi
}

func (p *InAppPurchaseApi) VerifyPurchase(inappur.InAppCheckoutItem) error {
	return nil
}

func (p *InAppPurchaseApi) VerifySubscription(inappur.InAppSubscriptionItem) error {
	return nil
}

func (p *InAppPurchaseApi) PurchaseGuardMiddleware(inappur.InAppCheckoutItem) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}

func (p *InAppPurchaseApi) SubscriptionGuardMiddleware(inappur.InAppSubscriptionItem) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}

func NewInAppPurchaseApi(plugin *PluginApi) *InAppPurchaseApi {
	return &InAppPurchaseApi{plugin}
}
