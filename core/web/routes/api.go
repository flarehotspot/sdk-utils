package routes

import (
	// "net/http"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/controllers/api/v1"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
	routenames "github.com/flarehotspot/core/web/routes/names"
)

func ApiRoutes(g *globals.CoreGlobals) {

	// portal apis
	portalApiRouter := router.PortalApiRouter
	deviceMiddleware := middlewares.DeviceMiddleware(g.Db, g.ClientRegister)
	portalApiCtrl := apiv1.NewPortalApiCtrl(g)
	portalApiRouter.Use(deviceMiddleware)
	portalApiRouter.HandleFunc("/navs", portalApiCtrl.PortalNavs).Methods("GET").Name(routenames.PortalNavItems)

	// auth apis
	// authApiRouter := router.AuthApiRouter
	// adminAuthCtrl := apiv1.NewAdminAuthCtrl(g)
	// isAuthenticatedCtrl := adminAuthMw(http.HandlerFunc(adminAuthCtrl.IsAuthenticated))
	// authApiRouter.HandleFunc("/login", adminAuthCtrl.Login).Methods("POST").Name(routenames.AuthLogin)
	// authApiRouter.HandleFunc("/logout", adminAuthCtrl.Logout).Methods("POST").Name(routenames.AuthLogout)
	// authApiRouter.Handle("/is-authenticated", isAuthenticatedCtrl).Methods("GET").Name(routenames.AuthIsAuthenticated)

	// admin apis
	adminApiRouter := router.AdminApiRouter
	adminApiCtrl := apiv1.NewAdminApiCtrl(g)
	adminApiRouter.HandleFunc("/navs", adminApiCtrl.GetAdminNavs).Methods("GET").Name(routenames.AdminNavs)
}
