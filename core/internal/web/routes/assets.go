package routes

import (
	"net/http"

	"core/internal/plugins"
	webutils "core/internal/utils/web"
	"core/internal/web/controllers"
	"core/internal/web/middlewares"

	sdkutils "github.com/flarehotspot/sdk-utils"
)

func AssetsRoutes(g *plugins.CoreGlobals) {
	cacheMw := middlewares.CacheResponse(365)
	assetsCtrl := controllers.NewAssetsCtrl(g)

	webutils.RootRouter.Handle("/favicon.ico", cacheMw(http.HandlerFunc(assetsCtrl.GetFavicon)))

	allPlugins := g.PluginMgr.All()
	for _, p := range allPlugins {
		assetsDist := p.Resource("assets/dist")
		fs := http.FileServer(http.Dir(assetsDist))
		prefix := p.Http().Helpers().AssetPath("")
		fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
		webutils.RootRouter.PathPrefix(prefix).Handler(fileserver)
	}

	for _, p := range allPlugins {
		resourcesDir := p.Resource("")
		fs := http.FileServer(http.Dir(resourcesDir))
		prefix := p.Http().Helpers().ResourcePath("")
		fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
		webutils.RootRouter.PathPrefix(prefix).Handler(fileserver)
	}

	// set public static files
	assetPathMw := middlewares.AssetPath
	publicDir := sdkutils.PathPublicDir
	fs := http.FileServer(http.Dir(publicDir))
	prefix := "/public"
	fileserver := cacheMw(assetPathMw(http.StripPrefix(prefix, fs)))
	webutils.RootRouter.PathPrefix(prefix).Handler(fileserver)
}

func CoreAssets(g *plugins.CoreGlobals) {
	assetsDir := g.CoreAPI.Utl.Resource("assets")
	fs := http.FileServer(http.Dir(assetsDir))
	prefix := g.CoreAPI.Http().Helpers().AssetPath("")
	fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
	webutils.RootRouter.PathPrefix(prefix).Handler(fileserver)
}
