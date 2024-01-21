package plugins

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/sdk/api/themes"
	"github.com/flarehotspot/core/web/response"
	"github.com/flarehotspot/core/web/router"
	routenames "github.com/flarehotspot/core/web/routes/names"
)

func NewThemesApi(api *PluginApi) *ThemesApi {
	return &ThemesApi{api: api}
}

type ThemesApi struct {
	api                           *PluginApi
	adminTheme                    themes.AdminTheme
	portalTheme                   themes.PortalTheme
	AdminLayoutComponentFullPath  string
	AdminIndexComponentFullPath   string
	AdminLoginComponentFullPath   string
	PortalLayoutComponentFullPath string
	PortalIndexComponentFullPath  string
}

func (t *ThemesApi) AdminThemeComponent(theme themes.AdminTheme) {
	layoutRouteName := fmt.Sprintf("%s:%s", t.api.Pkg(), routenames.AdminThemeLayout)
	indexRouterName := fmt.Sprintf("%s:%s", t.api.Pkg(), routenames.AdminThemeIndex)
	loginRouterName := fmt.Sprintf("%s:%s", t.api.Pkg(), routenames.AdminThemeLogin)

	httpRouter := t.api.HttpApi().HttpRouter().(*HttpRouterApi)
	r := httpRouter.pluginRouter.mux
	r = r.PathPrefix("/vue/theme/admin/components").Subrouter()
	r.HandleFunc("/layout.vue", t.GetComponentHandler(theme.LayoutComponent)).Methods("GET").Name(layoutRouteName)
	r.HandleFunc("/index.vue", t.GetComponentHandler(theme.IndexComponentPath)).Methods("GET").Name(indexRouterName)
	r.HandleFunc("/login.vue", t.GetComponentHandler(theme.LoginComponentPath)).Methods("GET").Name(loginRouterName)

	adminLayoutPath, _ := r.Get(layoutRouteName).GetPathTemplate()
	adminIndexPath, _ := r.Get(indexRouterName).GetPathTemplate()
	adminLoginPath, _ := r.Get(loginRouterName).GetPathTemplate()

	t.AdminLayoutComponentFullPath = adminLayoutPath
	t.AdminIndexComponentFullPath = adminIndexPath
	t.AdminLoginComponentFullPath = adminLoginPath
	t.adminTheme = theme
}

func (t *ThemesApi) PortalThemeComponent(theme themes.PortalTheme) {
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

func (t *ThemesApi) GetComponentHandler(comp themes.ThemeComponent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers := t.api.HttpApi().Helpers()
		compfile := filepath.Join(t.api.Resource(filepath.Join("components", comp.ComponentPath)))
		data := comp.Data
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		response.Text(w, compfile, helpers, data)
	}
}
