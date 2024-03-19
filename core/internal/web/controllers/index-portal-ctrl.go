package controllers

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/internal/config"
	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/utils/assets"
	sse "github.com/flarehotspot/core/internal/utils/sse"
	"github.com/flarehotspot/core/internal/web/response"
	routenames "github.com/flarehotspot/core/internal/web/routes/names"
)

func PortalIndexPage(g *plugins.CoreGlobals) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg, err := config.ReadThemesConfig()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		themePkg := cfg.Portal
		themePlugin, ok := g.PluginMgr.FindByPkg(themePkg)
		if !ok {
			http.Error(w, "Invalid admin theme", 500)
			return
		}

		themesApi := themePlugin.Themes().(*plugins.ThemesApi)
		routes := g.PluginMgr.Utils().GetPortalRoutes()

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

		routesData := map[string]any{
			"Routes": string(routesJson),
		}

		ssePath := g.CoreAPI.HttpAPI.Helpers().UrlForRoute(routenames.RoutePortalSse)

		jsFiles := []assets.AssetWithData{
			// libs
			{File: g.CoreAPI.Utl.Resource("assets/libs/nprogress-0.2.0.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/toastify-1.12.0.min.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/promise-polyfill.min.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/event-source.polyfill.min.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/vue-2.7.16.min.js")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/vue-router-3.6.5.min.js")},

			// app
			{File: g.CoreAPI.Utl.Resource("assets/services/require-config.js")},
			{File: g.CoreAPI.Utl.Resource("assets/services/vue-lazy-load.js")},
			{File: g.CoreAPI.Utl.Resource("assets/services/basic-http.js")},
			{File: g.CoreAPI.Utl.Resource("assets/services/utils.js")},
			{File: g.CoreAPI.Utl.Resource("assets/services/events.js"), Data: ssePath},
			{File: g.CoreAPI.Utl.Resource("assets/services/vue-http.js")},
			{File: g.CoreAPI.Utl.Resource("assets/services/notify.js")},
			{File: g.CoreAPI.Utl.Resource("assets/portal/router.js"), Data: routesData},
		}

		portalAssets := themesApi.GetPortalThemeAssets()
		for _, path := range portalAssets.Scripts {
			file := themePlugin.Resource(filepath.Join("assets", path))
			jsFiles = append(jsFiles, assets.AssetWithData{File: file})
		}

		cssFiles := []assets.AssetWithData{
			{File: g.CoreAPI.Utl.Resource("assets/libs/nprogress-0.2.0.css")},
			{File: g.CoreAPI.Utl.Resource("assets/libs/toastify-1.12.0.min.css")},
		}

		for _, path := range portalAssets.Styles {
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
		api.Http().HttpResponse().PortalView(w, r, "index.html", vdata)
	})
}

func PortalSseHandler(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := sse.NewSocket(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		clnt, err := g.CoreAPI.HttpAPI.GetClientDevice(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sse.AddSocket(clnt.MacAddr(), s)
		s.Listen()
	}
}

func PortalItemsHandler(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api := g.CoreAPI
		items := api.Http().GetPortalItems(r)
		res := api.Http().VueResponse()
		res.Json(w, items, 200)
	}
}
