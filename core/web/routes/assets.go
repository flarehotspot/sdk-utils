package routes

import (
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/globals"
	paths "github.com/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/core/web/controllers"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
	routenames "github.com/flarehotspot/core/web/routes/names"
)

func AssetsRoutes(g *globals.CoreGlobals) {
	cacheMw := middlewares.CacheResponse(365)
	assetsCtrl := controllers.NewAssetsCtrl(g)

	router.RootRouter.Handle("/favicon.ico", cacheMw(http.HandlerFunc(assetsCtrl.GetFavicon)))

	vueR := router.RootRouter.PathPrefix("/vue-dynamic-components").Subrouter()
	vueR.Use(cacheMw)

	vueR.HandleFunc("/with-helper/{pkg}/{version}/{path:.*}", assetsCtrl.AssetWithHelpers).
		Methods("GET").
		Name(routenames.AssetWithHelpers)

	vueR.HandleFunc("/plugin-components/{pkg}/{version}/{path:.*}", assetsCtrl.VueComponent).
		Methods("GET").
		Name(routenames.AssetVueComponent)

	allPlugins := g.PluginMgr.All()
	for _, p := range allPlugins {
		assetsDir := filepath.Join(p.Utils().Resource("assets"))
		fs := http.FileServer(http.Dir(assetsDir))
		prefix := p.HttpApi().Helpers().AssetPath("")
		fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
		router.RootRouter.PathPrefix(prefix).Handler(fileserver)
	}

	assetPathMw := middlewares.AssetPath
	publicDir := paths.PublicDir
	fs := http.FileServer(http.Dir(publicDir))
	prefix := "/public"
	fileserver := cacheMw(assetPathMw(http.StripPrefix(prefix, fs)))
	router.RootRouter.PathPrefix(prefix).Handler(fileserver)
}

func CoreAssets(g *globals.CoreGlobals) {
	assetsDir := g.CoreAPI.Utl.Resource("assets")
	fs := http.FileServer(http.Dir(assetsDir))
	prefix := g.CoreAPI.HttpApi().Helpers().AssetPath("")
	fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
	router.RootRouter.PathPrefix(prefix).Handler(fileserver)
}
