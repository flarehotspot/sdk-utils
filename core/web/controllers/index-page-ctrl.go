package controllers

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/config/themecfg"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/plugins"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"
	themes "github.com/flarehotspot/core/sdk/api/themes"
	"github.com/flarehotspot/core/utils/assets"
	"github.com/flarehotspot/core/web/response"
)

func NewIndexPageCtrl(g *globals.CoreGlobals) IndexPageCtrl {
	return IndexPageCtrl{g}
}

type IndexPageCtrl struct {
	g *globals.CoreGlobals
}

func (c *IndexPageCtrl) PortalIndex(w http.ResponseWriter, r *http.Request) {
	themePkg := themecfg.Read().Portal
	themePlugin, ok := c.g.PluginMgr.FindByPkg(themePkg)
	if !ok {
		http.Error(w, "Invalid portal theme", 500)
		return
	}

	themesApi := themePlugin.ThemesApi().(*plugins.ThemesApi)

	portalRoutes := c.portalRoutes()

	c.render(w, r, themePlugin, portalRoutes, themesApi.GetPortalThemeAssets())
}

func (c *IndexPageCtrl) AdminIndex(w http.ResponseWriter, r *http.Request) {
	themePkg := themecfg.Read().Admin
	themePlugin, ok := c.g.PluginMgr.FindByPkg(themePkg)
	if !ok {
		http.Error(w, "Invalid admin theme", 500)
		return
	}

	themesApi := themePlugin.ThemesApi().(*plugins.ThemesApi)
	adminRoutes := c.adminRoutes(themesApi)
	c.render(w, r, themePlugin, adminRoutes, themesApi.GetAdminThemeAssets())
}

func (c *IndexPageCtrl) MainJs(w http.ResponseWriter, r *http.Request) {
	mainjs := filepath.Join(c.g.CoreApi.Resource("views/scripts/main.tpl.js"))
	helpers := c.g.CoreApi.HttpApi().Helpers()
	w.Header().Set("Content-Type", "application/javascript")
	response.Text(w, mainjs, helpers, nil)
}

func (c *IndexPageCtrl) adminRoutes(themesApi *plugins.ThemesApi) []map[string]any {
	routes := []*plugins.VueRouteComponent{}
	for _, p := range c.g.PluginMgr.All() {
		vueR := p.HttpApi().VueRouter().(*plugins.VueRouterApi)
		adminRoutes := vueR.GetAdminRoutes()
		routes = append(routes, adminRoutes...)
	}

	children := []map[string]any{}
	for _, r := range routes {
		children = append(children, map[string]any{
			"path":      r.VueRoutePath,
			"name":      r.VueRouteName,
			"component": r.HttpComponentFullPath,
			"meta": map[string]any{
				"data_path": r.HttpDataFullPath,
			},
		})
	}

	children = append(children, map[string]any{
		"path":      "/",
		"name":      "index",
		"component": themesApi.AdminIndexComponentFullPath,
	})

	routesMap := []map[string]any{
		{
			"path":      "/",
			"name":      "theme-layout",
			"component": themesApi.AdminLayoutComponentFullPath,
			"children":  children,
			"meta": map[string]any{
				"requireAuth": true,
			},
		},
		{
			"path":      "/login",
			"name":      "login",
			"component": themesApi.AdminLoginComponentFullPath,
			"meta": map[string]any{
				"requireNoAuth": true,
			},
		},
	}

	return routesMap
}

func (c *IndexPageCtrl) portalRoutes() []*plugins.VueRouteComponent {
	routes := []*plugins.VueRouteComponent{}
	for _, p := range c.g.PluginMgr.All() {
		vueR := p.HttpApi().VueRouter().(*plugins.VueRouterApi)
		portalRoutes := vueR.GetPortalRoutes()
		routes = append(routes, portalRoutes...)
	}
	return routes
}

func (c *IndexPageCtrl) render(w http.ResponseWriter, r *http.Request, themePlugin plugin.IPluginApi, routes any, themeAssets themes.ThemeAssets) {

	routesJson, err := json.Marshal(routes)
	if err != nil {
		response.ErrorHtml(w, err.Error())
		return
	}

	routesData := map[string]any{"Routes": string(routesJson)}

	jsFiles := []assets.AssetWithData{
		{File: c.g.CoreApi.Resource("assets/libs/toastify-1.12.0.min.js")},
		{File: c.g.CoreApi.Resource("assets/libs/basic-http-1.0.0.js")},
		{File: c.g.CoreApi.Resource("assets/libs/promise-polyfill.min.js")},
		{File: c.g.CoreApi.Resource("assets/libs/event-source.polyfill.min.js")},
		{File: c.g.CoreApi.Resource("assets/libs/vue-2.7.16.min.js")},
		{File: c.g.CoreApi.Resource("assets/libs/vue-router-3.6.5.min.js")},
		{File: c.g.CoreApi.Resource("assets/app/vue-http.js")},
		{File: c.g.CoreApi.Resource("assets/app/require-config.js")},
		{File: c.g.CoreApi.Resource("assets/app/notification.js")},
		{File: c.g.CoreApi.Resource("assets/app/auth.tpl.js")},
		{File: c.g.CoreApi.Resource("assets/app/router.tpl.js"), Data: routesData},
		{File: c.g.CoreApi.Resource("assets/app/flare-view.js")},
	}

	for _, path := range themeAssets.Scripts {
		file := themePlugin.Resource(filepath.Join("assets", path))
		jsFiles = append(jsFiles, assets.AssetWithData{File: file})
	}

	cssFiles := []assets.AssetWithData{}

	cssFiles = append(cssFiles, assets.AssetWithData{File: c.g.CoreApi.Resource("assets/libs/toastify-1.12.0.min.css")})

	for _, path := range themeAssets.Styles {
		file := themePlugin.Resource(filepath.Join("assets", path))
		cssFiles = append(cssFiles, assets.AssetWithData{File: file})
	}

	jsBundle, err := c.g.CoreApi.Utl.BundleAssetsWithHelper(w, r, jsFiles...)
	if err != nil {
		response.ErrorHtml(w, err.Error())
		return
	}

	cssBundle, err := c.g.CoreApi.Utl.BundleAssetsWithHelper(w, r, cssFiles...)
	if err != nil {
		response.ErrorHtml(w, err.Error())
		return
	}

	vdata := map[string]any{
		"VendorScripts": jsBundle.PublicPath,
		"VendorStyles":  cssBundle.PublicPath,
	}

	api := c.g.CoreApi
	api.HttpApi().HttpResponse().View(w, r, "index.html", vdata)
}
