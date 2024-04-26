package routes

import (
	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/controllers"
	routenames "github.com/flarehotspot/core/internal/web/routes/names"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
)

func PaymentRoutes(g *plugins.CoreGlobals) {

	portalR := g.CoreAPI.HttpAPI.HttpRouter().PluginRouter()
	vueR := g.CoreAPI.HttpAPI.VueRouter()

	portalR.Group("/payments", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/options", controllers.PaymentOptionsCtrl(g)).Name(routenames.RoutePortalPaymentOptions)
	})

	vueR.RegisterPortalRoutes(sdkhttp.VuePortalRoute{
		RouteName: routenames.RoutePaymentOptions,
		RoutePath: "/payments/options",
		Component: "payments/customer/PaymentOptions.vue",
	})
}
