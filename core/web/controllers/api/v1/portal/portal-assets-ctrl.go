package portalApiV1

import (
	"encoding/json"
	"net/http"

	"github.com/flarehotspot/core/config/themecfg"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/plugins"
	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/gorilla/mux"
)

func NewPortalAssetsCtrl(g *globals.CoreGlobals) *PortalAssetsCtrl {
	return &PortalAssetsCtrl{g}
}

type PortalAssetsCtrl struct {
	g *globals.CoreGlobals
}

func (c *PortalAssetsCtrl) MainJs(w http.ResponseWriter, r *http.Request) {
	themePkg := themecfg.Read().Portal
	themePlugin := c.g.PluginMgr.FindByPkg(themePkg)
	themesApi := themePlugin.ThemesApi().(*plugins.ThemesApi)
	portalComponent, ok := themesApi.GetPortalComponent()
	if !ok {
		http.Error(w, "No portal theme component path defined", 500)
		return
	}

	allPlugins := c.g.PluginMgr.All()
	routes := []*plugins.VueRoute{}

	for _, p := range allPlugins {
		vueRouter := p.HttpApi().VueRouter().(*plugins.VueRouter)
		portalRoutes := vueRouter.GetPortalRoutes(r)
		routes = append(routes, portalRoutes...)
	}

	routesJson, err := json.Marshal(routes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := map[string]any{
		"CoreApi":       c.g.CoreApi,
		"Routes":        string(routesJson),
		"Theme": map[string]any{
			"LayoutComponent": themePlugin.HttpApi().AssetPath(portalComponent.ThemeComponentPath),
			"IndexComponent":  themePlugin.HttpApi().AssetPath(portalComponent.IndexComponentPath),
		},
	}

	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	c.g.CoreApi.HttpAPI.Respond().Text(w, r, "templates/js/main.tpl.js", data)
}

func (c *PortalAssetsCtrl) HelpersJs(g *globals.CoreGlobals) http.HandlerFunc {
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
		routes := vueRouter.GetPortalRoutes(r)
		routesJson, err := json.Marshal(routes)
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
		c.g.CoreApi.HttpAPI.Respond().Text(w, r, "templates/js/helpers.tpl.js", vdata)
	}
}
