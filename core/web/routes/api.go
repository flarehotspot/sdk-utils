package routes

import (
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/controllers/apictrl"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes/names"
)

func ApiRoutes(g *globals.CoreGlobals) {
	r := router.RootRouter()
	sessCtrl := apictrl.NewSessionsApiCtrl(g)
	clntCtrl := apictrl.NewClientApiCtrl(g)
	navsCtrl := apictrl.NewNavsCtrl(g.PluginMgr)
	sseCtrl := apictrl.NewSseApiCtrl()
	deviceMw := middlewares.DeviceMiddleware(g.Db, g.ClientRegister)

	// portal api subrouter
	r = r.PathPrefix("/api").Subrouter()
	portalR := r.PathPrefix("/portal").Subrouter()
	portalR.Use(deviceMw)
	portalR.HandleFunc("/events", sseCtrl.PortalEvents).Methods("GET")
	portalR.HandleFunc("/client", clntCtrl.ClientData).Methods("GET")
	portalR.HandleFunc("/sessions", sessCtrl.Index).Methods("GET")

	// admin api subrouter
	adminR := r.PathPrefix("/admin").Subrouter()
	adminR.Use(middlewares.AdminAuth)
	adminR.HandleFunc("/events", sseCtrl.AdminEvents).Methods("GET")
	adminR.HandleFunc("/navigation/prefetch.json", navsCtrl.NavSearchJson).Methods("GET").Name(names.RouteApiNavsPrefetch)
}
