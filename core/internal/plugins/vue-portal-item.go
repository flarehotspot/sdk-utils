package plugins

import (
	"sdk/api/http"
)

func NewVuePortalItem(api *PluginApi, nav sdkhttp.VuePortalItem) sdkhttp.PortalItem {
	var routePath, routeName string
	routePath = sdkhttp.VueNotFoundPath

	vueRouter := api.Http().VueRouter().(*VueRouterApi)
	if route, ok := vueRouter.FindVueRoute(nav.RouteName); ok {
		pairs := []string{}
		for k, v := range nav.RouteParams {
			pairs = append(pairs, k, v)
		}
		routePath = route.VueRoutePath.URL(pairs...)
		routeName = route.VueRouteName
	}

	return sdkhttp.PortalItem{
		Label:          nav.Label,
		IconUrl:        api.HttpAPI.Helpers().AssetPath(nav.IconPath),
		VueRouteName:   routeName,
		VueRoutePath:   routePath,
		VueRouteParams: nav.RouteParams,
	}
}
