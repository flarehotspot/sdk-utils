package routes

import (
	"net/http"
	"path/filepath"

	"core/internal/plugins"
	"core/internal/web/controllers"
	"core/internal/web/middlewares"
	"core/internal/web/router"
	routenames "core/internal/web/routes/names"

	paths "github.com/flarehotspot/go-utils/paths"
)

func AssetsRoutes(g *plugins.CoreGlobals) {
	cacheMw := middlewares.CacheResponse(365)
	assetsCtrl := controllers.NewAssetsCtrl(g)

	router.RootRouter.Handle("/favicon.ico", cacheMw(http.HandlerFunc(assetsCtrl.GetFavicon)))

	vueR := router.AssetsRouter.PathPrefix("/vue-dynamic-components").Subrouter()
	vueR.Use(cacheMw)

	vueR.HandleFunc("/with-helper/{pkg}/{version}/{path:.*}", assetsCtrl.AssetWithHelpers).
		Methods("GET").
		Name(routenames.RouteAssetWithHelpers)

	vueR.HandleFunc("/plugin-components/{pkg}/{version}/{path:.*}", assetsCtrl.VueComponent).
		Methods("GET").
		Name(routenames.RouteAssetVueComponent)

	allPlugins := g.PluginMgr.All()
	for _, p := range allPlugins {
		assetsDir := filepath.Join(p.Resource("assets"))
		fs := http.FileServer(http.Dir(assetsDir))
		prefix := p.Http().Helpers().AssetPath("")
		fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
		router.RootRouter.PathPrefix(prefix).Handler(fileserver)
	}

	// set public static files
	assetPathMw := middlewares.AssetPath
	publicDir := paths.PublicDir
	fs := http.FileServer(http.Dir(publicDir))
	prefix := "/public"
	fileserver := cacheMw(assetPathMw(http.StripPrefix(prefix, fs)))
	router.RootRouter.PathPrefix(prefix).Handler(fileserver)
}

func CoreAssets(g *plugins.CoreGlobals) {
	assetsDir := g.CoreAPI.Utl.Resource("assets")
	fs := http.FileServer(http.Dir(assetsDir))
	prefix := g.CoreAPI.Http().Helpers().AssetPath("")
	fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
	router.RootRouter.PathPrefix(prefix).Handler(fileserver)
}
