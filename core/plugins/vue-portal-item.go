package plugins

import (
	"net/http"

	"github.com/flarehotspot/core/sdk/api/http"
)

func NewVuePortalItem(api *PluginApi, r *http.Request, nav sdkhttp.VuePortalItem) sdkhttp.PortalItem {
	var routePath, routeName string
	routePath = sdkhttp.VueNotFoundPath

	vueRouter := api.HttpApi().VueRouter().(*VueRouterApi)
	if route, ok := vueRouter.FindVueRoute(nav.RouteName); ok {
		routePath = route.VueRoutePath
		routeName = route.VueRouteName
	}

	return sdkhttp.PortalItem{
		Label:        api.Utils().Translate("label", nav.TranslateLabel),
		IconUri:      api.HttpAPI.Helpers().AssetPath(nav.IconPath),
		VueRouteName: routeName,
		VueRoutePath: routePath,
	}
}
