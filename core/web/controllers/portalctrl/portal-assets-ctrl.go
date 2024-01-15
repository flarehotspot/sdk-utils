package portalctrl

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/flarehotspot/core/config/appcfg"
	"github.com/flarehotspot/core/config/themecfg"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/plugins"
	"github.com/gorilla/mux"
)

func NewPortalAssetsCtrl(g *globals.CoreGlobals) *PortalAssetsCtrl {
	return &PortalAssetsCtrl{g}
}

type PortalAssetsCtrl struct {
	g *globals.CoreGlobals
}

func (c *PortalAssetsCtrl) FaviconIco(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := os.ReadFile(c.g.CoreApi.Resource("assets/images/default-favicon-32x32.png"))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}

func (c *PortalAssetsCtrl) MainJs(w http.ResponseWriter, r *http.Request) {
	themePkg := themecfg.Read().CaptivePortal
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

	appConfig, err := appcfg.Read()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := map[string]any{
		"CoreApi":       c.g.CoreApi,
		"Routes":        string(routesJson),
		"AssetsVersion": appConfig.AssetsVersion,
		"PortalTheme": map[string]any{
			"LayoutComponent": filepath.Join(themePlugin.Pkg(), portalComponent.ThemeComponentPath),
			"IndexComponent":  filepath.Join(themePlugin.Pkg(), portalComponent.IndexComponentPath),
		},
	}

	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	c.g.CoreApi.HttpAPI.Respond().Text(w, r, "templates/portal/js/main.tpl.js", data)
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
			"Plugin": plugin,
			"Routes": string(routesJson),
		}

		w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
		c.g.CoreApi.HttpAPI.Respond().Text(w, r, "templates/portal/js/helpers.tpl.js", vdata)
	}
}
