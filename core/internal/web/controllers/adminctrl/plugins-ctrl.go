package adminctrl

import (
	"core/internal/plugins"
	"core/internal/utils/pkg"
	"net/http"
	sdkplugin "sdk/api/plugin"
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
