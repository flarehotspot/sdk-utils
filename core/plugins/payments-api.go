package plugins

import (
	"log"
	"net/http"

	sdkpayments "github.com/flarehotspot/core/sdk/api/payments"
	"github.com/flarehotspot/core/web/helpers"
	routenames "github.com/flarehotspot/core/web/routes/names"
)

func NewPaymentsApi(plugin *PluginApi, pmgr *PaymentsMgr) *PaymentsApi {
	return &PaymentsApi{
		api:         plugin,
		paymentsMgr: pmgr,
	}
}

type PaymentsApi struct {
	api         *PluginApi
	paymentsMgr *PaymentsMgr
}

func (self *PaymentsApi) NewPaymentProvider(provider sdkpayments.PaymentProvider) {
	log.Println("Registering payment method:", provider.Name())
	self.paymentsMgr.NewPaymentProvider(self.api, provider)
}

func (self *PaymentsApi) Checkout(w http.ResponseWriter, r *http.Request, p sdkpayments.PurchaseRequest) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		coreApi := self.api.CoreAPI
		coreApi.HttpAPI.VueResponse().Redirect(w, routenames.RoutePaymentOptions)
	}

	purMw := self.api.HttpAPI.middlewares.PendingPurchaseMw()
	purMw(http.HandlerFunc(handler)).ServeHTTP(w, r)
}

func (self *PaymentsApi) GetPendingPurchase(r *http.Request) (sdkpayments.Purchase, error) {
	mdls := self.api.models
	clnt, err := helpers.CurrentClient(r)
	if err != nil {
        log.Println("helpers.CurrentClient error:", err)
		return nil, err
	}
	p, err := mdls.Purchase().FindByDeviceId(r.Context(), clnt.Id())
	if err != nil {
        log.Println("mdls.Purchase().FindByDeviceId error:", err)
		return nil, err
	}
	purchase := NewPurchase(self.api, r.Context(), clnt.Id(), p)
	return purchase, nil
}
