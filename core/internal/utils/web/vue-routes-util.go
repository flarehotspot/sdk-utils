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
			Path:      string(r.VueRoutePath),
			Name:      r.VueRouteName,
			Component: r.HttpComponentPath,
		})
	}

	data.LayoutComponent = ThemeComponent{
		Path:      themeApi.PortalLayoutRoute.Path,
		Name:      themeApi.PortalLayoutRoute.Name,
		Component: themeApi.PortalLayoutRoute.HttpComponentPath,
	}

	data.IndexComponent = ThemeComponent{
		Path:      themeApi.PortalIndexRoute.Path,
		Name:      themeApi.PortalIndexRoute.Name,
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
			Path:      string(r.VueRoutePath),
			Name:      r.VueRouteName,
			Component: r.HttpComponentPath,
			RouteMeta: RouteMeta{
				RequireAuth: true,
			},
		})
	}

	data.LayoutComponent = ThemeComponent{
		Path:      themeApi.AdminLayoutRoute.Path,
		Name:      themeApi.AdminLayoutRoute.Name,
		Component: themeApi.AdminLayoutRoute.HttpComponentPath,
	}

	data.DashboardComponent = ThemeComponent{
		Path:      themeApi.AdminDashboardRoute.Path,
		Name:      themeApi.AdminDashboardRoute.Name,
		Component: themeApi.AdminDashboardRoute.HttpComponentPath,
	}

	data.LoginComponent = ThemeComponent{
		Path:      themeApi.AdminLoginRoute.Path,
		Name:      themeApi.AdminLoginRoute.Name,
		Component: themeApi.AdminLoginRoute.HttpComponentPath,
	}

	return data, nil
}
