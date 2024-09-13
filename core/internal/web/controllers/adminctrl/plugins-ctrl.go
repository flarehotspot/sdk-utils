package adminctrl

import (
	"core/internal/plugins"
	coremachine_v0_0_1 "core/internal/rpc/machines/coremachines/v0_0_1"
	"core/internal/rpc/twirp"
	"core/internal/utils/pkg"
	"errors"
	"log"
	"net/http"
	sdkplugin "sdk/api/plugin"
	"sdk/libs/go-json"
	sdkstr "sdk/utils/strings"
	"strings"
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

		res.Json(w, plugins, http.StatusOK)
	}
}

func PluginsStoreIndexCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	type Plugin struct {
		Id      int
		Name    string
		Package string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		srv, ctx := twirp.GetCoreMachineTwirpServiceAndCtx()
		qPlugins, err := srv.FetchPlugins(ctx, &coremachine_v0_0_1.FetchPluginsRequest{})
		if err != nil {
			log.Println("Error:", err)
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if qPlugins == nil {
			err := errors.New("queried plugins is nil")
			log.Println("Error:", err)
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// parse plugins
		var plugins []Plugin
		for _, qP := range qPlugins.Plugins {
			plugins = append(plugins, Plugin{
				Id:      int(qP.PluginId),
				Name:    qP.Name,
				Package: qP.Package,
			})
		}

		res.Json(w, plugins, http.StatusOK)
	}
}

func ViewPluginCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	type PluginRelease struct {
		Id         int
		Major      int
		Minor      int
		Patch      int
		ZipFileUrl string
	}

	type Plugin struct {
		Id       int
		Name     string
		Package  string
		Releases []PluginRelease
	}

	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		// parse query
		pluginId := sdkstr.AtoiOrDefault(r.URL.Query().Get("id"), 0)

		if pluginId == 0 {
			err := errors.New("invalid plugin id")
			log.Println("Error:", err)
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		srv, ctx := twirp.GetCoreMachineTwirpServiceAndCtx()
		qPlugin, err := srv.FetchPlugin(ctx, &coremachine_v0_0_1.FetchPluginRequest{
			PluginId: int32(pluginId),
		})
		if err != nil {
			log.Println("Error:", err)
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if qPlugin == nil {
			err := errors.New("queried plugin is nil")
			log.Println("Error:", err)
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// parse plugin
		var pluginReleases []PluginRelease
		for _, qpr := range qPlugin.Releases {
			pluginReleases = append(pluginReleases, PluginRelease{
				Id:         int(qpr.PluginReleaseId),
				Major:      int(qpr.Major),
				Minor:      int(qpr.Minor),
				Patch:      int(qpr.Patch),
				ZipFileUrl: qpr.ZipFileUrl,
			})
		}

		plugin := Plugin{
			Id:       int(qPlugin.Plugin.PluginId),
			Name:     qPlugin.Plugin.Name,
			Package:  qPlugin.Plugin.Package,
			Releases: pluginReleases,
		}

		res.Json(w, plugin, http.StatusOK)
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
