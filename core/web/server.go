package web

import (
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes"
)

func SetupBootRoutes(g *globals.CoreGlobals) {
	routes.BootRoutes(g)

	coreDir := filepath.Join(paths.AppDir, "core/resources/public")
	routes.PublicAssets("core", coreDir, router.BootingRrouter())
}

func SetupAllRoutes(g *globals.CoreGlobals) {
	router.AdminRouter().
		Use(middlewares.AdminAuth)

	routes.PortalRoutes(g)
	routes.AuthRoutes(g)
	routes.AdminRoutes(g)
	routes.PaymentRoutes(g)
	routes.ApiRoutes(g)

    // plugins public assets
	rootRouter := router.RootRouter()
	plugins := g.PluginMgr.All()
	for _, plugin := range plugins {
		pkg := plugin.Pkg()
		dir := filepath.Join(plugin.Dir(), "resources/public")
		routes.PublicAssets(pkg, dir, rootRouter)
	}

    // core public assets
	coreDir := filepath.Join(paths.AppDir, "core/resources/public")
	routes.PublicAssets("core", coreDir, rootRouter)

    // root public dir
    routes.PublicAssets("", paths.PublicDir, rootRouter)

	router.RootRouter().NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})
}
