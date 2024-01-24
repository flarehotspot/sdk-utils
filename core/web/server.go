package web

import (
	"fmt"
	"net/http"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes"
	"github.com/gorilla/mux"
)

func SetupBootRoutes(g *globals.CoreGlobals) {
	routes.BootRoutes(g)
	routes.CoreAssets(g)
}

func SetupAllRoutes(g *globals.CoreGlobals) {
	router.AdminApiRouter.Use(middlewares.AdminAuth)
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

	router.RootRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		// met, err2 := route.GetMethods()
		fmt.Println(tpl)
		return nil
	})
}
