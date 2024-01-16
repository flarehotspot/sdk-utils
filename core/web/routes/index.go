package routes

import (
	"net/http"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/controllers"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes/names"
)

func IndexRoutes(g *globals.CoreGlobals) {
	rootR := router.RootRouter
	deviceMw := middlewares.DeviceMiddleware(g.Db, g.ClientRegister)
	indexCtrl := controllers.NewIndexPageCtrl(g)
	portalIndexCtrl := deviceMw(http.HandlerFunc(indexCtrl.PortalIndex))

	rootR.Handle("/", portalIndexCtrl).Methods("GET").Name(routenames.RoutePortalIndex)
	rootR.HandleFunc("/admin", indexCtrl.AdminIndex).Methods("GET").Name(routenames.RouteAdminIndex)

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
