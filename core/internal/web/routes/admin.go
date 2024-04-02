package routes

import (
	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/controllers"
	"github.com/flarehotspot/core/internal/web/router"
	routenames "github.com/flarehotspot/core/internal/web/routes/names"
	sdkacct "github.com/flarehotspot/sdk/api/accounts"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
)

func AdminRoutes(g *plugins.CoreGlobals) {
	rootR := router.RootRouter
	adminR := g.CoreAPI.HttpAPI.HttpRouter().AdminRouter()

	adminIndexCtrl := controllers.AdminIndexPage(g)
	adminSseCtrl := controllers.AdminSseHandler(g)

	rootR.Handle("/admin", adminIndexCtrl).Methods("GET").Name(routenames.RouteAdminIndex)
	adminR.Get("/events", adminSseCtrl).Name(routenames.RouteAdminSse)
	adminR.Post("/themes", controllers.SaveThemeSettings(g)).Name(routenames.RouteAdminThemes)

	g.CoreAPI.HttpAPI.VueRouter().RegisterAdminRoutes([]sdkhttp.VueAdminRoute{
		{
<<<<<<< HEAD
			RouteName: "admin.welcome",
			RoutePath: "/welcome/:name",
			Component: "admin/Welcome.vue",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				api.LoggerAPI.Info("Handling admin welcome route")
				name := api.Http().MuxVars(r)["name"]
				data := map[string]string{
					"name": name,
				}
				g.CoreAPI.HttpAPI.VueResponse().Json(w, data, 200)
			},
			Middlewares: []func(http.Handler) http.Handler{},
			PermitFn: func(perms []string) bool {
				return true
			},
		},
	}...)

	adminR.Group("/themes", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// return the settings view
=======
			RouteName:   "theme-picker",
			RoutePath:   "/theme-picker",
			HandlerFunc: controllers.GetAvailableThemes(g),
			Component:   "admin/ThemePicker.vue",
		},
		{
			RouteName:   "logger",
			RoutePath:   "/logger",
			HandlerFunc: controllers.GetLogs(g),
			Component:   "admin/LogViewer.vue",
		},
	}...)
>>>>>>> dev/feat-logviewer

	g.CoreAPI.HttpAPI.VueRouter().AdminNavsFunc(func(acct sdkacct.Account) []sdkhttp.VueAdminNav {
		return []sdkhttp.VueAdminNav{
			{
				Category:  sdkhttp.NavCategoryThemes,
				Label:     "Select Theme",
				RouteName: "theme-picker",
			},
			{
				Category:  sdkhttp.NavCategorySystem,
				Label:     "View Logs",
				RouteName: "logger",
			},
		}
	})
}
