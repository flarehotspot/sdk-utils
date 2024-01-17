package controllers

import (
	"net/http"

	"github.com/flarehotspot/core/config/themecfg"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/plugins"
)

func NewIndexPageCtrl(g *globals.CoreGlobals) IndexPageCtrl {
	return IndexPageCtrl{g}
}

type IndexPageCtrl struct {
	g *globals.CoreGlobals
}

func (c *IndexPageCtrl) PortalIndex(w http.ResponseWriter, r *http.Request) {
	themePkg := themecfg.Read().Portal
	themePlugin := c.g.PluginMgr.FindByPkg(themePkg)
	themesApi := themePlugin.ThemesApi().(*plugins.ThemesApi)
	portalComponent, ok := themesApi.GetPortalThemeComponents()
	if !ok {
		http.Error(w, "No portal theme component path defined", 500)
		return
	}

	vdata := map[string]any{
		"CoreApi":      c.g.CoreApi,
		"ThemeApi":     themePlugin.(*plugins.PluginApi),
		"ThemeScripts": portalComponent.ThemeAssets.Scripts,
		"ThemeStyles":  portalComponent.ThemeAssets.Styles,
	}

	api := c.g.CoreApi
	api.HttpApi().Respond().View(w, r, "portal/vue-layout.html", vdata)
}

func (c *IndexPageCtrl) AdminIndex(w http.ResponseWriter, r *http.Request) {
	themePkg := themecfg.Read().Admin
	themePlugin := c.g.PluginMgr.FindByPkg(themePkg)
	themesApi := themePlugin.ThemesApi().(*plugins.ThemesApi)
	adminThemeComponent, ok := themesApi.GetAdminLayoutComponents()
	if !ok {
		http.Error(w, "No admin theme component path defined", 500)
		return
	}

	vdata := map[string]any{
		"CoreApi":      c.g.CoreApi,
		"ThemeApi":     themePlugin.(*plugins.PluginApi),
		"ThemeScripts": adminThemeComponent.ThemeAssets.Scripts,
		"ThemeStyles":  adminThemeComponent.ThemeAssets.Styles,
	}

	api := c.g.CoreApi
	api.HttpApi().Respond().View(w, r, "admin/vue-layout.html", vdata)
}
