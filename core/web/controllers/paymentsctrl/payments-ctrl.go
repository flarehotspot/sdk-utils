package paymentsctrl

// import (
// 	"database/sql"
// 	"errors"
// 	"log"
// 	"net/http"

// 	"github.com/flarehotspot/flarehotspot/core/db"
// 	coreM "github.com/flarehotspot/flarehotspot/core/db/models"
// 	"github.com/flarehotspot/flarehotspot/core/globals"
// 	"github.com/flarehotspot/flarehotspot/core/payments"
// 	"github.com/flarehotspot/flarehotspot/core/plugins"
// 	mdls "github.com/flarehotspot/sdk/api/models"
// 	"github.com/flarehotspot/sdk/utils/flash"
// 	"github.com/flarehotspot/flarehotspot/core/web/helpers"
// 	"github.com/flarehotspot/flarehotspot/core/web/response"
// 	"github.com/flarehotspot/flarehotspot/core/web/router"
// 	"github.com/flarehotspot/flarehotspot/core/web/routes/names"
// 	"github.com/gorilla/mux"
// )

// type PaymentsCtrl struct {
// 	db          *db.Database
// 	mdls        *coreM.Models
// 	paymentsMgr *payments.PaymentsMgr
// 	capi        *plugins.PluginApi
// 	errR        *response.ErrRedirect
// }

// func NewPaymentsCtrl(g *globals.CoreGlobals) *PaymentsCtrl {
// 	errR := response.NewErrRoute(routenames.RoutePaymentOptions)
// 	return &PaymentsCtrl{g.Db, g.Models, g.PaymentsMgr, g.CoreApi, errR}
// }

// func (self *PaymentsCtrl) PaymentOptions(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	clnt, err := helpers.CurrentClient(r)
// 	if err != nil {
// 		log.Println("Error getting client...")
// 		self.Error(w, r, err)
// 		return
// 	}

// 	methods := map[string]string{}
// 	purchase, err := self.mdls.Purchase().PendingPurchase(ctx, clnt.Id())
// 	if err != nil && !errors.Is(err, sql.ErrNoRows) {
// 		log.Println("Error getting purchase", err)
// 		self.Error(w, r, err)
// 		return
// 	}

// 	items := []mdls.IPurchaseItem{}
// 	if err == nil {
// 		items, err = purchase.PurchaseItems(ctx)
// 		if err != nil {
// 			log.Println("Error getting purchase items")
// 			self.Error(w, r, err)
// 			return
// 		}
// 	}

// 	for _, method := range self.paymentsMgr.Options(clnt) {
// 		route := router.FindRoute(routenames.RoutePaymentSelected)
// 		url, err := route.URL("uuid", method.Uuid())
// 		if err != nil {
// 			self.Error(w, r, err)
// 			return
// 		}
// 		methods[method.IOption().Name()] = url.EscapedPath() + "?" + r.URL.Query().Encode()
// 	}

// 	log.Println("Items:", items)
// 	log.Println("Methods: ", methods)

// 	cancelUrl, err := router.UrlForRoute(routenames.RoutePaymentCancel)
// 	if err != nil {
// 		self.Error(w, r, err)
// 		return
// 	}

// 	data := map[string]interface{}{
// 		"methods":   methods,
// 		"purchase":  purchase,
// 		"items":     items,
// 		"cancelUrl": cancelUrl,
// 	}

// 	self.capi.HttpApi().Respond().AdminView(w, r, "payments/payment-options.html", data)
// }

// func (self *PaymentsCtrl) PaymentOptionSelected(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	clnt, err := helpers.CurrentClient(r)
// 	if err != nil {
// 		self.Error(w, r, err)
// 		return
// 	}

// 	params := mux.Vars(r)
// 	method := self.paymentsMgr.FindByUuid(clnt, params["uuid"])
// 	if method == nil {
// 		self.Error(w, r, errors.New("invalid payment method"))
// 		return
// 	}

// 	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
// 	if err != nil {
// 		log.Println("Error make db tx.")
// 		self.Error(w, r, err)
// 		return
// 	}
// 	defer tx.Rollback()

// 	if pur, err := self.mdls.Purchase().PendingPurchaseTx(tx, ctx, clnt.Id()); err == nil {
// 		method.IOption().PaymentHandler(w, r, clnt, pur)
// 		return
// 	}

// 	req, err := payments.ParsePurchaseReq(r)
// 	if err != nil {
// 		log.Println("Error parse purchase request.")
// 		self.Error(w, r, err)
// 		return
// 	}

// 	pur, err := self.mdls.Purchase().CreateTx(tx, ctx, clnt.Id(), req.VarPrice, req.CallbackUrl)
// 	if err != nil {
// 		log.Println("Error creating purchase.")
// 		self.Error(w, r, err)
// 		return
// 	}

// 	for _, pitem := range req.Items {
// 		_, err := self.mdls.PurchaseItem().CreateTx(tx, ctx, pur.Id(), pitem.Sku, pitem.Name, pitem.Description, pitem.Price)
// 		if err != nil {
// 			log.Println("Error creating purchase item.")
// 			self.Error(w, r, err)
// 			return
// 		}
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		log.Println("Error committing transaction.")
// 		self.Error(w, r, err)
// 		return
// 	}

// 	method.IOption().PaymentHandler(w, r, clnt, pur)
// }

// func (self *PaymentsCtrl) CancelPurchase(w http.ResponseWriter, r *http.Request) {
// 	clnt, err := helpers.CurrentClient(r)
// 	if err != nil {
// 		self.Error(w, r, err)
// 		return
// 	}

// 	ctx := r.Context()
// 	pur, err := self.mdls.Purchase().PendingPurchase(ctx, clnt.Id())
// 	if err != nil {
// 		self.Error(w, r, err)
// 		return
// 	}

// 	err = pur.Cancel(ctx)
// 	if err != nil {
// 		self.Error(w, r, err)
// 		return
// 	}

// 	flash.SetFlashMsg(w, flash.Info, "Purchase has been cancelled.")
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

// func (self *PaymentsCtrl) Error(w http.ResponseWriter, r *http.Request, err error) {
// 	self.errR.Redirect(w, r, err)
// }
