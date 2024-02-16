package plugins

// import (
// 	"net/http"

// 	"github.com/flarehotspot/sdk/api/http/router"
// )

// func NewVueAdminRoute(api *PluginApi, route router.VueAdminRoute) VueAdminRoute {
// 	return VueAdminRoute{
// 		VueRouteName:  VueRouteName(api, route.RouteName),
// 		MuxRouteName:  api.HttpAPI.httpRouter.MuxRouteName(router.PluginRouteName(route.RouteName)),
// 		HttpRoutePath: HttpRoutePath(api, route.RoutePath),
// 		HttpHandler:   route.HandlerFunc,
// 		HttpMethods:   route.Methods,
// 		VueRoutePath:  VueRoutePath(api, route.RoutePath),
// 		ComponentPath: VueComponentPath(api, route.ComponentPath),
// 	}
// }

// type VueAdminRoute struct {
// 	VueRouteName        string              `json:"name"`
// 	MuxRouteName        router.MuxRouteName `json:"muxname"`
// 	HttpRoutePath       string              `json:"-"`
// 	HttpHandler         http.HandlerFunc    `json:"-"`
// 	HttpMethods         []string            `json:"-"`
// 	VueRoutePath        string              `json:"path"`
// 	ComponentPath       string              `json:"component"`
// 	PermissionsRequired []string            `json:"permissions_required"`
// 	PermissionsAnyOf    []string            `json:"permissions_any_of"`
// }

