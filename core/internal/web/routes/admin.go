package routes

import (
	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/controllers"
	"github.com/flarehotspot/core/internal/web/router"
	"github.com/flarehotspot/core/internal/web/routes/names"
)

func AdminRoutes(g *plugins.CoreGlobals) {
	rootR := router.RootRouter
	adminR := g.CoreAPI.HttpAPI.HttpRouter().AdminRouter()

	adminIndexCtrl := controllers.AdminIndexPage(g)
	adminSseCtrl := controllers.AdminSseHandler(g)

	rootR.Handle("/admin", adminIndexCtrl).Methods("GET").Name(routenames.RouteAdminIndex)
	adminR.Get("/events", adminSseCtrl).Name(routenames.RouteAdminSse)
}
