package adminctrl

import (
	"core/internal/plugins"
	"core/internal/utils/pkg"
	"net/http"
)

func PluginsIndexCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()
		plugins := pkg.InstalledPluginsList()
		res.Json(w, plugins, http.StatusOK)
	}
}
