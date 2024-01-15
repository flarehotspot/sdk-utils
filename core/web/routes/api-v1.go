package routes

import (
	"github.com/flarehotspot/core/globals"
	adminApiV1 "github.com/flarehotspot/core/web/controllers/api/v1/admin"
	portalApiV1 "github.com/flarehotspot/core/web/controllers/api/v1/portal"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
)

func ApiRoutesV1(g *globals.CoreGlobals) {

	// portal assets
	portalAssetsCtrl := portalApiV1.NewPortalAssetsCtrl(g)
	portalAssetsRouterV1 := router.AssetsApiRouterV1.PathPrefix("/portal").Subrouter()
	portalAssetsRouterV1.HandleFunc("/main.js", portalAssetsCtrl.MainJs).Methods("GET")
	portalAssetsRouterV1.HandleFunc("/{pkg}/helpers.js", portalAssetsCtrl.HelpersJs(g)).Methods("GET")

	// admin assets
	adminAssetsCtrl := adminApiV1.NewAdminAssetsCtrl(g)
	adminAssetsRouterV1 := router.AssetsApiRouterV1.PathPrefix("/admin").Subrouter()
	adminAssetsRouterV1.HandleFunc("/main.js", adminAssetsCtrl.MainJs).Methods("GET")
	adminAssetsRouterV1.HandleFunc("/{pkg}/helpers.js", adminAssetsCtrl.HelpersJs(g)).Methods("GET")

	// portal apis
	portalApiRouterV1 := router.PortalApiRouterV1
	deviceMiddleware := middlewares.DeviceMiddleware(g.Db, g.ClientRegister)
	portalApiRouterV1.Use(deviceMiddleware)

	portalApiCtrl := portalApiV1.NewPortalApiCtrl(g)
	portalApiRouterV1.HandleFunc("/navs", portalApiCtrl.PortalNavs).Methods("GET")

}
