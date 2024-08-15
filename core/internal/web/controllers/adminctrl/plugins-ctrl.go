package adminctrl

import (
	"core/internal/plugins"
	"core/internal/utils/pkg"
	"encoding/json"
	"net/http"
	sdkplugin "sdk/api/plugin"
	"strings"
)

func PluginsIndexCtrl(g *plugins.CoreGlobals) http.HandlerFunc {

	type Plugin struct {
		Info sdkplugin.PluginInfo
		Src  pkg.PluginInstalledMark
	}

	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()
		sources := pkg.InstalledPluginsList()
		plugins := []Plugin{}

		for _, src := range sources {
			ok, path := pkg.IsPluginInstalled(src.Def)
			if !ok {
				continue
			}

			info, err := pkg.PluginInfo(path)
			if err != nil {
				res.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			p := Plugin{Info: info, Src: src}
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
