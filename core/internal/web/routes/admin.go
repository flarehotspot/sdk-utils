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

	getAvailableThemes := controllers.RespondJsonThemes(g)
	saveThemesAvailable := controllers.SaveThemeSettings(g)

	adminIndexCtrl := controllers.AdminIndexPage(g)
	adminSseCtrl := controllers.AdminSseHandler(g)

	rootR.Handle("/admin", adminIndexCtrl).Methods("GET").Name(routenames.RouteAdminIndex)
	adminR.Get("/events", adminSseCtrl).Name(routenames.RouteAdminSse)

	adminR.Group("/themes", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/", getAvailableThemes).Name(routenames.AdminThemeSettings)
		subrouter.Post("/save", saveThemesAvailable).Name(routenames.AdminSaveThemeSettings)
	})
	g.CoreAPI.HttpAPI.VueRouter().AdminNavsFunc(func(acct sdkacct.Account) []sdkhttp.VueAdminNav {
		return []sdkhttp.VueAdminNav{
			{
				Category:  sdkhttp.NavCategoryThemes,
				Label:     "Theme Picker",
				RouteName: "settings-theme",
			},
		}
	})
	g.CoreAPI.HttpAPI.VueRouter().RegisterAdminRoutes([]sdkhttp.VueAdminRoute{
		{
			RouteName:   "settings-theme",
			RoutePath:   "/themes",
			HandlerFunc: getAvailableThemes,
			Component:   "ThemeSettings.vue",
		},
	}...)
}
