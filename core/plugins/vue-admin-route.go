package plugins

import (
	"github.com/flarehotspot/core/sdk/api/http/router"
)

func NewVueAdminRoute(api *PluginApi, route *router.VueAdminRoute) *VueAdminRoute {
	return &VueAdminRoute{
		RouteName:     VueRouteName(api, route.RouteName),
		RoutePath:     VueRoutePath(api, route.RoutePath),
		ComponentPath: VueComponentPath(api, route.ComponentPath),
	}
}

type VueAdminRoute struct {
	RouteName           string   `json:"name"`
	RoutePath           string   `json:"path"`
	ComponentPath       string   `json:"component"`
	PermissionsRequired []string `json:"permissions_required"`
	PermissionsAnyOf    []string `json:"permissions_any_of"`
}
