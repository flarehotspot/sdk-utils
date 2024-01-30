package routes

import (
	"net/http"

	"github.com/flarehotspot/core/plugins"
	"github.com/flarehotspot/core/web/controllers"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes/urls"
)

func BootRoutes(g *plugins.CoreGlobals) {
	bootCtrl := controllers.NewBootCtrl(g, g.PluginMgr, g.CoreAPI)
	r := router.BootingRouter
	r.Use(bootCtrl.Middleware)
	r.HandleFunc(urls.BOOT_URL, bootCtrl.IndexPage).Methods("GET")
	r.HandleFunc(urls.BOOT_STATUS_URL, bootCtrl.SseHandler).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, urls.BOOT_URL, http.StatusFound)
	})
}
