package adminctrl

import (
	"core/internal/plugins"
	rpc "core/internal/rpc"
	"core/internal/utils/pkg"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	sdkplugin "sdk/api/plugin"
	"sdk/libs/go-json"
	sdkpaths "sdk/utils/paths"
	sdkstr "sdk/utils/strings"
	"strings"
)

type PluginRelease struct {
	Id         int
	Major      int
	Minor      int
	Patch      int
	ZipFileUrl string
}

type PluginData struct {
	Id               int
	Info             sdkplugin.PluginInfo
	Src              pkg.PluginInstallData
	HasPendingUpdate bool
	ToBeRemoved      bool
	IsInstalled      bool
	Releases         []PluginRelease
}

func PluginsIndexCtrl(g *plugins.CoreGlobals) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()
		plugins := getInstalledPlugins()

		res.Json(w, plugins, http.StatusOK)
	}
}

func PluginsStoreIndexCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		srv, ctx := rpc.GetCoreMachineTwirpServiceAndCtx()
		qPlugins, err := srv.FetchPlugins(ctx, &rpc.FetchPluginsRequest{})
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

		installedPlugins := getInstalledPlugins()

		// parse pluginsData
		var pluginsData []PluginData
		for _, qP := range qPlugins.Plugins {
			pluginsData = append(pluginsData, PluginData{
				Id: int(qP.PluginId),
				Info: sdkplugin.PluginInfo{
					Name:        qP.Name,
					Package:     qP.Package,
					Description: "",
				},
				IsInstalled: isPluginInstalled(qP.Package, &installedPlugins),
			})
		}

		res.Json(w, pluginsData, http.StatusOK)
	}
}

func isPluginInstalled(pluginPkg string, installedPlugins *[]PluginData) bool {
	for _, p := range *installedPlugins {
		if pluginPkg == p.Info.Package {
			return true
		}
	}
	return false
}

func ViewPluginCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
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

		srv, ctx := rpc.GetCoreMachineTwirpServiceAndCtx()
		qPlugin, err := srv.FetchPlugin(ctx, &rpc.FetchPluginRequest{
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

		plugin := PluginData{
			Id: int(qPlugin.Plugin.PluginId),
			Info: sdkplugin.PluginInfo{
				Name:        qPlugin.Plugin.Name,
				Package:     qPlugin.Plugin.Package,
				Description: "", // TODO: add the description
			},
			Releases: pluginReleases,
		}

		res.Json(w, plugin, http.StatusOK)
	}
}

func UploadFileCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.HttpResponse()

		// limit file upload size 10 * (2 ** 20) = 10MB
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			log.Println("Error in parsing multi part form:", err)
			res.Json(w, "", http.StatusInternalServerError)
			return
		}

		// get uploaded file
		uploadedFile, handler, err := r.FormFile("file")
		if err != nil {
			log.Println("Error in opening form file: ", err)
			res.Json(w, "Error: invalid multipart file", http.StatusInternalServerError)
			return
		}
		defer uploadedFile.Close()

		// prepare parent path
		parentPath := filepath.Join(sdkpaths.UploadsDir, sdkstr.Rand(6))

		// ensure parent directory exists
		if err := os.MkdirAll(parentPath, 0755); err != nil {
			log.Println("Error creating parent dir:", err)
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// create destination file
		filePath := filepath.Join(parentPath, handler.Filename)
		prZipFile, err := os.Create(filePath)
		if err != nil {
			log.Println("Error creating pr zip file:", err)
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer prZipFile.Close()

		// copy the contents of the uploaded file on to the created destination file
		if _, err := io.Copy(prZipFile, uploadedFile); err != nil {
			log.Println("Error copying file:", err)
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("%s successfully uploaded", filePath)
		res.Json(w, filePath, http.StatusOK)
	}
}

// TODO: update for multiple files for future use-case
func UploadFilesCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.HttpResponse()

		// TODO: implementation of multiple file uploads

		res.Json(w, "", http.StatusOK)
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

func getInstalledPlugins() []PluginData {
	sources := pkg.InstalledPluginsList()
	plugins := []PluginData{}

	for _, src := range sources {
		info, err := pkg.GetPluginInfo(src.Def)
		if err != nil {
			return nil
		}

		p := PluginData{
			Info:             info,
			Src:              src,
			HasPendingUpdate: pkg.HasPendingUpdate(info.Package),
			ToBeRemoved:      pkg.IsToBeRemoved(info.Package),
		}

		plugins = append(plugins, p)
	}

	return plugins
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
