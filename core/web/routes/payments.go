package routes

import (
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/controllers/paymentsctrl"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes/names"
)

func PaymentRoutes(g *globals.CoreGlobals) {
	ctrl := paymentsctrl.NewPaymentsCtrl(g)
	deviceMw := middlewares.DeviceMiddleware(g.Db, g.ClientRegister)
	r := router.RootRouter().PathPrefix("/payments").Subrouter()

	r.Use(deviceMw)
	r.HandleFunc("/options", ctrl.PaymentOptions).
		Methods("GET").Name(names.RoutePaymentOptions)

	r.HandleFunc("/{uuid}/selected", ctrl.PaymentOptionSelected).
		Methods("GET").Name(names.RoutePaymentSelected)

	r.HandleFunc("/cancel", ctrl.CancelPurchase).
		Methods("GET").Name(names.RoutePaymentCancel)
}
