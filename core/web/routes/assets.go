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
	assetsCtrl := controllers.NewAssetsCtrl(g)

	router.RootRouter.HandleFunc("/favicon.ico", assetsCtrl.GetFavicon)

	router.AssetsRouter.
		HandleFunc("/with-helper/{pkg}/{version}/{path:.*}", assetsCtrl.AssetWithHelpers).
		Methods("GET").
		Name(routenames.AssetWithHelpers)

	allPlugins := g.PluginMgr.All()
	for _, p := range allPlugins {
		assetsDir := filepath.Join(p.Resource("assets"))
		fs := http.FileServer(http.Dir(assetsDir))
		prefix := p.HttpApi().Helpers().AssetPath("")
		fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
		router.RootRouter.PathPrefix(prefix).Handler(fileserver)
	}

	publicDir := paths.PublicDir
	fs := http.FileServer(http.Dir(publicDir))
	prefix := "/public"
	fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
	router.RootRouter.PathPrefix(prefix).Handler(fileserver)
}

func CoreAssets(g *globals.CoreGlobals) {
	assetsDir := g.CoreApi.Resource("assets")
	fs := http.FileServer(http.Dir(assetsDir))
	prefix := g.CoreApi.HttpApi().Helpers().AssetPath("")
	fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
	router.RootRouter.PathPrefix(prefix).Handler(fileserver)
}
