package plugins

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/flarehotspot/sdk/api/accounts"
	"github.com/flarehotspot/sdk/api/connmgr"
	"github.com/flarehotspot/sdk/api/http"
)

func NewVueRouterApi(api *PluginApi) *VueRouterApi {
	return &VueRouterApi{
		api:          api,
		AdminRoutes:  []*VueRouteComponent{},
		PortalRoutes: []*VueRouteComponent{},
	}
}

type VueRouterApi struct {
	api          *PluginApi
	AdminRoutes  []*VueRouteComponent
	PortalRoutes []*VueRouteComponent
	LoginRoute   *VueRouteComponent
	AdminNavsFn  sdkhttp.VueAdminNavsFunc
	PortalNavsFn sdkhttp.VuePortalItemsFunc
}

func (self *VueRouterApi) AddAdminRoutes(route ...*VueRouteComponent) {
	self.AdminRoutes = append(self.AdminRoutes, route...)
}

func (self *VueRouterApi) AddPortalRoutes(route ...*VueRouteComponent) {
	self.PortalRoutes = append(self.PortalRoutes, route...)
}

func (self *VueRouterApi) SetLoginRoute(route *VueRouteComponent) {
	self.LoginRoute = route
}

func (self *VueRouterApi) RegisterAdminRoutes(routes ...sdkhttp.VueAdminRoute) {
	for _, r := range routes {
		route := NewVueRouteComponent(self.api, r.RouteName, r.RoutePath, r.Component, nil, nil)

		if _, ok := self.FindVueRoute(route.VueRouteName); ok {
			log.Println("Warning: Admin route name \"" + r.RouteName + "\" already exists in admin routes ")
		}

		self.AddAdminRoutes(route)
	}
}

func (self *VueRouterApi) RegisterPortalRoutes(routes ...sdkhttp.VuePortalRoute) {
	for _, r := range routes {
		route := NewVueRouteComponent(self.api, r.RouteName, r.RoutePath, r.Component, nil, nil)

		if _, ok := self.FindVueRoute(route.VueRouteName); ok {
			log.Println("Warning: Admin route name \"" + r.RouteName + "\" already exists in admin routes ")
		}

		self.AddPortalRoutes(route)
	}
}

func (self *VueRouterApi) AdminNavsFunc(fn sdkhttp.VueAdminNavsFunc) {
	self.AdminNavsFn = fn
}

func (self *VueRouterApi) GetAdminNavs(acct sdkacct.Account) []sdkhttp.AdminNavItem {
	navs := []sdkhttp.AdminNavItem{}
	if self.AdminNavsFn != nil {
		for _, nav := range self.AdminNavsFn(acct) {
			adminNav, ok := NewVueAdminNav(self.api, acct, nav)
			if ok {
				navs = append(navs, adminNav)
			}
		}
	}
	return navs
}

func (self *VueRouterApi) PortalItemsFunc(fn sdkhttp.VuePortalItemsFunc) {
	self.PortalNavsFn = fn
}

func (self *VueRouterApi) GetPortalItems(clnt sdkconnmgr.ClientDevice) []sdkhttp.PortalItem {
	navs := []sdkhttp.PortalItem{}

	if self.PortalNavsFn != nil {
		for _, nav := range self.PortalNavsFn(clnt) {
			navs = append(navs, NewVuePortalItem(self.api, nav))
		}
		return navs
	}

	return navs
}

func (self *VueRouterApi) FindVueRoute(name string) (*VueRouteComponent, bool) {
	routeName := self.MakeVueRouteName(name)
	for _, route := range self.AdminRoutes {
		if route.VueRouteName == routeName {
			return route, true
		}
	}

	for _, route := range self.PortalRoutes {
		if route.VueRouteName == routeName {
			return route, true
		}
	}

	if self.LoginRoute != nil && self.LoginRoute.VueRouteName == self.MakeVueRouteName(name) {
		return self.LoginRoute, true
	}

	return nil, false
}

func (self *VueRouterApi) ReloadPortalItems(clnt sdkconnmgr.ClientDevice) {
	items := self.api.HttpAPI.GetPortalItems(clnt)
	clnt.Emit("portal:items:reload", items)
}

func (self *VueRouterApi) VueRouteName(name string) string {
	return self.MakeVueRouteName(name)
}

func (self *VueRouterApi) VueRoutePath(name string, pairs ...string) string {
	var path VueRoutePath
	route, ok := self.api.HttpAPI.vueRouter.FindVueRoute(name)
	if !ok {
		path = sdkhttp.VueNotFoundPath
	}
	path = route.VueRoutePath
	return path.URL(pairs...)
}

func (self *VueRouterApi) VuePkgRoutePath(pkg string, name string, pairs ...string) string {
	otherPkg, ok := self.api.PluginsMgrApi.FindByPkg(pkg)
	if !ok {
		return ""
	}
	return otherPkg.Http().VueRouter().VueRoutePath(name, pairs...)
}

func (self *VueRouterApi) MakeVueRouteName(name string) string {
	name = fmt.Sprintf("%s.%s", self.api.Pkg(), name)
	return name
}

func (self *VueRouterApi) MakeVueRoutePath(p string) VueRoutePath {
	p = path.Join("/", self.api.Pkg(), self.api.Version(), p)
	return VueRoutePath(strings.TrimSuffix(p, "/"))
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
