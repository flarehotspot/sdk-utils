package routes

import (
	"core/internal/plugins"
	"core/internal/web/controllers"
	"core/internal/web/controllers/adminctrl"
	"core/internal/web/router"
	sdkacct "sdk/api/accounts"
	sdkhttp "sdk/api/http"
)

func AdminRoutes(g *plugins.CoreGlobals) {
	rootR := router.RootRouter
	adminR := g.CoreAPI.HttpAPI.HttpRouter().AdminRouter()

	adminIndexCtrl := controllers.AdminIndexPage(g)
	adminSseCtrl := controllers.AdminSseHandler(g)

	rootR.Handle("/admin", adminIndexCtrl).Methods("GET").Name("admin:index")
	adminR.Get("/events", adminSseCtrl).Name("admin:sse")

	adminR.Group("/core", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/fetch", adminctrl.FetchLatestCoreReleaseCtrl(g)).Name("admin:core:fetch")
		subrouter.Get("/current", adminctrl.GetCurrentCoreVersionCtrl(g)).Name("admin:core:current")
		subrouter.Post("/download", adminctrl.DownloadUpdatesCtrl(g)).Name("admin:core:download")
		subrouter.Post("/update", adminctrl.UpateCoreCtrl(g)).Name("admin:core:update")
	})

	adminR.Group("/themes", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/index", adminctrl.GetAvailableThemes(g)).Name("admin:themes:list")
		subrouter.Post("/save", adminctrl.SaveThemeSettings(g)).Name("admin:themes:save")
	})

	adminR.Group("/logs", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/index", adminctrl.GetLogs(g)).Name("admin:logs:index")
		subrouter.Post("/clear", adminctrl.ClearLogs(g)).Name("admin:logs:clear")
	})

	adminR.Group("/plugins", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/index", adminctrl.PluginsIndexCtrl(g)).
			Name("admin:plugins:index")

		subrouter.Group("/store", func(storeSubrouter sdkhttp.HttpRouterInstance) {
			storeSubrouter.Get("/index", adminctrl.PluginsStoreIndexCtrl(g)).
				Name("admin:plugins:store:index")

			storeSubrouter.Get("/plugins/plugin", adminctrl.ViewPluginCtrl(g)).
				Name("admin:plugins:store:plugin")
		})

		subrouter.Post("/install", adminctrl.PluginsInstallCtrl(g)).
			Name("admin:plugins:install")

		subrouter.Post("/uninstall", adminctrl.UninstallPluginCtrl(g)).
			Name("admin:plugins:uninstall")
	})

	adminR.Group("/upload", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Post("/file", adminctrl.UploadFileCtrl(g)).
			Name("admin:upload:file")

		// TODO: for future use-case
		// subrouter.Post("/files", adminctrl.UploadFilesCtrl(g)).
		// Name("admin:upload:files")
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
		{
			RouteName: "plugins-index",
			RoutePath: "/plugins",
			Component: "admin/plugins/Index.vue",
		},
		{
			RouteName: "plugins-new",
			RoutePath: "/plugins/new",
			Component: "admin/plugins/NewInstall.vue",
		},
		{
			RouteName: "plugins-store",
			RoutePath: "/plugins/store",
			Component: "admin/plugins/PluginsStore.vue",
		},
		{
			RouteName: "plugin",
			RoutePath: "/plugins/store/plugin",
			Component: "admin/plugins/PluginDetail.vue",
		},
		{
			RouteName: "core-updates",
			RoutePath: "/system-updates",
			Component: "admin/CoreUpdates.vue",
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
			{
				Category:  sdkhttp.NavCategorySystem,
				Label:     "Manage Plugins",
				RouteName: "plugins-index",
			},
			{
				Category:  sdkhttp.NavCategorySystem,
				Label:     "System Updates",
				RouteName: "core-updates",
			},
		}
	})
}
