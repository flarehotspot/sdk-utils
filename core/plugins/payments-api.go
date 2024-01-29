package plugins

import (
	"log"
	"net/http"

	sdkmodels "github.com/flarehotspot/core/sdk/api/models"
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

func (self *PaymentsApi) Checkout(w http.ResponseWriter, r *http.Request, p sdkpayments.PurchaseRequest) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		mdls := self.api.models
		res := self.api.HttpAPI.VueResponse()
		clnt, err := helpers.CurrentClient(r)
		if err != nil {
			res.FlashMsg("error", err.Error())
			res.Json(w, nil, http.StatusInternalServerError)
			return
		}

		route, ok := self.api.HttpAPI.vueRouter.FindVueRoute(p.CallbackVueRouteName)
		if !ok {
			res.FlashMsg("error", "Invalid callback route name")
			res.Json(w, nil, http.StatusInternalServerError)
			return
		}

		purchase, err := mdls.Purchase().Create(r.Context(), clnt.Id(), p.Sku, p.Name, p.Description, p.Price, p.AnyPrice, route.VueRouteName)
		if err != nil {
			res.FlashMsg("error", err.Error())
			res.Json(w, nil, http.StatusInternalServerError)
			return
		}

		coreApi := self.api.CoreAPI
		coreApi.HttpAPI.VueResponse().Redirect(w, routenames.RoutePaymentOptions, "token", purchase.Token())
	}

	purMw := self.api.HttpAPI.middlewares.PendingPurchaseMw()
	purMw(http.HandlerFunc(handler)).ServeHTTP(w, r)
}

func (self *PaymentsApi) PaymentReceived(token string, optname string, amount float64) error {
	return nil
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
