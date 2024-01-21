package plugins

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	routerI "github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/gorilla/mux"
)

func NewVueRouterApi(api *PluginApi) *VueRouterApi {
	return &VueRouterApi{
		api:          api,
		adminRoutes:  []*VueComponentRoute{},
		portalRoutes: []*VueComponentRoute{},
	}
}

type VueRouterApi struct {
	api          *PluginApi
	adminRoutes  []*VueComponentRoute
	portalRoutes []*VueComponentRoute
	adminNavsFn  routerI.VueAdminNavsHandler
	portalNavsFn routerI.VuePortalItemsHandler
}

func (self *VueRouterApi) AdminRoutes(routes []routerI.VueAdminRoute) {
	if routes != nil {
		compRouter := self.api.HttpAPI.httpRouter.pluginRouter.mux.PathPrefix("/vue-route/admin-components").Subrouter()
		dataRouter := self.api.HttpAPI.httpRouter.adminRouter.mux.PathPrefix("/vue-route/admin-data").Subrouter()

		for _, r := range routes {
			route := NewVueComponentRoute(self.api, r.RouteName, r.RoutePath, r.HandlerFn, r.Component, r.DisableCache, true, nil, nil)

			if _, ok := self.FindAdminRoute(route.VueRouteName); ok {
				log.Println("Warning: Admin route name \"" + r.RouteName + "\" already exists in admin routes ")
			}

			compRouter.
				HandleFunc(route.HttpComponentPath, route.GetComponentHandler()).
				Methods("GET").
				Name(string(route.MuxCompRouteName))

			handler := http.HandlerFunc(route.GetDataHandler())

			var handlerFunc http.Handler
			if r.Middlewares != nil {
				for _, m := range r.Middlewares {
					handlerFunc = m(handler)
				}
			} else {
				handlerFunc = handler
			}

			dataRouter.
				Handle(route.HttpDataPath, handlerFunc).
				Methods("GET").
				Name(string(route.MuxDataRouteName))

			router := dataRouter
			compR := router.Get(string(route.MuxCompRouteName))
			dataR := router.Get(string(route.MuxDataRouteName))
			comppath, _ := compR.GetPathTemplate()
			datapath, _ := dataR.GetPathTemplate()
			route.HttpComponentFullPath = comppath
			route.HttpDataFullPath = datapath
			self.adminRoutes = append(self.adminRoutes, route)
		}
	}
}

func (self *VueRouterApi) PortalRoutes(routes []routerI.VuePortalRoute) {
	if routes != nil {

		pluginRouter := self.api.HttpAPI.httpRouter.pluginRouter
		compRouter := pluginRouter.mux.PathPrefix("/vue-route/portal-components").Subrouter()
		dataRouter := pluginRouter.mux.PathPrefix("/vue-route/portal-data").Subrouter()
		for _, r := range routes {
			route := NewVueComponentRoute(self.api, r.RouteName, r.RoutePath, r.HandlerFn, r.Component, r.DisableCache, false, nil, nil)

			if _, ok := self.FindPortalRoute(route.VueRouteName); ok {
				log.Println("Warning: Portal route name \"" + r.RouteName + "\" already exists in portal routes ")
			}

			compRouter.
				HandleFunc(route.HttpComponentPath, route.GetComponentHandler()).
				Methods("GET").
				Name(string(route.MuxCompRouteName))

			handler := http.HandlerFunc(route.GetDataHandler())

			var handlerFunc http.Handler
			if r.Middlewares != nil {
				for _, m := range r.Middlewares {
					handlerFunc = m(handlerFunc)
				}
			} else {
				handlerFunc = handler
			}

			dataRouter.
				HandleFunc(route.HttpDataPath, route.GetDataHandler()).
				Methods("GET").
				Name(string(route.MuxDataRouteName))

			router := compRouter
			compR := router.Get(string(route.MuxCompRouteName))
			dataR := router.Get(string(route.MuxDataRouteName))
			comppath, _ := compR.GetPathTemplate()
			datapath, _ := dataR.GetPathTemplate()
			route.HttpComponentFullPath = comppath
			route.HttpDataFullPath = datapath
			self.portalRoutes = append(self.portalRoutes, route)
		}

	}
}

func (self *VueRouterApi) GetAdminRoutes() []*VueComponentRoute {
	return self.adminRoutes
}

func (self *VueRouterApi) GetPortalRoutes() []*VueComponentRoute {
	return self.portalRoutes
}

func (self *VueRouterApi) FindAdminRoute(vueRouteName string) (*VueComponentRoute, bool) {
	vueR := self.api.HttpAPI.vueRouter
	routeName := vueR.VueRouteName(vueRouteName)
	for _, route := range self.GetAdminRoutes() {
		if route.VueRouteName == routeName {
			return route, true
		}
	}
	return nil, false
}

func (self *VueRouterApi) AdminNavs(fn routerI.VueAdminNavsHandler) {
	self.adminNavsFn = fn
}

func (self *VueRouterApi) GetAdminNavs(r *http.Request) []VueAdminNav {
	navs := []VueAdminNav{}
	if self.adminNavsFn != nil {
		vars := mux.Vars(r)
		for _, nav := range self.adminNavsFn(r, vars) {
			navs = append(navs, NewVueAdminNav(self.api, r, nav))
		}
	}

	return navs
}

func (self *VueRouterApi) FindPortalRoute(name string) (*VueComponentRoute, bool) {
	vueR := self.api.HttpAPI.vueRouter
	routeName := vueR.VueRouteName(name)

	for _, route := range self.GetPortalRoutes() {
		if route.VueRouteName == routeName {
			return route, true
		}
	}

	return nil, false
}

func (self *VueRouterApi) FindVueComponent(name string) (VueComponentRoute, bool) {
	return VueComponentRoute{}, true
}

func (self *VueRouterApi) PortalItems(fn routerI.VuePortalItemsHandler) {
	self.portalNavsFn = fn
}

func (self *VueRouterApi) GetPortalItems(r *http.Request) []VuePortalItem {
	navs := []VuePortalItem{}
	vars := mux.Vars(r)

	if self.portalNavsFn != nil {
		for _, nav := range self.portalNavsFn(r, vars) {
			navs = append(navs, NewVuePortalItem(self.api, r, nav))
		}

		return navs
	}

	return navs
}

func (self *VueRouterApi) FindVueRoute(name string) (*VueComponentRoute, bool) {
	vueR := self.api.HttpAPI.vueRouter
	routeName := vueR.VueRouteName(name)
	for _, route := range self.GetAdminRoutes() {
		if route.VueRouteName == routeName {
			return route, true
		}
	}
	for _, route := range self.GetPortalRoutes() {
		if route.VueRouteName == routeName {
			return route, true
		}
	}
	return nil, false
}

func (self *VueRouterApi) VueRouteName(name string) string {
	name = fmt.Sprintf("%s-%s", self.api.Slug(), name)
	name = strings.ReplaceAll(name, "-", "")
	return name
}

func (self *VueRouterApi) VueRoutePath(path string) string {
	path = filepath.Join("/", self.api.Pkg(), path)
	return strings.TrimSuffix(path, "/")
}
func (self *VueRouterApi) HttpDataPath(path string) string {
	path = self.MuxPathFromVue(path)
	// path = filepath.Join("/data", path)
	return strings.TrimSuffix(path, "/")
}

func (self *VueRouterApi) HttpComponentPath(name string) string {
	name = filepath.Join("/components/", name)
	if !strings.HasSuffix(name, ".vue") {
		name = name + ".vue"
	}
	return name
}

func (self *VueRouterApi) MuxPathFromVue(path string) string {
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
