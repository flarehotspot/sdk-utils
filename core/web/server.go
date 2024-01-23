package web

import (
	"net/http"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes"
)

func SetupBootRoutes(g *globals.CoreGlobals) {
	routes.BootRoutes(g)
	routes.CoreAssets(g)
}

func SetupAllRoutes(g *globals.CoreGlobals) {

	routes.IndexRoutes(g)
	routes.AssetsRoutes(g)
	routes.ApiRoutes(g)
	// routes.AuthRoutes(g)
	// routes.AdminRoutes(g)
	// routes.PaymentRoutes(g)
	// routes.ApiRoutes(g)

	router.RootRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})
}
