package routes

import (
	"net/http"

	"github.com/flarehotspot/core/plugins"
	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/web/helpers"
	routenames "github.com/flarehotspot/core/web/routes/names"
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
				res.FlashMsg("error", err.Error())
				res.Json(w, nil, http.StatusInternalServerError)
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
