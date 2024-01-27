package plugins

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/flarehotspot/core/sdk/api/http"
)

func NewVueRouterApi(api *PluginApi) *VueRouterApi {
	return &VueRouterApi{
		api:          api,
		adminRoutes:  []*VueRouteComponent{},
		portalRoutes: []*VueRouteComponent{},
	}
}

type VueRouterApi struct {
	api          *PluginApi
	adminRoutes  []*VueRouteComponent
	portalRoutes []*VueRouteComponent
	loginRoute   *VueRouteComponent
	adminNavsFn  sdkhttp.VueAdminNavsFunc
	portalNavsFn sdkhttp.VuePortalItemsFunc
}

func (self *VueRouterApi) AddAdminRoutes(route ...*VueRouteComponent) {
	self.adminRoutes = append(self.adminRoutes, route...)
}

func (self *VueRouterApi) AddPortalRoutes(route ...*VueRouteComponent) {
	self.portalRoutes = append(self.portalRoutes, route...)
}

func (self *VueRouterApi) SetLoginRoute(route *VueRouteComponent) {
	self.loginRoute = route
}

func (self *VueRouterApi) RegisterAdminRoutes(routes ...sdkhttp.VueAdminRoute) {
	dataRouter := self.api.HttpAPI.httpRouter.adminRouter.mux
	for _, r := range routes {
		route := NewVueRouteComponent(self.api, r.RouteName, r.RoutePath, r.HandlerFunc, r.Component, nil, nil)

		if _, ok := self.FindVueRoute(route.VueRouteName); ok {
			log.Println("Warning: Admin route name \"" + r.RouteName + "\" already exists in admin routes ")
		}

		route.MountRoute(dataRouter, r.Middlewares...)
		self.AddAdminRoutes(route)
	}
}

func (self *VueRouterApi) RegisterPortalRoutes(routes ...sdkhttp.VuePortalRoute) {
	dataRouter := self.api.HttpAPI.httpRouter.pluginRouter.mux.PathPrefix("/vue/portal").Subrouter()
	for _, r := range routes {
		route := NewVueRouteComponent(self.api, r.RouteName, r.RoutePath, r.HandlerFunc, r.Component, nil, nil)

		if _, ok := self.FindVueRoute(route.VueRouteName); ok {
			log.Println("Warning: Admin route name \"" + r.RouteName + "\" already exists in admin routes ")
		}

		route.MountRoute(dataRouter, r.Middlewares...)
		self.AddPortalRoutes(route)
	}
}

func (self *VueRouterApi) AdminNavsFunc(fn sdkhttp.VueAdminNavsFunc) {
	self.adminNavsFn = fn
}

func (self *VueRouterApi) GetAdminNavs(r *http.Request) []sdkhttp.AdminNavItem {
	navs := []sdkhttp.AdminNavItem{}
	if self.adminNavsFn != nil {
		for _, nav := range self.adminNavsFn(r) {
			adminNav, ok := NewVueAdminNav(self.api, r, nav)
			if ok {
				navs = append(navs, adminNav)
			}
		}
	}
	return navs
}

func (self *VueRouterApi) PortalItemsFunc(fn sdkhttp.VuePortalItemsFunc) {
	self.portalNavsFn = fn
}

func (self *VueRouterApi) GetPortalItems(r *http.Request) []sdkhttp.PortalItem {
	navs := []sdkhttp.PortalItem{}

	if self.portalNavsFn != nil {
		for _, nav := range self.portalNavsFn(r) {
			navs = append(navs, NewVuePortalItem(self.api, r, nav))
		}

		return navs
	}

	return navs
}

func (self *VueRouterApi) FindVueRoute(name string) (*VueRouteComponent, bool) {
	routeName := self.VueRouteName(name)
	for _, route := range self.adminRoutes {
		if route.VueRouteName == routeName {
			return route, true
		}
	}

	for _, route := range self.portalRoutes {
		if route.VueRouteName == routeName {
			return route, true
		}
	}

	if self.loginRoute != nil && self.loginRoute.VueRouteName == self.VueRouteName(name) {
		return self.loginRoute, true
	}

	return nil, false
}

func (self *VueRouterApi) VueRouteName(name string) string {
	name = fmt.Sprintf("%s.%s", self.api.Pkg(), name)
	return name
}

func (self *VueRouterApi) VueRoutePath(path string) string {
	path = filepath.Join("/", self.api.Pkg(), path)
	return strings.TrimSuffix(path, "/")
}

func (self *VueRouterApi) VuePathToMuxPath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	parts := strings.Split(path, "/")
	routPath := strings.Builder{}
	for _, p := range parts {
		if strings.HasPrefix(p, ":") {
			routPath.WriteString(fmt.Sprintf("{%s}", strings.TrimPrefix(p, ":")))
		} else {
			routPath.WriteString(p)
		}
		routPath.WriteString("/")
	}
	path = routPath.String()
	return strings.TrimSuffix(path, "/")
}
