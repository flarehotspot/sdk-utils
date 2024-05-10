package routes

import (
	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/controllers"
	"github.com/flarehotspot/core/internal/web/router"
	routenames "github.com/flarehotspot/core/internal/web/routes/names"
)

func PortalRoutes(g *plugins.CoreGlobals) {
	rootR := router.RootRouter
	portalR := g.CoreAPI.HttpAPI.HttpRouter().PluginRouter()
	pendingPurchaseMw := g.CoreAPI.HttpAPI.Middlewares().PendingPurchase()

	portalIndexCtrl := controllers.PortalIndexPage(g)
	portalSseCtrl := controllers.PortalSseHandler(g)
	portalItemsCtrl := controllers.PortalItemsHandler(g)

	rootR.Handle("/", portalIndexCtrl).Methods("GET").Name(routenames.RoutePortalIndex)
	portalR.Get("/events", portalSseCtrl).Name(routenames.RoutePortalSse)
	portalR.Get("/nav/items", portalItemsCtrl, pendingPurchaseMw).Name(routenames.RoutePortalItems)
}
