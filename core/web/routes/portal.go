package routes

import (
	"net/http"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/controllers/portalctrl"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes/names"
)

func PortalRoutes(g *globals.CoreGlobals) {
	r := router.RootRouter()
	portalCtrl := portalctrl.NewPortalCtrl(g)
	deviceMw := middlewares.DeviceMiddleware(g.Db, g.ClientRegister)
	pendingPurMw := middlewares.PendingPurchaseMw(g.Db, g.Models, g.PaymentsMgr)

	portalIndexCtrl := pendingPurMw(http.HandlerFunc(portalCtrl.IndexPage))
	portalIndexCtrl = deviceMw(portalIndexCtrl)
	r.Handle("/", portalIndexCtrl).Methods("GET").Name(names.RoutePortalIndex)

	// sub paths
	r = r.PathPrefix("/portal").Subrouter()
	r.Use(deviceMw)

	r.HandleFunc("/session/add", portalCtrl.Test).Methods("GET")
	r.HandleFunc("/test/template", portalCtrl.TestTemplate).Methods("GET")
}
