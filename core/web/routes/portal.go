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
	deviceMw := middlewares.DeviceMiddleware(g.Db, g.ClientRegister)
	pendingPurMw := middlewares.PendingPurchaseMw(g.Db, g.Models, g.PaymentsMgr)
	portalCtrl := portalctrl.NewPortalCtrl(g)
	portalAssetsCtrl := portalctrl.NewPortalAssetsCtrl(g)

	r.HandleFunc("/favicon.ico", portalAssetsCtrl.FaviconIco)

	portalIndexCtrl := pendingPurMw(http.HandlerFunc(portalCtrl.IndexPage))
	portalIndexCtrl = deviceMw(portalIndexCtrl)
	r.Handle("/", portalIndexCtrl).Methods("GET").Name(names.RoutePortalIndex)

	// portal assets subpath
	corePkg := g.CoreApi.Pkg()
	assetsR := r.PathPrefix("/assets/{version}").Subrouter()
	assetsR.HandleFunc("/"+corePkg+"/portal/main.js", portalAssetsCtrl.MainJs).Methods("GET")
	assetsR.HandleFunc("/{pkg}/portal/helpers.js", portalAssetsCtrl.HelpersJs(g)).Methods("GET")

	// sub paths
	r = r.PathPrefix("/portal").Subrouter()
	r.Use(deviceMw)

	r.HandleFunc("/session/add", portalCtrl.Test).Methods("GET")
	r.HandleFunc("/test/template", portalCtrl.TestTemplate).Methods("GET")

	// portal api subpath
	portalApiCtrl := portalctrl.NewPortalApiCtrl(g)
	apiR := r.PathPrefix("/api").Subrouter()

	apiR.HandleFunc("/navs", portalApiCtrl.PortalNavs).Methods("GET")
}
