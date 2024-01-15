package adminctrl

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/config/appcfg"
	"github.com/flarehotspot/core/config/themecfg"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/plugins"
	"github.com/flarehotspot/core/sdk/api/http/router"
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
	adminComponent, ok := themesApi.GetPortalComponent()
	if !ok {
		http.Error(w, "No admin theme component path defined", 500)
		return
	}

	allPlugins := c.g.PluginMgr.All()
	vueRoutes := []*plugins.VuePortalRoute{}

	for _, p := range allPlugins {
		vueRouter := p.HttpApi().VueRouter().(*plugins.VueRouter)
		portalRoutes := vueRouter.GetAdminRoutes(r)
		vueRoutes = append(vueRoutes, portalRoutes...)
	}

	routesJson, err := json.Marshal(vueRoutes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	appConfig, err := appcfg.Read()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := map[string]any{
		"CoreApi":       c.g.CoreApi,
		"Routes":        string(routesJson),
		"AssetsVersion": appConfig.AssetsVersion,
		"Theme": map[string]any{
			"LayoutComponent": filepath.Join(themePlugin.Pkg(), adminComponent.ThemeComponentPath),
			"IndexComponent":  filepath.Join(themePlugin.Pkg(), adminComponent.IndexComponentPath),
		},
	}

	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	c.g.CoreApi.HttpAPI.Respond().Text(w, r, "templates/js/main.tpl.js", data)
}

func (c *AdminAssetsCtrl) HelpersJs(g *globals.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		pkg, ok := vars["pkg"]
		if !ok {
			http.Error(w, "need to specify plugin package name", 500)
			return
		}

		plugin := g.PluginMgr.FindByPkg(pkg)
		if plugin == nil {
			http.Error(w, "invalid plugin package name", 500)
			return
		}

		vueRouter := plugin.HttpApi().VueRouter().(*plugins.VueRouter)
		adminRoutes := vueRouter.GetAdminRoutes(r)
		routesJson, err := json.Marshal(adminRoutes)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		vdata := map[string]any{
			"Plugin":       plugin,
			"Routes":       string(routesJson),
			"NotFoundPath": router.NotFoundVuePath,
		}

		w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
		c.g.CoreApi.HttpAPI.Respond().Text(w, r, "templates/portal/js/helpers.tpl.js", vdata)
	}
}
