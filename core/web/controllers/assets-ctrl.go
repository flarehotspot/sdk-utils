package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/flarehotspot/core/config"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/plugins"
	fs "github.com/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/core/web/response"
	"github.com/gorilla/mux"
)

func NewAssetsCtrl(g *globals.CoreGlobals) *AssetsCtrl {
	return &AssetsCtrl{g}
}

type AssetsCtrl struct {
	g *globals.CoreGlobals
}

func (c *AssetsCtrl) GetFavicon(w http.ResponseWriter, r *http.Request) {
	contents, err := os.ReadFile(c.g.CoreAPI.Utl.Resource("assets/images/default-favicon-32x32.png"))
	if err != nil {
		response.ErrorHtml(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(contents)
}

func (c *AssetsCtrl) AssetWithHelpers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkg := vars["pkg"]
	assetPath := vars["path"]
	pluginApi, ok := c.g.PluginMgr.FindByPkg(pkg)
	if !ok {
		http.Error(w, "Plugin not found: "+pkg, 404)
		return
	}

	assetPath = filepath.Join(pluginApi.Utils().Resource("assets"), assetPath)
	if !fs.Exists(assetPath) {
		http.Error(w, "Asset not found: "+assetPath, 404)
		return
	}

	response.File(w, assetPath, c.g.CoreAPI.HttpApi().Helpers(), nil)
}

func (c *AssetsCtrl) VueComponent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkg := vars["pkg"]
	componentPath := vars["path"]
	pluginApi, ok := c.g.PluginMgr.FindByPkg(pkg)
	if !ok {
		c.g.CoreAPI.HttpAPI.VueResponse().Component(w, "empty-component.vue", vars)
		return
	}

	res := pluginApi.HttpApi().VueResponse()
	res.Component(w, componentPath, nil)
}

func (c *AssetsCtrl) FormField(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	comp := vars["component"] // e.g. input.vue

	cfg, err := config.ReadThemesConfig()
	if err != nil {
		response.ErrorHtml(w, err.Error())
		return
	}

	theme, ok := c.g.PluginMgr.FindByPkg(cfg.Admin)
	if !ok {
		response.ErrorHtml(w, "Theme not found")
		return
	}

	inputFile := filepath.Join(theme.Utils().Resource("components/forms/" + comp))
	if !fs.IsFile(inputFile) {
		themesApi := theme.ThemesApi().(*plugins.ThemesApi)
		csslib := themesApi.AdminTheme.CssLib
		inputFile = filepath.Join(c.g.CoreAPI.Utils().Resource("components/forms/" + string(csslib) + "/" + comp))
	}

	response.File(w, inputFile, nil, 200)
}
