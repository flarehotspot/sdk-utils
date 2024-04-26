package plugins

import (
	"net/http"

	"github.com/flarehotspot/sdk/api/themes"
	"github.com/flarehotspot/sdk/utils/fs"
)

func NewThemesApi(api *PluginApi) *ThemesApi {
	return &ThemesApi{api: api}
}

type ThemesApi struct {
	api         *PluginApi
	AdminTheme  *sdkthemes.AdminTheme
	PortalTheme *sdkthemes.PortalTheme

	AdminLayoutRoute    *VueRouteComponent
	AdminDashboardRoute *VueRouteComponent
	AdminLoginRoute     *VueRouteComponent

	PortalLayoutRoute *VueRouteComponent
	PortalIndexRoute  *VueRouteComponent
}

func (self *ThemesApi) NewAdminTheme(theme sdkthemes.AdminTheme) {
	layoutComp := NewVueRouteComponent(self.api, theme.LayoutComponent.RouteName, "/theme/layout", theme.LayoutComponent.Component, nil, nil)

	loginComp := NewVueRouteComponent(self.api, theme.LoginComponent.RouteName, "/theme/login", theme.LoginComponent.Component, nil, nil)

	dashComp := NewVueRouteComponent(self.api, theme.DashboardComponent.RouteName, "/theme/dashboard", theme.DashboardComponent.Component, nil, nil)

	self.AdminLayoutRoute = layoutComp
	self.AdminDashboardRoute = dashComp
	self.AdminLoginRoute = loginComp
	self.api.HttpAPI.vueRouter.AddAdminRoutes(dashComp)
	self.api.HttpAPI.vueRouter.SetLoginRoute(loginComp)
	self.AdminTheme = &theme
}

func (self *ThemesApi) NewPortalTheme(theme sdkthemes.PortalTheme) {
	layoutComp := NewVueRouteComponent(self.api, theme.LayoutComponent.RouteName, "/theme/layout", theme.LayoutComponent.Component, nil, nil)

	indexComp := NewVueRouteComponent(self.api, theme.IndexComponent.RouteName, "/theme/index", theme.IndexComponent.Component, nil, nil)

	self.PortalLayoutRoute = layoutComp
	self.PortalIndexRoute = indexComp
	self.api.HttpAPI.vueRouter.AddPortalRoutes(indexComp)
	self.PortalTheme = &theme
}

func (self *ThemesApi) GetAdminThemeAssets() sdkthemes.ThemeAssets {
	assets := sdkthemes.ThemeAssets{Scripts: []string{}, Styles: []string{}}
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

func (self *ThemesApi) GetPortalThemeAssets() sdkthemes.ThemeAssets {
	assets := sdkthemes.ThemeAssets{Scripts: []string{}, Styles: []string{}}
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
		clnt, err := self.api.HttpAPI.GetClientDevice(r)
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		items := self.api.HttpAPI.GetPortalItems(clnt)
		res.Json(w, items, http.StatusOK)
	}
}
