package plugins

import (
	"net/http"

	"github.com/flarehotspot/core/sdk/api/http"
	translate "github.com/flarehotspot/core/sdk/utils/translate"
)

func NewVuePortalItem(api *PluginApi, r *http.Request, nav sdkhttp.VuePortalItem) VuePortalItem {
	vueRouter := api.HttpApi().VueRouter().(*VueRouterApi)
	label := api.Translate(translate.Label, nav.TranslateLabel)
	path := sdkhttp.VueNotFoundPath

	if route, ok := vueRouter.FindPortalRoute(nav.RouteName); ok {
		path = route.HttpDataPath
	}

	return VuePortalItem{
		Label:     label,
		RoutePath: path,
	}
}

type VuePortalItem struct {
	Label     string `json:"label"`
	RoutePath string `json:"path"`
}
