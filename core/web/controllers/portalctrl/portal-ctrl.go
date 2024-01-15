package portalctrl

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/flarehotspot/core/config/appcfg"
	"github.com/flarehotspot/core/config/themecfg"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/plugins"
	"github.com/flarehotspot/core/web/helpers"
	"github.com/flarehotspot/core/web/response"
	"github.com/flarehotspot/core/web/routes/names"
	"github.com/gorilla/mux"
)

type PortalCtrl struct {
	g *globals.CoreGlobals
}

func NewPortalCtrl(g *globals.CoreGlobals) PortalCtrl {
	return PortalCtrl{g}
}

func (c *PortalCtrl) IndexPage(w http.ResponseWriter, r *http.Request) {
	themePkg := themecfg.Read().CaptivePortal
	themePlugin := c.g.PluginMgr.FindByPkg(themePkg)
	themesApi := themePlugin.ThemesApi().(*plugins.ThemesApi)
	portalComponent, ok := themesApi.GetPortalComponent()
	if !ok {
		http.Error(w, "No portal theme component path defined", 500)
		return
	}

	scripts := []string{}
	styles := []string{}

	if portalComponent.ThemeAssets != nil {
		if portalComponent.ThemeAssets.Scripts != nil {
			for _, script := range *portalComponent.ThemeAssets.Scripts {
				jsPath := themePlugin.HttpApi().Helpers(w, r).AssetPath(script)
				scripts = append(scripts, jsPath)
			}
		}

		if portalComponent.ThemeAssets.Styles != nil {
			for _, style := range *portalComponent.ThemeAssets.Styles {
				cssPath := themePlugin.HttpApi().Helpers(w, r).AssetPath(style)
				styles = append(styles, cssPath)
			}
		}
	}

	vdata := map[string]any{
		"ThemeScripts": scripts,
		"ThemeStyles":  styles,
	}

	api := c.g.CoreApi
	api.HttpApi().Respond().View(w, r, "captive-portal/layout.html", vdata)
}

func (c *PortalCtrl) FaviconIco(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := os.ReadFile(c.g.CoreApi.Resource("assets/images/default-favicon-32x32.png"))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}

func (c *PortalCtrl) MainJs(w http.ResponseWriter, r *http.Request) {
	themePkg := themecfg.Read().CaptivePortal
	themePlugin := c.g.PluginMgr.FindByPkg(themePkg)
	themesApi := themePlugin.ThemesApi().(*plugins.ThemesApi)
	portalComponent, ok := themesApi.GetPortalComponent()
	if !ok {
		http.Error(w, "No portal theme component path defined", 500)
		return
	}

	allPlugins := c.g.PluginMgr.All()
	routes := []*plugins.VueRoute{}

	for _, p := range allPlugins {
		vueRouter := p.HttpApi().VueRouter().(*plugins.VueRouter)
		portalRoutes := vueRouter.GetPortalRoutes(r)
		routes = append(routes, portalRoutes...)
	}

	routesJson, err := json.Marshal(routes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	appConfig, err := appcfg.Read()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := map[string]any{
		"CoreApi":       c.g.CoreApi,
		"Routes":        string(routesJson),
		"AssetsVersion": appConfig.AssetsVersion,
		"PortalTheme": map[string]any{
			"LayoutComponent": filepath.Join(themePlugin.Pkg(), portalComponent.ThemeComponentPath),
			"IndexComponent":  filepath.Join(themePlugin.Pkg(), portalComponent.IndexComponentPath),
		},
	}

	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	c.g.CoreApi.HttpAPI.Respond().Text(w, r, "templates/portal/js/main.tpl.js", data)
}

func (c *PortalCtrl) HelpersJs(g *globals.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		pkg, ok := vars["pkg"]
		if !ok {
			http.Error(w, "need to specify plugin package name", 500)
			return
		}

		plugin := g.PluginMgr.FindByPkg(pkg)
		if plugin == nil {
			http.Error(w, "invalid plugin package name", 500)
			return
		}

		vueRouter := plugin.HttpApi().VueRouter().(*plugins.VueRouter)
		routes := vueRouter.GetPortalRoutes(r)
		routesJson, err := json.Marshal(routes)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		vdata := map[string]any{
			"Plugin": plugin,
			"Routes": string(routesJson),
		}

		w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
		c.g.CoreApi.HttpAPI.Respond().Text(w, r, "templates/portal/js/helpers.tpl.js", vdata)
	}
}

// TODO: Remove this test method
func (c *PortalCtrl) Test(w http.ResponseWriter, r *http.Request) {
	clnt, err := helpers.CurrentClient(r)
	if err != nil {
		c.Error(w, r, err)
		return
	}

	c.g.Models.Session().Create(r.Context(), clnt.Id(), 0, 30, 0, nil, 1, 1, false)

	w.WriteHeader(200)
}

type testpage struct {
	Title string
	Data  map[string]any
}

func (self *PortalCtrl) TestTemplate(w http.ResponseWriter, r *http.Request) {
	p := &testpage{
		Title: "Some page title",
		Data:  map[string]any{"data": "data value"},
	}
	t, err := template.New("page").Parse("Title is \"{{ .Title }}\" and data is \"{{ .Data.data }}\".")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, p)
	if err != nil {
		panic(err)
	}
}

func (self *PortalCtrl) Error(w http.ResponseWriter, r *http.Request, err error) {
	e := response.NewErrRoute(names.RoutePortalIndex)
	e.Redirect(w, r, err)
}
