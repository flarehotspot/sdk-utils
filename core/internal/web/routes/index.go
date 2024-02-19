package routes

import (
	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/controllers"
	"github.com/flarehotspot/core/internal/web/router"
	"github.com/flarehotspot/core/internal/web/routes/names"
)

func IndexRoutes(g *plugins.CoreGlobals) {
	rootR := router.RootRouter
	portalIndexCtrl := controllers.PortalIndexPage(g)
	adminIndexCtrl := controllers.AdminIndexPage(g)

	rootR.Handle("/", portalIndexCtrl).Methods("GET").Name(routenames.RoutePortalIndex)
	rootR.Handle("/admin", adminIndexCtrl).Methods("GET").Name(routenames.RouteAdminIndex)
	// rootR.HandleFunc("/scripts/main-"+g.CoreApi.Version()+".js", indexCtrl.MainJs).Methods("GET").Name(routenames.AssetMainJs)

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
