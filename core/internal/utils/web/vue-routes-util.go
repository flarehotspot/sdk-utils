package webutil

import (
	"github.com/flarehotspot/core/internal/plugins"
)

type RouteMeta struct {
	RequireAuth bool `json:"requireAuth"`
}

type ChildRoutes struct {
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	Component string    `json:"component"`
	RouteMeta RouteMeta `json:"meta"`
	Redirect  string    `json:"redirect"`
}

type ThemeComponent struct {
	Path      string `json:"path"`
	Name      string `json:"name"`
	Component string `json:"component"`
}

type portalRoutesData struct {
	LayoutComponent ThemeComponent `json:"layout_component"`
	IndexComponent  ThemeComponent `json:"index_component"`
	ChildRoutes     []ChildRoutes  `json:"child_routes"`
}

type adminRoutesData struct {
	LayoutComponent    ThemeComponent `json:"layout_component"`
	DashboardComponent ThemeComponent `json:"dashboard_component"`
	LoginComponent     ThemeComponent `json:"login_component"`
	ChildRoutes        []ChildRoutes  `json:"child_routes"`
}

func GetPortalRoutesData(g *plugins.CoreGlobals, themeApi *plugins.ThemesApi) (portalRoutesData, error) {
	var data portalRoutesData

	routes := []*plugins.VueRouteComponent{}
	for _, p := range g.PluginMgr.All() {
		vueR := p.Http().VueRouter().(*plugins.VueRouterApi)
		portalRoutes := vueR.PortalRoutes
		routes = append(routes, portalRoutes...)
	}

	for _, r := range routes {
		data.ChildRoutes = append(data.ChildRoutes, ChildRoutes{
			Path:      r.VueRoutePath.URL(),
			Name:      r.VueRouteName,
			Component: r.HttpComponentPath,
		})
	}

	data.LayoutComponent = ThemeComponent{
		Path:      themeApi.PortalLayoutRoute.VueRoutePath.URL(),
		Name:      themeApi.PortalLayoutRoute.VueRouteName,
		Component: themeApi.PortalLayoutRoute.HttpComponentPath,
	}

	data.IndexComponent = ThemeComponent{
		Path:      themeApi.PortalIndexRoute.VueRoutePath.URL(),
		Name:      themeApi.PortalIndexRoute.VueRouteName,
		Component: themeApi.PortalIndexRoute.HttpComponentPath,
	}

	return data, nil
}

func GetAdminRoutesData(g *plugins.CoreGlobals, themeApi *plugins.ThemesApi) (adminRoutesData, error) {
	var data adminRoutesData

	routes := []*plugins.VueRouteComponent{}
	for _, p := range g.PluginMgr.All() {
		vueR := p.Http().VueRouter().(*plugins.VueRouterApi)
		adminRoutes := vueR.AdminRoutes
		routes = append(routes, adminRoutes...)
	}

	for _, r := range routes {
		data.ChildRoutes = append(data.ChildRoutes, ChildRoutes{
			Path:      r.VueRoutePath.URL(),
			Name:      r.VueRouteName,
			Component: r.HttpComponentPath,
			RouteMeta: RouteMeta{
				RequireAuth: true,
			},
		})
	}

	data.LayoutComponent = ThemeComponent{
		Path:      themeApi.AdminLayoutRoute.VueRoutePath.URL(),
		Name:      themeApi.AdminLayoutRoute.VueRouteName,
		Component: themeApi.AdminLayoutRoute.HttpComponentPath,
	}

	data.DashboardComponent = ThemeComponent{
		Path:      themeApi.AdminDashboardRoute.VueRoutePath.URL(),
		Name:      themeApi.AdminDashboardRoute.VueRouteName,
		Component: themeApi.AdminDashboardRoute.HttpComponentPath,
	}

	data.LoginComponent = ThemeComponent{
		Path:      themeApi.AdminLoginRoute.VueRoutePath.URL(),
		Name:      themeApi.AdminLoginRoute.VueRouteName,
		Component: themeApi.AdminLoginRoute.HttpComponentPath,
	}

	return data, nil
}
