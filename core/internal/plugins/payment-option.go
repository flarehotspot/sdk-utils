package plugins

import (
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	payments "github.com/flarehotspot/sdk/api/payments"
	plugin "github.com/flarehotspot/sdk/api/plugin"
)

func NewPaymentOpt(api plugin.PluginApi, opt payments.PaymentOpt) PaymentOption {
	uuid := api.Pkg() + "::" + opt.OptName
	vrouter := api.Http().VueRouter().(*VueRouterApi)
	var path, name string
	path = sdkhttp.VueNotFoundPath
	if route, ok := vrouter.FindVueRoute(opt.VueRouteName); ok {
		pairs := []string{}
		for k, v := range opt.RouteParams {
			pairs = append(pairs, k, v)
		}
		path = route.VueRoutePath.URL(pairs...)
		name = route.VueRouteName
	}

	return PaymentOption{api, opt, uuid, name, path}
}

type PaymentOption struct {
	api          plugin.PluginApi
	Opt          payments.PaymentOpt
	UUID         string
	VueRouteName string
	VueRoutePath string
}
