package plugins

import (
	"log"
	"net/http"

	sdkmodels "github.com/flarehotspot/core/sdk/api/models"
	sdkpayments "github.com/flarehotspot/core/sdk/api/payments"
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

func (self *PaymentsApi) Checkout(w http.ResponseWriter, r *http.Request, params sdkpayments.PurchaseRequest) {
}

func (self *PaymentsApi) ExecCallback(w http.ResponseWriter, r *http.Request, purchase sdkmodels.IPurchase) {
}

func (self *PaymentsApi) ConfirmPurchase(w http.ResponseWriter, r *http.Request, purchase sdkmodels.IPurchase) {
	log.Println("TODO: execute purchase")
}

func (self *PaymentsApi) CancelPurchase(w http.ResponseWriter, r *http.Request, purchase sdkmodels.IPurchase) {
	log.Println("TODO: cancel purchase")
}

func (self *PaymentsApi) ParsePaymentInfo(r *http.Request) (*sdkpayments.PaymentInfo, error) {
	return ParsePaymentInfo(self.api.db, self.api.models, r)
}

func (self *PaymentsApi) NewPaymentProvider(provider sdkpayments.IPaymentProvider) {
	log.Println("Registering payment method:", provider.Name())
	self.paymentsMgr.NewPaymentProvider(self.api, provider)
}
