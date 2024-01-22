package plugins

import (
	"fmt"
	"net/http"
	"path/filepath"

	themes "github.com/flarehotspot/core/sdk/api/themes"
	"github.com/flarehotspot/core/web/response"
	"github.com/flarehotspot/core/web/router"
	routenames "github.com/flarehotspot/core/web/routes/names"
)

func NewThemesApi(api *PluginApi) *ThemesApi {
	return &ThemesApi{api: api}
}

type ThemesApi struct {
	api                             *PluginApi
	adminTheme                      themes.AdminTheme
	portalTheme                     themes.PortalTheme
	AdminLayoutComponentFullPath    string
	AdminLoginComponentFullPath     string
	PortalLayoutComponentFullPath   string
	PortalIndexComponentFullPath    string
}

func (t *ThemesApi) NewAdminTheme(theme themes.AdminTheme) {
	t.adminTheme = theme

	loginRouterName := fmt.Sprintf("%s:%s", t.api.Pkg(), routenames.AdminThemeLogin)
	layoutRouteName := fmt.Sprintf("%s:%s", t.api.Pkg(), routenames.AdminThemeLayout)

	r := t.api.HttpApi().HttpRouter().(*HttpRouterApi).pluginRouter.mux
	r = r.PathPrefix("/vue/theme/admin/components").Subrouter()
	r.HandleFunc("/Login.vue", t.GetComponentHandler(theme.LoginComponentPath)).Methods("GET").Name(loginRouterName)
	r.HandleFunc("/Layout.vue", t.GetComponentHandler(theme.LayoutComponent)).Methods("GET").Name(layoutRouteName)

	adminLoginPath, _ := r.Get(loginRouterName).GetPathTemplate()
	adminLayoutPath, _ := r.Get(layoutRouteName).GetPathTemplate()

	t.AdminLoginComponentFullPath = adminLoginPath
	t.AdminLayoutComponentFullPath = adminLayoutPath
}

func (t *ThemesApi) NewPortalTheme(theme themes.PortalTheme) {
	t.portalTheme = theme
	r := router.RootRouter
	r = r.PathPrefix("/vue/theme/portal/components").Subrouter()
	r.HandleFunc("/layout.vue", t.GetComponentHandler(theme.LayoutComponent)).Methods("GET").Name(routenames.PortalThemeLayout)
	r.HandleFunc("/index.vue", t.GetComponentHandler(theme.IndexComponent)).Methods("GET").Name(routenames.PortalThemeIndex)

	portalLayoutPath, _ := r.Get(routenames.PortalThemeLayout).GetPathTemplate()
	portalIndexPath, _ := r.Get(routenames.PortalThemeIndex).GetPathTemplate()

	t.PortalLayoutComponentFullPath = portalLayoutPath
	t.PortalIndexComponentFullPath = portalIndexPath
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

func (t *ThemesApi) GetDashboarComponentHandler() http.HandlerFunc {
	route, ok := t.api.HttpAPI.vueRouter.FindAdminRoute(t.adminTheme.DashboardRoute)
	if !ok {
		return func(w http.ResponseWriter, r *http.Request) {
			response.ErrorJson(w, "Invalid dashboard route: "+t.adminTheme.DashboardRoute)
		}
	}

	return route.GetComponentHandler()
}

func (t *ThemesApi) GetDashboardVueRoute() (*VueRouteComponent, bool) {
	return t.api.HttpAPI.vueRouter.FindAdminRoute(t.adminTheme.DashboardRoute)
}

func (t *ThemesApi) GetComponentHandler(comp themes.ThemeComponent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers := t.api.HttpApi().Helpers()
		compfile := filepath.Join(t.api.Resource(filepath.Join("components", comp.ComponentPath)))
		data := comp.Data
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		response.Text(w, compfile, helpers, data)
	}
}
