package routes

import (
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/controllers/adminctrl"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes/names"
)

func AdminRoutes(g *globals.CoreGlobals) {
	dashboardCtrl := adminctrl.NewDashboardCtrl(g)
	pluginsCtrl := adminctrl.NewPluginsCtrl(g.PluginMgr, g.CoreApi)
	bandwidthCtrl := adminctrl.NewBandwidthCtrl(g, g.CoreApi)
	r := router.AdminRouter()

	r.HandleFunc("/", dashboardCtrl.RedirectToDash)
	r.HandleFunc("/dashboard", dashboardCtrl.Index).Methods("GET").Name(names.RouteAdminDashboardIndex)
	r.HandleFunc("/plugins", pluginsCtrl.Index).Methods("GET").Name(names.RouteAdminPluginsIndex)
	r.HandleFunc("/plugins/new", pluginsCtrl.New).Methods("GET").Name(names.RouteAdminPluginsNew)
	r.HandleFunc("/plugins/upload", pluginsCtrl.Upload).Methods("POST").Name(names.RouteAdminPluginUpload)
	r.HandleFunc("/bandwidth/index", bandwidthCtrl.Index).Methods("GET").Name(names.RouteAdminBandwidthIndex)
	r.HandleFunc("/bandwidth/{ifname}/save", bandwidthCtrl.Save).Methods("POST").Name(names.RouteAdminBandwidthSave)
}
