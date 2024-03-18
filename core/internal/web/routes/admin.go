package routes

import (
	"net/http"

	"github.com/flarehotspot/core/internal/config"
	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/controllers"
	"github.com/flarehotspot/core/internal/web/router"
	routenames "github.com/flarehotspot/core/internal/web/routes/names"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
)

func AdminRoutes(g *plugins.CoreGlobals) {
	api := g.CoreAPI
	rootR := router.RootRouter
	adminR := g.CoreAPI.HttpAPI.HttpRouter().AdminRouter()

	adminIndexCtrl := controllers.AdminIndexPage(g)
	adminSseCtrl := controllers.AdminSseHandler(g)

	rootR.Handle("/admin", adminIndexCtrl).Methods("GET").Name(routenames.RouteAdminIndex)
	adminR.Get("/events", adminSseCtrl).Name(routenames.RouteAdminSse)

	adminR.Group("/themes", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// return the settings view

			allPlugins := api.PluginsMgr().All()

			for _, p := range allPlugins {
				pluginApi := p.(*plugins.PluginApi)
				if pluginApi.ThemesAPI.AdminTheme != nil {
					// plugin has admin theme
				}

				if pluginApi.ThemesAPI.PortalTheme != nil {
					// plugin has portal theme
				}
			}

		}).Name("admin.themes.settings")

		subrouter.Post("/save", func(w http.ResponseWriter, r *http.Request) {
			// save the settings
			cfg := config.ThemesConfig{
				Portal: "com.flarego.xxxx",
				Admin:  "com.flarego.xxxx",
			}
			config.WriteThemesConfig(cfg)
			// check error
		}).Name("admin.themes.save")
	})

}
