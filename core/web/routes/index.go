package routes

import (
	"net/http"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/controllers/adminctrl"
	"github.com/flarehotspot/core/web/controllers/portalctrl"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes/names"
)

func IndexRoutes(g *globals.CoreGlobals) {
	rootR := router.RootRouter
	deviceMw := middlewares.DeviceMiddleware(g.Db, g.ClientRegister)
	portalCtrl := portalctrl.NewPortalCtrl(g)
	portalIndexCtrl := deviceMw(http.HandlerFunc(portalCtrl.IndexPage))
	rootR.Handle("/", portalIndexCtrl).Methods("GET").Name(names.RoutePortalIndex)

	adminCtrl := adminctrl.NewAdminCtrl(g)
	rootR.HandleFunc("/admin", adminCtrl.AdminIndex)

	// portal assets subpath
	// assetsR := router.AssetsApiRouterV1
	// assetsR.HandleFunc("/"+g.CoreApi.Pkg()+"/portal/main.js", portalAssetsCtrl.MainJs).Methods("GET")
	// assetsR.HandleFunc("/{pkg}/portal/helpers.js", portalAssetsCtrl.HelpersJs(g)).Methods("GET")

	// portal api subpath
	// portalApiRouter := router.ApiRouterV1.PathPrefix("/portal").Subrouter()
	// portalApiRouter.Use(deviceMiddleware)

	// portalApiCtrl := apiv1.NewPortalApiCtrl(g)

	// portalApiRouter.HandleFunc("/navs", portalApiCtrl.PortalNavs).Methods("GET")
}
