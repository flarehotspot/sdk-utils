package controllers

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/config"
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
	cfg, err := config.ReadThemesConfig()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	themePkg := cfg.Portal
	themePlugin, ok := c.g.PluginMgr.FindByPkg(themePkg)
	if !ok {
		http.Error(w, "Invalid portal theme", 500)
		return
	}

	themesApi := themePlugin.ThemesApi().(*plugins.ThemesApi)
	portalRoutes := c.g.PluginMgr.Utils().GetPortalRoutes()
	c.render(w, r, themePlugin, portalRoutes, themesApi.GetPortalThemeAssets())
}

func (c *IndexPageCtrl) AdminIndex(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.ReadThemesConfig()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	themePkg := cfg.Admin
	themePlugin, ok := c.g.PluginMgr.FindByPkg(themePkg)
	if !ok {
		http.Error(w, "Invalid admin theme", 500)
		return
	}

	themesApi := themePlugin.ThemesApi().(*plugins.ThemesApi)
	adminRoutes := c.g.PluginMgr.Utils().GetAdminRoutes()
	c.render(w, r, themePlugin, adminRoutes, themesApi.GetAdminThemeAssets())
}

func (c *IndexPageCtrl) MainJs(w http.ResponseWriter, r *http.Request) {
	mainjs := filepath.Join(c.g.CoreApi.Utl.Resource("views/scripts/main.tpl.js"))
	helpers := c.g.CoreApi.HttpApi().Helpers()
	w.Header().Set("Content-Type", "application/javascript")
	response.Text(w, mainjs, helpers, nil)
}

func (c *IndexPageCtrl) render(w http.ResponseWriter, r *http.Request, themePlugin plugin.IPluginApi, routes any, themeAssets themes.ThemeAssets) {

	routesJson, err := json.Marshal(routes)
	if err != nil {
		response.ErrorHtml(w, err.Error())
		return
	}

	routesData := map[string]any{"Routes": string(routesJson)}

	jsFiles := []assets.AssetWithData{
		{File: c.g.CoreApi.Utl.Resource("assets/libs/nprogress-0.2.0.js")},
		{File: c.g.CoreApi.Utl.Resource("assets/libs/toastify-1.12.0.min.js")},
		{File: c.g.CoreApi.Utl.Resource("assets/libs/basic-http-1.0.0.js")},
		{File: c.g.CoreApi.Utl.Resource("assets/libs/promise-polyfill.min.js")},
		{File: c.g.CoreApi.Utl.Resource("assets/libs/event-source.polyfill.min.js")},
		{File: c.g.CoreApi.Utl.Resource("assets/libs/vue-2.7.16.min.js")},
		{File: c.g.CoreApi.Utl.Resource("assets/libs/vue-router-3.6.5.min.js")},
		{File: c.g.CoreApi.Utl.Resource("assets/app/vue-http.js")},
		{File: c.g.CoreApi.Utl.Resource("assets/app/require-config.js")},
		{File: c.g.CoreApi.Utl.Resource("assets/app/notify.js")},
		{File: c.g.CoreApi.Utl.Resource("assets/app/auth.js")},
		{File: c.g.CoreApi.Utl.Resource("assets/app/router.js"), Data: routesData},
	}

	for _, path := range themeAssets.Scripts {
		file := themePlugin.Utils().Resource(filepath.Join("assets", path))
		jsFiles = append(jsFiles, assets.AssetWithData{File: file})
	}

	cssFiles := []assets.AssetWithData{
        {File: c.g.CoreApi.Utl.Resource("assets/libs/nprogress-0.2.0.css")},
        {File: c.g.CoreApi.Utl.Resource("assets/libs/toastify-1.12.0.min.css")},
    }

	for _, path := range themeAssets.Styles {
		file := themePlugin.Utils().Resource(filepath.Join("assets", path))
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
