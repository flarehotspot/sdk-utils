package plugins

import (
	"github.com/flarehotspot/core/sdk/api/http/router"
	"net/http"
)

func NewVueRouter(api *PluginApi) *VueRouter {
	return &VueRouter{api: api}
}

type VueRouter struct {
	api            *PluginApi
	adminRoutesFn  func(r *http.Request) []*router.VueRoute
	adminNavsFn    func(r *http.Request) []*router.VueAdminNav
	portalRoutesFn func(r *http.Request) []*router.VueRoute
	portalNavsFn   func(r *http.Request) []*router.VuePortalItem
}

func (self *VueRouter) AdminRoutes(fn func(r *http.Request) []*router.VueRoute) {
	self.adminRoutesFn = fn
}

func (self *VueRouter) GetAdminRoutes(r *http.Request) []*VueRoute {
	routes := []*VueRoute{}
	if self.adminRoutesFn != nil {
		for _, route := range self.adminRoutesFn(r) {
			routes = append(routes, NewVueRoute(self.api, route))
		}
	}
	return routes
}

func (self *VueRouter) FindAdminRoute(r *http.Request, name string) (*VueRoute, bool) {
	routeName := VueRouteName(self.api, name)
	for _, route := range self.GetAdminRoutes(r) {
		if route.RouteName == routeName {
			return route, true
		}
	}

	return nil, false
}

func (self *VueRouter) AdminNavs(fn func(r *http.Request) []*router.VueAdminNav) {
	self.adminNavsFn = fn
}

func (self *VueRouter) GetAdminNavs(r *http.Request) []*VueAdminNav {
	navs := []*VueAdminNav{}
	if self.adminNavsFn != nil {
		for _, nav := range self.adminNavsFn(r) {
			navs = append(navs, NewVueAdminNav(self.api, r, nav))
		}
	}

	return navs
}

func (self *VueRouter) PortalRoutes(fn func(r *http.Request) []*router.VueRoute) {
	self.portalRoutesFn = fn
}

func (self *VueRouter) GetPortalRoutes(r *http.Request) []*VueRoute {
	routes := []*VueRoute{}
	if self.portalRoutesFn != nil {
		for _, route := range self.portalRoutesFn(r) {
			routes = append(routes, NewVueRoute(self.api, route))
		}
	}
	return routes
}

func (self *VueRouter) FindPortalRoute(r *http.Request, name string) (*VueRoute, bool) {
	routeName := VueRouteName(self.api, name)

	for _, route := range self.GetPortalRoutes(r) {
		if route.RouteName == routeName {
			return route, true
		}
	}

	return nil, false
}

func (self *VueRouter) PortalItems(fn func(r *http.Request) []*router.VuePortalItem) {
	self.portalNavsFn = fn
}

func (self *VueRouter) GetPortalItems(r *http.Request) []*VuePortalItem {
	navs := []*VuePortalItem{}

	if self.portalNavsFn != nil {
		for _, nav := range self.portalNavsFn(r) {
			navs = append(navs, NewVuePortalItem(self.api, r, nav))
		}

		return navs
	}

	return navs
}
