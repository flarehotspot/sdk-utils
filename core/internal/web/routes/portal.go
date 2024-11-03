package routes

import (
	"core/internal/plugins"
	"core/internal/web/controllers"
	"core/internal/web/router"
)

func PortalRoutes(g *plugins.CoreGlobals) {
	rootR := router.RootRouter
	portalR := g.CoreAPI.HttpAPI.HttpRouter().PluginRouter()
	// pendingPurchaseMw := g.CoreAPI.HttpAPI.Middlewares().PendingPurchase()

	portalIndexCtrl := controllers.PortalIndexPage(g)
	portalSseCtrl := controllers.PortalSseHandler(g)
	// portalItemsCtrl := controllers.PortalItemsHandler(g)

	rootR.Handle("/", portalIndexCtrl).Methods("GET").Name("portal:index")
	portalR.Get("/events", portalSseCtrl).Name("portal:sse")
	// portalR.Get("/nav/items", portalItemsCtrl, pendingPurchaseMw).Name("portal:navs:items")
}
