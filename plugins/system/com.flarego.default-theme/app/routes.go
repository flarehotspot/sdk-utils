package app

import (
	sdkplugin "sdk/api/plugin"

	"com.flarego.default-theme/app/controllers"
)

const (
	RouteNameLogin   = "auth.login"
	RouteNameLogout  = "auth.logout"
	RoutePortalItems = "portal.items"
	RouteAdminNavs   = "admin.navs"
	RoutePayments    = "save.settings"
)

func SetupRoutes(api sdkplugin.PluginApi) {
	pluginRouter := api.Http().HttpRouter().PluginRouter()
	pluginRouter.Get("/test", controllers.IndexCtrl(api)).Name("index")
	// pluginRouter.Group("/auth", func(subrouter sdkhttp.HttpRouterInstance) {
	// 	subrouter.Post("/login", controllers.LoginCtrl(api)).Name(RouteNameLogin)
	// 	subrouter.Post("/logout", controllers.LogoutCtrl(api)).Name(RouteNameLogout)
	// })

	// adminRouter := api.Http().HttpRouter().AdminRouter()
	// adminRouter.Get("/navs", controllers.GetAdminNavs(api)).Name(RouteAdminNavs)

}
