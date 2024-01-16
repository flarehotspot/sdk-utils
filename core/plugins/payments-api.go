package plugins

import (
	"log"
	"net/http"

	"github.com/flarehotspot/core/payments"
	"github.com/flarehotspot/core/sdk/api/models"
	paymentsI "github.com/flarehotspot/core/sdk/api/payments"
	"github.com/flarehotspot/core/web/response"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes/names"
)

type PaymentsApi struct {
	api         *PluginApi
	paymentsMgr *payments.PaymentsMgr
}

func (self *PaymentsApi) Checkout(w http.ResponseWriter, r *http.Request, params *paymentsI.PurchaseRequest) {
	url, err := router.UrlForRoute(routenames.RoutePaymentOptions)
	if err != nil {
		response.Error(w, err)
		return
	}
	query, err := params.ToQueryParams()
	if err != nil {
		response.Error(w, err)
	}
	url = url + "?" + query
	response.Redirect(w, r, url, http.StatusSeeOther)
}

func (self *PaymentsApi) ExecCallback(w http.ResponseWriter, r *http.Request, purchase models.IPurchase) {
	url := purchase.CallbackUrl() + "?token=" + purchase.Token()
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (self *PaymentsApi) ConfirmPurchase(w http.ResponseWriter, r *http.Request, purchase models.IPurchase) {
	log.Println("TODO: execute purchase")
}

func (self *PaymentsApi) CancelPurchase(w http.ResponseWriter, r *http.Request, purchase models.IPurchase) {
	log.Println("TODO: cancel purchase")
}

func (self *PaymentsApi) ParsePaymentInfo(r *http.Request) (*paymentsI.PaymentInfo, error) {
	return payments.ParsePaymentInfo(self.api.db, self.api.models, r)
}

func (self *PaymentsApi) NewPaymentProvider(provider paymentsI.IPaymentProvider) {
	log.Println("Registering payment method:", provider.Name())
	self.paymentsMgr.NewPaymentProvider(self.api, provider)
}

func NewPaymentsApi(plugin *PluginApi, pmgr *payments.PaymentsMgr) *PaymentsApi {
	return &PaymentsApi{
		api:         plugin,
		paymentsMgr: pmgr,
	}
}
