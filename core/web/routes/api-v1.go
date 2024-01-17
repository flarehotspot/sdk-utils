package routes

import (
	"net/http"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/controllers/api/v1"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
	routenames "github.com/flarehotspot/core/web/routes/names"
)

func ApiRoutesV1(g *globals.CoreGlobals) {

	rootR := router.RootRouter
	assetsR := router.AssetsApiRouterV1
	portalAssetsRouterV1 := assetsR.PathPrefix("/portal").Subrouter()
	adminAssetsRouterV1 := assetsR.PathPrefix("/admin").Subrouter()

	commonApiCtrl := apiv1.NewCommonApiCtrl(g)
	assetsR.HandleFunc("/{pkg}/api.js", commonApiCtrl.ApiJs).Methods("GET").Name(routenames.ApiJs)

	// portal assets
	portalAssetsCtrl := apiv1.NewPortalAssetsCtrl(g)
	rootR.HandleFunc("/api/scripts/main-portal.js", portalAssetsCtrl.MainJs).Methods("GET").Name(routenames.PortalMainJs)
	portalAssetsRouterV1.HandleFunc("/{pkg}/helpers.js", portalAssetsCtrl.HelpersJs).Methods("GET").Name(routenames.PortalHelperJs)

	// admin assets
	adminAssetsCtrl := apiv1.NewAdminAssetsCtrl(g)
	rootR.HandleFunc("/api/scripts/main-admin.js", adminAssetsCtrl.MainJs).Methods("GET").Name(routenames.AdminMainJs)
	adminAssetsRouterV1.HandleFunc("/{pkg}/helpers.js", adminAssetsCtrl.HelpersJs).Methods("GET").Name(routenames.AdminHelperJs)

	// portal apis
	portalApiRouterV1 := router.PortalApiRouterV1
	deviceMiddleware := middlewares.DeviceMiddleware(g.Db, g.ClientRegister)
	portalApiCtrl := apiv1.NewPortalApiCtrl(g)
	portalApiRouterV1.Use(deviceMiddleware)
	portalApiRouterV1.HandleFunc("/navs", portalApiCtrl.PortalNavs).Methods("GET").Name(routenames.PortalNavItems)

	// auth apis
	adminAuthMw := middlewares.AdminAuth
	authApiRouterV1 := router.AuthApiRouterV1
	adminAuthCtrl := apiv1.NewAdminAuthCtrl(g)
	isAuthenticatedCtrl := adminAuthMw(http.HandlerFunc(adminAuthCtrl.IsAuthenticated))
	authApiRouterV1.HandleFunc("/login", adminAuthCtrl.Login).Methods("POST").Name(routenames.AuthLogin)
	authApiRouterV1.HandleFunc("/logout", adminAuthCtrl.Logout).Methods("POST").Name(routenames.AuthLogout)
	authApiRouterV1.Handle("/is-authenticated", isAuthenticatedCtrl).Methods("GET").Name(routenames.AuthIsAuthenticated)

	// admin apis
	adminApiRouterV1 := router.AdminApiRouterV1
	adminApiRouterV1.Use(adminAuthMw)
	adminApiCtrl := apiv1.NewAdminApiCtrl(g)
	adminApiRouterV1.HandleFunc("/navs", adminApiCtrl.GetAdminNavs).Methods("GET").Name(routenames.AdminNavs)
}
