package web

import (
	"log"
	"net/http"

	"core/internal/plugins"
	"core/internal/web/router"
	"core/internal/web/routes"
)

func SetupBootRoutes(g *plugins.CoreGlobals) {
	routes.AssetsRoutes(g)
	routes.BootRoutes(g)
	routes.CoreAssets(g)
}

func SetupAllRoutes(g *plugins.CoreGlobals) {
	routes.AssetsRoutes(g)
	routes.PortalRoutes(g)
	routes.AdminRoutes(g)
	routes.PaymentRoutes(g)

	router.RootRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Warning: unknown route requested: ", r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
	})
}
