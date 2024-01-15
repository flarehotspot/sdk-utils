package plugins

import (
	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/sdk/utils/translate"
	"net/http"
)

func NewVueAdminNav(api *PluginApi, r *http.Request, nav *router.VueAdminNav) *VueAdminNav {
	vueRouter := api.HttpAPI.vueRouter
	path := router.NotFoundVuePath
	if route, ok := vueRouter.FindAdminRoute(r, nav.RouteName); ok {
		path = route.RoutePath
	}

	label := api.Translate(translate.Label, nav.TranslateLabel)
	return &VueAdminNav{
		Label: label,
		Path:  path,
	}
}

type VueAdminNav struct {
	Category router.INavCategory
	Label    string `json:"label"`
	Path     string `json:"path"`
}
