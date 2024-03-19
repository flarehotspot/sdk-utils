package plugins

import (
	"net/http"
	"path/filepath"

	themes "github.com/flarehotspot/sdk/api/themes"
	sdkfs "github.com/flarehotspot/sdk/utils/fs"
)

func NewThemesApi(api *PluginApi) *ThemesApi {
	return &ThemesApi{api: api}
}

type ThemesApi struct {
	api         *PluginApi
	AdminTheme  *themes.AdminTheme
	PortalTheme *themes.PortalTheme

	AdminLayoutRoute    *VueRouteComponent
	AdminDashboardRoute *VueRouteComponent
	AdminLoginRoute     *VueRouteComponent

	PortalLayoutRoute *VueRouteComponent
	PortalIndexRoute  *VueRouteComponent
}

func (self *ThemesApi) NewAdminTheme(theme themes.AdminTheme) {
	adminRouter := self.api.HttpAPI.httpRouter.adminRouter.mux
	compRouter := self.api.HttpAPI.httpRouter.pluginRouter.mux

	layoutComp := NewVueRouteComponent(self.api, theme.LayoutComponent.RouteName, "/theme/layout", theme.LayoutComponent.HandlerFunc, theme.LayoutComponent.Component, nil, nil)
	layoutComp.MountRoute(compRouter)

	loginComp := NewVueRouteComponent(self.api, theme.LoginComponent.RouteName, "/theme/login", theme.LoginComponent.HandlerFunc, theme.LoginComponent.Component, nil, nil)
	loginComp.MountRoute(compRouter)

	dashComp := NewVueRouteComponent(self.api, theme.DashboardComponent.RouteName, "/theme/dashboard", theme.DashboardComponent.HandlerFunc, theme.DashboardComponent.Component, nil, nil)
	dashComp.MountRoute(adminRouter)

	self.AdminLayoutRoute = layoutComp
	self.AdminDashboardRoute = dashComp
	self.AdminLoginRoute = loginComp
	self.api.HttpAPI.vueRouter.AddAdminRoutes(dashComp)
	self.api.HttpAPI.vueRouter.SetLoginRoute(loginComp)
	self.AdminTheme = &theme
}

func (self *ThemesApi) NewPortalTheme(theme themes.PortalTheme) {
	compRouter := self.api.HttpAPI.httpRouter.pluginRouter.mux.PathPrefix("/portal/vue/components").Subrouter()

	layoutComp := NewVueRouteComponent(self.api, theme.LayoutComponent.RouteName, "/theme/layout", theme.LayoutComponent.HandlerFunc, theme.LayoutComponent.Component, nil, nil)
	layoutComp.MountRoute(compRouter)

	purMw := self.api.HttpAPI.middlewares.PendingPurchaseMw()
	indexComp := NewVueRouteComponent(self.api, theme.IndexComponent.RouteName, "/theme/index", theme.IndexComponent.HandlerFunc, theme.IndexComponent.Component, nil, nil)
	indexComp.WrapperFile = self.api.CoreAPI.Resource(filepath.Join("components", "portal", "IndexWrapper.vue"))
	indexComp.MountRoute(compRouter, purMw)

	self.PortalLayoutRoute = layoutComp
	self.PortalIndexRoute = indexComp
	self.api.HttpAPI.vueRouter.AddPortalRoutes(indexComp)
	self.PortalTheme = &theme
}

func (self *ThemesApi) GetAdminThemeAssets() themes.ThemeAssets {
	assets := themes.ThemeAssets{Scripts: []string{}, Styles: []string{}}
	if self.AdminTheme.ThemeAssets != nil {
		if self.AdminTheme.ThemeAssets.Scripts != nil {
			assets.Scripts = self.AdminTheme.ThemeAssets.Scripts
		}
		if self.AdminTheme.ThemeAssets.Styles != nil {
			assets.Styles = self.AdminTheme.ThemeAssets.Styles
		}
	}
	return assets
}

func (self *ThemesApi) GetPortalThemeAssets() themes.ThemeAssets {
	assets := themes.ThemeAssets{Scripts: []string{}, Styles: []string{}}
	if self.PortalTheme.ThemeAssets != nil {
		if self.PortalTheme.ThemeAssets.Scripts != nil {
			assets.Scripts = self.PortalTheme.ThemeAssets.Scripts
		}
		if self.PortalTheme.ThemeAssets.Styles != nil {
			assets.Styles = self.PortalTheme.ThemeAssets.Styles
		}
	}
	return assets
}

func (self *ThemesApi) GetFormFieldPath(vuefile string) (uri string) {
	file := self.api.Utl.Resource("assets/components/forms/" + vuefile)
	if sdkfs.IsFile(file) {
		return self.api.HttpAPI.Helpers().VueComponentPath("forms/" + vuefile)
	}
	vuepath := "forms/" + string(self.AdminTheme.CssLib) + "/" + vuefile
	return self.api.CoreAPI.HttpAPI.Helpers().VueComponentPath(vuepath)
}

func (self *ThemesApi) PortalItemsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := self.api.CoreAPI.HttpAPI.VueResponse()
		items := self.api.HttpAPI.GetPortalItems(r)
		res.Json(w, items, http.StatusOK)
	}
}
