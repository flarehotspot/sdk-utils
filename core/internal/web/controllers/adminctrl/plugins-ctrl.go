package adminctrl

import (
	"core/internal/plugins"
	"core/internal/utils/pkg"
	"log"
	"net/http"
	sdkplugin "sdk/api/plugin"
	"sdk/libs/go-json"
	"strings"

	coremachine_v0_0_1 "core/internal/rpc/machines/coremachines/v0_0_1"
	"core/internal/rpc/twirp"
)

func PluginsIndexCtrl(g *plugins.CoreGlobals) http.HandlerFunc {

	type PluginData struct {
		Info             sdkplugin.PluginInfo
		Src              pkg.PluginInstallData
		HasPendingUpdate bool
		ToBeRemoved      bool
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
				ToBeRemoved:      pkg.IsToBeRemoved(info.Package),
			}

			plugins = append(plugins, p)
		}

		srv, ctx := twirp.GetCoreMachineTwirpServiceAndCtx()
		_, err := srv.FetchPlugins(ctx, &coremachine_v0_0_1.FetchPluginsRequest{})
		if err != nil {
			log.Println("Error:", err)
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

func UninstallPluginCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()
		// read post body as json
		var data struct {
			Pkg string `json:"pkg"`
		}

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = pkg.MarkToRemove(data.Pkg)
		if err != nil {
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, http.StatusOK)
	}
}
