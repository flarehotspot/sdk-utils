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
	portalComponent, ok := themesApi.GetPortalComponent()
	if !ok {
		http.Error(w, "No portal theme component path defined", 500)
		return
	}

	scripts := []string{}
	styles := []string{}

	if portalComponent.ThemeAssets != nil {
		if portalComponent.ThemeAssets.Scripts != nil {
			for _, script := range portalComponent.ThemeAssets.Scripts {
				jsPath := themePlugin.HttpApi().Helpers(w, r).AssetPath(script)
				scripts = append(scripts, jsPath)
			}
		}

		if portalComponent.ThemeAssets.Styles != nil {
			for _, style := range portalComponent.ThemeAssets.Styles {
				cssPath := themePlugin.HttpApi().Helpers(w, r).AssetPath(style)
				styles = append(styles, cssPath)
			}
		}
	}

	vdata := map[string]any{
		"CoreApi":      c.g.CoreApi,
		"ThemeScripts": scripts,
		"ThemeStyles":  styles,
	}

	api := c.g.CoreApi
	api.HttpApi().Respond().View(w, r, "portal/layout.html", vdata)
}

func (c *IndexPageCtrl) AdminIndex(w http.ResponseWriter, r *http.Request) {
	themePkg := themecfg.Read().Admin
	themePlugin := c.g.PluginMgr.FindByPkg(themePkg)
	themesApi := themePlugin.ThemesApi().(*plugins.ThemesApi)
	adminThemeComponent, ok := themesApi.GetAdminThemeComponent()
	if !ok {
		http.Error(w, "No admin theme component path defined", 500)
		return
	}

	scripts := []string{}
	styles := []string{}

	if adminThemeComponent.ThemeAssets != nil {
		if adminThemeComponent.ThemeAssets.Scripts != nil {
			for _, script := range adminThemeComponent.ThemeAssets.Scripts {
				jsPath := themePlugin.HttpApi().Helpers(w, r).AssetPath(script)
				scripts = append(scripts, jsPath)
			}
		}

		if adminThemeComponent.ThemeAssets.Styles != nil {
			for _, style := range adminThemeComponent.ThemeAssets.Styles {
				cssPath := themePlugin.HttpApi().Helpers(w, r).AssetPath(style)
				styles = append(styles, cssPath)
			}
		}
	}

	vdata := map[string]any{
		"CoreApi":      c.g.CoreApi,
		"ThemeScripts": scripts,
		"ThemeStyles":  styles,
	}

	api := c.g.CoreApi
	api.HttpApi().Respond().View(w, r, "admin/layout.html", vdata)
}
