package routes

import (
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
)

func PluginAssets(g *globals.CoreGlobals) {
	allPlugins := g.PluginMgr.All()
	for _, p := range allPlugins {
		assetsDir := filepath.Join(p.Resource("assets"))
		fs := http.FileServer(http.Dir(assetsDir))
		prefix := p.HttpApi().AssetPath("")
		fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
		router.RootRouter.PathPrefix(prefix).Handler(fileserver)
	}
}

func CoreAssets(g *globals.CoreGlobals) {
	p := g.CoreApi
	assetsDir := filepath.Join(p.Resource("assets"))
	fs := http.FileServer(http.Dir(assetsDir))
	prefix := p.HttpApi().AssetPath("")
	fileserver := middlewares.AssetPath(http.StripPrefix(prefix, fs))
	router.RootRouter.PathPrefix(prefix).Handler(fileserver)
}
