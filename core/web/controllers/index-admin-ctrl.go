package controllers

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/config"
	"github.com/flarehotspot/core/plugins"
	"github.com/flarehotspot/core/utils/assets"
	"github.com/flarehotspot/core/web/response"
)

func AdminIndexPage(g *plugins.CoreGlobals) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg, err := config.ReadThemesConfig()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		themePkg := cfg.Admin
		themePlugin, ok := g.PluginMgr.FindByPkg(themePkg)
		if !ok {
			http.Error(w, "Invalid admin theme", 500)
			return
		}

		themesApi := themePlugin.Themes().(*plugins.ThemesApi)
		routes := g.PluginMgr.Utils().GetAdminRoutes()

		appcfg, err := config.ReadApplicationConfig()
		if err != nil {
			response.ErrorHtml(w, err.Error())
			return
		}

		routesJson, err := json.Marshal(routes)
		if err != nil {
			response.ErrorHtml(w, err.Error())
			return
		}

		loginVueName := themesApi.AdminLoginRoute.VueRouteName
		routesData := map[string]any{
			"Routes":         string(routesJson),
			"LoginRouteName": loginVueName,
		}

		jsFiles := []assets.AssetWithData{
			// libs
			{File: g.CoreAPI.Utl.Resource("assets/libs/nprogress-0.2.0.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/toastify-1.12.0.min.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/promise-polyfill.min.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/event-source.polyfill.min.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/vue-2.7.16.min.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/vue-router-3.6.5.min.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/vue-demi-0.14.6.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/vuelidate-core.min.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/vuelidate-validators.min.js")},

			// app
			{File: g.CoreAPI.Utl.Resource("assets/services/require-config.js")},
			{File: g.CoreAPI.Utl.Resource("assets/services/vue-lazy-load.js")},
			{File: g.CoreAPI.Utl.Resource("assets/services/basic-http.js")},
			{File: g.CoreAPI.Utl.Resource("assets/services/utils.js")},
			{File: g.CoreAPI.Utl.Resource("assets/services/vue-http.js")},
			{File: g.CoreAPI.Utl.Resource("assets/services/notify.js")},
			{File: g.CoreAPI.Utl.Resource("assets/admin/router.js"), Data: routesData},
		}

		adminAssets := themesApi.GetAdminThemeAssets()
		for _, path := range adminAssets.Scripts {
			file := themePlugin.Resource(filepath.Join("assets", path))
			jsFiles = append(jsFiles, assets.AssetWithData{File: file})
		}

		cssFiles := []assets.AssetWithData{
			{File: g.CoreAPI.Utl.Resource("assets/libs/nprogress-0.2.0.css")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/toastify-1.12.0.min.css")},
		}

		for _, path := range adminAssets.Styles {
			file := themePlugin.Resource(filepath.Join("assets", path))
			cssFiles = append(cssFiles, assets.AssetWithData{File: file})
		}

		jsBundle, err := g.CoreAPI.Utl.BundleAssetsWithHelper(w, r, jsFiles...)
		if err != nil {
			response.ErrorHtml(w, err.Error())
			return
		}

		cssBundle, err := g.CoreAPI.Utl.BundleAssetsWithHelper(w, r, cssFiles...)
		if err != nil {
			response.ErrorHtml(w, err.Error())
			return
		}

		vdata := map[string]any{
			"Lang":          appcfg.Lang,
			"ThemesApi":     themesApi,
			"VendorScripts": jsBundle.PublicPath,
			"VendorStyles":  cssBundle.PublicPath,
		}

		api := g.CoreAPI
		api.Http().HttpResponse().View(w, r, "admin/index.html", vdata)
	})
}
