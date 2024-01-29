package plugins

import (
	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	payments "github.com/flarehotspot/core/sdk/api/payments"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"
)

func NewPaymentOpt(api plugin.IPluginApi, opt payments.PaymentOpt) PaymentOption {
	uuid := api.Pkg() + "::" + opt.OptName
	vrouter := api.HttpApi().VueRouter().(*VueRouterApi)
	var path, name string
	path = sdkhttp.VueNotFoundPath
	if route, ok := vrouter.FindVueRoute(opt.VueRouteName); ok {
		path = route.VueRoutePath
		name = route.VueRouteName
	}

	return PaymentOption{api, opt, uuid, name, path}
}

type PaymentOption struct {
	api          plugin.IPluginApi
	Opt          payments.PaymentOpt
	UUID         string
	VueRouteName string
	VueRoutePath string
}
