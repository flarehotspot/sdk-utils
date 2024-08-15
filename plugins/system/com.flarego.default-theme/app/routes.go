package app

import (
	"com.flarego.default-theme/app/controllers"
	sdkhttp "sdk/api/http"
	sdkplugin "sdk/api/plugin"
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
	pluginRouter.Group("/auth", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Post("/login", controllers.LoginCtrl(api)).Name(RouteNameLogin)
		subrouter.Post("/logout", controllers.LogoutCtrl(api)).Name(RouteNameLogout)
	})

	adminRouter := api.Http().HttpRouter().AdminRouter()
	adminRouter.Get("/navs", controllers.GetAdminNavs(api)).Name(RouteAdminNavs)

}
