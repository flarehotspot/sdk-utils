package routes

import (
	"net/http"

	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/helpers"
	routenames "github.com/flarehotspot/core/internal/web/routes/names"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
)

func PaymentRoutes(g *plugins.CoreGlobals) {
	g.CoreAPI.HttpAPI.VueRouter().RegisterPortalRoutes(sdkhttp.VuePortalRoute{
		RouteName: routenames.RoutePaymentOptions,
		RoutePath: "/payments/options",
		Component: "payments/customer/PaymentOptions.vue",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			res := g.CoreAPI.HttpAPI.VueResponse()
			clnt, err := helpers.CurrentClient(g.ClientRegister, r)
			if err != nil {
				res.SendFlashMsg(w, "error", err.Error(), http.StatusInternalServerError)
				return
			}

			methods := map[string]string{}
			for _, opt := range g.PaymentsMgr.Options(clnt) {
				methods[opt.Opt.OptName] = opt.VueRoutePath
			}

			res.Json(w, methods, http.StatusOK)
		},
	})
}
