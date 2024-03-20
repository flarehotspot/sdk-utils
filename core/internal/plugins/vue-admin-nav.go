package plugins

import (
	"github.com/flarehotspot/sdk/api/accounts"
	"github.com/flarehotspot/sdk/api/http"
)

func NewVueAdminNav(api *PluginApi, acct sdkacct.Account, nav sdkhttp.VueAdminNav) (sdkhttp.AdminNavItem, bool) {
	var routename, routepath string
	routepath = sdkhttp.VueNotFoundPath

	if vueRoute, ok := api.HttpAPI.vueRouter.FindVueRoute(nav.RouteName); ok {
		// nav.RouteParams is a map[string]string
		pairs := []string{}
		for k, v := range nav.RouteParams {
			pairs = append(pairs, k, v)
		}
		routename = vueRoute.VueRouteName
		routepath = vueRoute.VueRoutePath.URL(pairs...)
	}

	return sdkhttp.AdminNavItem{
		Category:       nav.Category,
		Label:          nav.Label,
		VueRouteName:   routename,
		VueRoutePath:   routepath,
		VueRouteParams: nav.RouteParams,
	}, true
}
