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

	AdminLayoutComponentFullPath string
	AdminLayoutDataFullPath      string

	AdminLoginComponentFullPath string
	AdminLoginDataFullPath      string

	AdminDashVuePath string

	PortalLayoutComponentFullPath string
	PortalIndexComponentFullPath  string
}

func (t *ThemesApi) NewAdminTheme(theme themes.AdminTheme) {
	adminRouter := t.api.HttpAPI.httpRouter.adminRouter.mux
	router := t.api.HttpAPI.httpRouter.pluginRouter.mux

	layoutComp := NewVueRouteComponent(t.api, theme.LayoutComponent.RouteName, "/theme/layout", theme.LayoutComponent.HandlerFunc, theme.LayoutComponent.ComponentPath, nil, nil)
	layoutComp.MountRoute(router)

	loginComp := NewVueRouteComponent(t.api, theme.LoginComponent.RouteName, "/theme/login", theme.LoginComponent.HandlerFunc, theme.LoginComponent.ComponentPath, nil, nil)
	loginComp.MountRoute(router)

	dashComp := NewVueRouteComponent(t.api, theme.DashboardComponent.RouteName, "/theme/dashboard", theme.DashboardComponent.HandlerFunc, theme.DashboardComponent.ComponentPath, nil, nil)
	dashComp.MountRoute(adminRouter)
	// register dashbord component to admin routes
	t.api.HttpAPI.vueRouter.adminRoutes = append(t.api.HttpAPI.vueRouter.adminRoutes, dashComp)

	t.AdminLayoutComponentFullPath = layoutComp.HttpComponentFullPath
	t.AdminLayoutDataFullPath = layoutComp.HttpDataFullPath

	t.AdminLoginComponentFullPath = loginComp.HttpComponentFullPath
	t.AdminLoginDataFullPath = loginComp.HttpDataFullPath

	t.AdminDashVuePath = dashComp.VueRoutePath
	t.adminTheme = theme
}

func (t *ThemesApi) NewPortalTheme(theme themes.PortalTheme) {
	// t.portalTheme = theme
	// r := router.RootRouter
	// r = r.PathPrefix("/vue/theme/portal/components").Subrouter()
	// r.HandleFunc("/layout.vue", t.GetComponentHandler(theme.LayoutComponent)).Methods("GET").Name(routenames.PortalThemeLayout)
	// r.HandleFunc("/index.vue", t.GetComponentHandler(theme.IndexComponent)).Methods("GET").Name(routenames.PortalThemeIndex)

	// portalLayoutPath, _ := r.Get(routenames.PortalThemeLayout).GetPathTemplate()
	// portalIndexPath, _ := r.Get(routenames.PortalThemeIndex).GetPathTemplate()

	// t.PortalLayoutComponentFullPath = portalLayoutPath
	// t.PortalIndexComponentFullPath = portalIndexPath
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

func (t *ThemesApi) GetThemeRouteComponent(comp themes.ThemeComponent, name string, path string) *VueRouteComponent {
	return NewVueRouteComponent(t.api, name, path, comp.HandlerFunc, comp.ComponentPath, nil, nil)
}

// func (t *ThemesApi) GetComponentHandler(comp themes.ThemeComponent) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		helpers := t.api.HttpApi().Helpers()
// 		compfile := filepath.Join(t.api.Utl.Resource(filepath.Join("components", comp.ComponentPath)))
// 		w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 		response.Text(w, compfile, helpers, nil)
// 	}
// }
