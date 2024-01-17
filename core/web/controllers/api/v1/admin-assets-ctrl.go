package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/flarehotspot/core/config/themecfg"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/plugins"
	routerI "github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/web/router"
	routenames "github.com/flarehotspot/core/web/routes/names"
	"github.com/gorilla/mux"
)

func NewAdminAssetsCtrl(g *globals.CoreGlobals) *AdminAssetsCtrl {
	return &AdminAssetsCtrl{g}
}

type AdminAssetsCtrl struct {
	g *globals.CoreGlobals
}

func (c *AdminAssetsCtrl) MainJs(w http.ResponseWriter, r *http.Request) {
	themePkg := themecfg.Read().Admin
	themePlugin := c.g.PluginMgr.FindByPkg(themePkg)
	themesApi := themePlugin.ThemesApi().(*plugins.ThemesApi)
	adminTheme, ok := themesApi.GetAdminLayoutComponents()
	if !ok {
		http.Error(w, "No admin theme component path defined", 500)
		return
	}

	allPlugins := c.g.PluginMgr.All()
	vueRoutes := []plugins.VueAdminRoute{
		{
			RouteName:     "theme-index",
			RoutePath:     "/",
			ComponentPath: adminTheme.IndexComponentPath,
		},
	}

	for _, p := range allPlugins {
		vueRouter := p.HttpApi().VueRouter().(*plugins.VueRouterApi)
		adminRoutes := vueRouter.GetAdminRoutes(r)
		vueRoutes = append(vueRoutes, adminRoutes...)
	}

	routesJson, err := json.Marshal(vueRoutes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	helperJsURL, _ := router.UrlForRoute(routenames.PortalHelperJs, "pkg", "PKG")
	data := map[string]any{
		"CoreApi":     c.g.CoreApi,
		"Routes":      string(routesJson),
		"HelperJsURL": helperJsURL,
		"Theme":       adminTheme,
	}

	c.g.CoreApi.HttpAPI.Respond().Script(w, r, "views/js/main-admin.tpl.js", data)
}

func (c *AdminAssetsCtrl) HelpersJs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkg, ok := vars["pkg"]
	if !ok {
		http.Error(w, "need to specify plugin package name", 500)
		return
	}

	plugin := c.g.PluginMgr.FindByPkg(pkg)
	if plugin == nil {
		http.Error(w, "invalid plugin package name", 500)
		return
	}

	vueRouter := plugin.HttpApi().VueRouter().(*plugins.VueRouterApi)
	adminRoutes := vueRouter.GetAdminRoutes(r)
	routesJson, err := json.Marshal(adminRoutes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	vdata := map[string]any{
		"CoreApi":      c.g.CoreApi,
		"Plugin":       plugin,
		"Routes":       string(routesJson),
		"NotFoundPath": routerI.NotFoundVuePath,
	}

	c.g.CoreApi.HttpAPI.Respond().Script(w, r, "views/portal/js/helpers.tpl.js", vdata)
}
