package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"core/internal/plugins"
	"core/internal/web/response"
	sdkfs "sdk/utils/fs"
	"github.com/gorilla/mux"
)

func NewAssetsCtrl(g *plugins.CoreGlobals) *AssetsCtrl {
	return &AssetsCtrl{g}
}

type AssetsCtrl struct {
	g *plugins.CoreGlobals
}

func (ctrl *AssetsCtrl) GetFavicon(w http.ResponseWriter, r *http.Request) {
	contents, err := os.ReadFile(ctrl.g.CoreAPI.Utl.Resource("assets/images/default-favicon-32x32.png"))
	if err != nil {
		response.ErrorHtml(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(contents)
}

func (ctrl *AssetsCtrl) AssetWithHelpers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkg := vars["pkg"]
	assetPath := vars["path"]
	pluginApi, ok := ctrl.g.PluginMgr.FindByPkg(pkg)
	if !ok {
		http.Error(w, "Plugin not found: "+pkg, 404)
		return
	}

	assetPath = filepath.Join(pluginApi.Resource("assets"), assetPath)
	if !sdkfs.Exists(assetPath) {
		http.Error(w, "Asset not found: "+assetPath, 404)
		return
	}

	response.File(w, assetPath, ctrl.g.CoreAPI.Http().Helpers(), nil)
}

func (ctrl *AssetsCtrl) VueComponent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkg := vars["pkg"]
	pluginApi, ok := ctrl.g.PluginMgr.FindByPkg(pkg)
	if !ok {
		res := ctrl.g.CoreAPI.HttpAPI.HttpResponse()
		res.File(w, r, "components/empty-component.vue", vars)
		return
	}

	componentPath := filepath.Join("components", vars["path"])

	res := pluginApi.Http().HttpResponse()
	res.File(w, r, componentPath, nil)
}
