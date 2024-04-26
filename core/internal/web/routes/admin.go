package routes

import (
	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/controllers"
	"github.com/flarehotspot/core/internal/web/controllers/adminctrl"
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

	adminR.Group("/themes", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/index", adminctrl.GetAvailableThemes(g))
		subrouter.Post("/save", adminctrl.SaveThemeSettings(g)).Name(routenames.RouteAdminThemesSave)
	})

	adminR.Group("/logs", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/index", adminctrl.GetLogs(g)).Name(routenames.RouteAdminLogsIndex)
	})

	g.CoreAPI.HttpAPI.VueRouter().RegisterAdminRoutes([]sdkhttp.VueAdminRoute{
		{
			RouteName: "theme-picker",
			RoutePath: "/theme-picker",
			Component: "admin/ThemePicker.vue",
		},
		{
			RouteName: "log-viewer",
			RoutePath: "/log-viewer",
			Component: "admin/LogViewer.vue",
		},
	}...)

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
				RouteName: "log-viewer",
			},
		}
	})
}
