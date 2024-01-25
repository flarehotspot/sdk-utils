package plugins

import (
	themes "github.com/flarehotspot/core/sdk/api/themes"
	// "github.com/flarehotspot/core/web/middlewares"
	// "github.com/flarehotspot/core/web/router"
	// routenames "github.com/flarehotspot/core/web/routes/names"
)

func NewThemesApi(api *PluginApi) *ThemesApi {
	return &ThemesApi{api: api}
}

type ThemesApi struct {
	api         *PluginApi
	adminTheme  themes.AdminTheme
	portalTheme themes.PortalTheme

	AdminLayoutRoute    *VueRouteComponent
	AdminDashboardRoute *VueRouteComponent
	AdminLoginRoute     *VueRouteComponent

	PortalLayoutRoute *VueRouteComponent
	PortalIndexRoute  *VueRouteComponent
}

func (t *ThemesApi) NewAdminTheme(theme themes.AdminTheme) {
	adminRouter := t.api.HttpAPI.httpRouter.adminRouter.mux
	compRouter := t.api.HttpAPI.httpRouter.pluginRouter.mux

	layoutComp := NewVueRouteComponent(t.api, theme.LayoutComponent.RouteName, "/theme/layout", theme.LayoutComponent.HandlerFunc, theme.LayoutComponent.ComponentPath, nil, nil)
	layoutComp.MountRoute(compRouter)

	loginComp := NewVueRouteComponent(t.api, theme.LoginComponent.RouteName, "/theme/login", theme.LoginComponent.HandlerFunc, theme.LoginComponent.ComponentPath, nil, nil)
	loginComp.MountRoute(compRouter)

	dashComp := NewVueRouteComponent(t.api, theme.DashboardComponent.RouteName, "/theme/dashboard", theme.DashboardComponent.HandlerFunc, theme.DashboardComponent.ComponentPath, nil, nil)
	dashComp.MountRoute(adminRouter)

	t.AdminLayoutRoute = layoutComp
	t.AdminDashboardRoute = dashComp
	t.AdminLoginRoute = loginComp
	t.api.HttpAPI.vueRouter.AddAdminRoutes(dashComp)
	t.api.HttpAPI.vueRouter.SetLoginRoute(loginComp)
	t.adminTheme = theme
}

func (t *ThemesApi) NewPortalTheme(theme themes.PortalTheme) {
	compRouter := t.api.HttpAPI.httpRouter.pluginRouter.mux.PathPrefix("/portal/vue/components").Subrouter()

	layoutComp := NewVueRouteComponent(t.api, theme.LayoutComponent.RouteName, "/theme/layout", theme.LayoutComponent.HandlerFunc, theme.LayoutComponent.ComponentPath, nil, nil)
	layoutComp.MountRoute(compRouter)

	indexComp := NewVueRouteComponent(t.api, theme.IndexComponent.RouteName, "/theme/index", theme.IndexComponent.HandlerFunc, theme.IndexComponent.ComponentPath, nil, nil)
	indexComp.MountRoute(compRouter)

	t.PortalLayoutRoute = layoutComp
	t.PortalIndexRoute = indexComp
	t.api.HttpAPI.vueRouter.AddPortalRoutes(indexComp)
	t.portalTheme = theme
}

func (t *ThemesApi) GetAdminThemeAssets() themes.ThemeAssets {
	assets := themes.ThemeAssets{Scripts: []string{}, Styles: []string{}}
	if t.adminTheme.ThemeAssets != nil {
		if t.adminTheme.ThemeAssets.Scripts != nil {
			assets.Scripts = t.adminTheme.ThemeAssets.Scripts
		}
		if t.adminTheme.ThemeAssets.Styles != nil {
			assets.Styles = t.adminTheme.ThemeAssets.Styles
		}
	}
	return assets
}

func (t *ThemesApi) GetPortalThemeAssets() themes.ThemeAssets {
	assets := themes.ThemeAssets{Scripts: []string{}, Styles: []string{}}
	if t.portalTheme.ThemeAssets != nil {
		if t.portalTheme.ThemeAssets.Scripts != nil {
			assets.Scripts = t.portalTheme.ThemeAssets.Scripts
		}
		if t.portalTheme.ThemeAssets.Styles != nil {
			assets.Styles = t.portalTheme.ThemeAssets.Styles
		}
	}
	return assets
}
