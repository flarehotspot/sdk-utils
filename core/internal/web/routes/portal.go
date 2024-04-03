package routes

import (
	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/controllers"
	"github.com/flarehotspot/core/internal/web/router"
	routenames "github.com/flarehotspot/core/internal/web/routes/names"
	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
)

func PortalRoutes(g *plugins.CoreGlobals) {
	rootR := router.RootRouter
	portalR := g.CoreAPI.HttpAPI.HttpRouter().PluginRouter()

	portalIndexCtrl := controllers.PortalIndexPage(g)
	portalSseCtrl := controllers.PortalSseHandler(g)
	portalItemsCtrl := controllers.PortalItemsHandler(g)

	rootR.Handle("/", portalIndexCtrl).Methods("GET").Name(routenames.RoutePortalIndex)
	portalR.Get("/events", portalSseCtrl).Name(routenames.RoutePortalSse)
	portalR.Get("/nav/items", portalItemsCtrl).Name(routenames.PortalItems)

	g.CoreAPI.HttpAPI.VueRouter().RegisterPortalRoutes(sdkhttp.VuePortalRoute{
		RouteName: "portal.sample.2",
		RoutePath: "/portal/sample",
        Component: "portal/Sample.vue",
	})

	g.CoreAPI.HttpAPI.VueRouter().PortalItemsFunc(func(clnt sdkconnmgr.ClientDevice) []sdkhttp.VuePortalItem {
		return []sdkhttp.VuePortalItem{
			{
                RouteName: "portal.sample.2",
                Label: "Sample Vue Component",
            },
		}
	})
}
