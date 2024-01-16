package plugins

import (
	"fmt"
	"path/filepath"

	"github.com/flarehotspot/core/sdk/api/http/router"
)

func VueRouteName(api *PluginApi, name string) string {
	return fmt.Sprintf("%s.%s", api.Pkg(), name)
}

func VueRoutePath(api *PluginApi, path string) string {
	return filepath.Join("/", api.Pkg(), path)
}

func VueComponentPath(api *PluginApi, path string) string {
	return api.HttpApi().AssetPath(path)
}

func NewVuePortalRoute(api *PluginApi, route router.VuePortalRoute) VuePortalRoute {
	return VuePortalRoute{
		RouteName:     VueRouteName(api, route.RouteName),
		RoutePath:     VueRoutePath(api, route.RoutePath),
		ComponentPath: VueComponentPath(api, route.ComponentPath),
	}
}

type VuePortalRoute struct {
	RouteName     string `json:"name"`
	RoutePath     string `json:"path"`
	ComponentPath string `json:"component"`
}
