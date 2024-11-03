package routes

import (
	"net/http"
	"path/filepath"

	"core/internal/plugins"
	"core/internal/web/controllers"
	"core/internal/web/middlewares"
	"core/internal/web/router"

	paths "github.com/flarehotspot/go-utils/paths"
)

func AssetsRoutes(g *plugins.CoreGlobals) {
	cacheMw := middlewares.CacheResponse(365)
	assetsCtrl := controllers.NewAssetsCtrl(g)

	router.RootRouter.Handle("/favicon.ico", cacheMw(http.HandlerFunc(assetsCtrl.GetFavicon)))

	allPlugins := g.PluginMgr.All()
	for _, p := range allPlugins {
		assetsDist := filepath.Join(p.Resource("assets/dist"))
		fs := http.FileServer(http.Dir(assetsDist))
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
