package adminctrl

import (
	"core/internal/plugins"
	"core/internal/utils/pkg"
	"net/http"
	sdkplugin "sdk/api/plugin"
	"sdk/libs/go-json"
	"strings"
)

func PluginsIndexCtrl(g *plugins.CoreGlobals) http.HandlerFunc {

	type PluginData struct {
		Info             sdkplugin.PluginInfo
		Src              pkg.PluginInstallData
		HasPendingUpdate bool
	}

	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()
		sources := pkg.InstalledPluginsList()
		plugins := []PluginData{}

		for _, src := range sources {
			info, err := pkg.GetPluginInfo(src.Def)
			if err != nil {
				res.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			p := PluginData{
				Info:             info,
				Src:              src,
				HasPendingUpdate: pkg.HasPendingUpdate(info.Package),
			}

			plugins = append(plugins, p)
		}

		res.Json(w, plugins, http.StatusOK)
	}
}

func PluginsInstallCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()
		// read post body as json
		var data pkg.PluginSrcDef
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var result strings.Builder
		info, err := pkg.InstallSrcDef(&result, data)
		if err != nil {
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, info, http.StatusOK)
	}
}
