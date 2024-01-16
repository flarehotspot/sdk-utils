package apiv1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flarehotspot/core/config/themecfg"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/plugins"
	routerI "github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/web/router"
	routenames "github.com/flarehotspot/core/web/routes/names"
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
	routes := []plugins.VuePortalRoute{}

	for _, p := range allPlugins {
		vueRouter := p.HttpApi().VueRouter().(*plugins.VueRouterApi)
		portalRoutes := vueRouter.GetPortalRoutes(r)
		routes = append(routes, portalRoutes...)
	}

	routesJson, err := json.Marshal(routes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	helperJsURL, _ := router.UrlForRoute(routenames.PortalHelperJs, "pkg", "PKG")
	data := map[string]any{
		"CoreApi":     c.g.CoreApi,
		"Routes":      string(routesJson),
		"HelperJsURL": helperJsURL,
		"Theme": map[string]any{
			"LayoutComponent": themePlugin.HttpApi().AssetPath(portalComponent.ThemeComponentPath),
			"IndexComponent":  themePlugin.HttpApi().AssetPath(portalComponent.IndexComponentPath),
		},
	}

	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	c.g.CoreApi.HttpAPI.Respond().Text(w, r, "views/js/main-portal.tpl.js", data)
}

func (c *PortalAssetsCtrl) HelpersJs(w http.ResponseWriter, r *http.Request) {
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
	routes := vueRouter.GetPortalRoutes(r)
	routesJson, err := json.Marshal(routes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println("CoreApi: ", c.g.CoreApi)

	vdata := map[string]any{
		"CoreApi":      c.g.CoreApi,
		"Plugin":       plugin,
		"Routes":       string(routesJson),
		"NotFoundPath": routerI.NotFoundVuePath,
	}

	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	c.g.CoreApi.HttpAPI.Respond().Text(w, r, "views/js/helpers-v1.tpl.js", vdata)
}
