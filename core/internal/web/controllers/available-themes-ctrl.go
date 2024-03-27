package controllers

import (
	"net/http"

	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/utils"
)

func GetAvailableThemes(g *plugins.CoreGlobals) utils.TmpThemeList {
	allPlugins := g.PluginMgr.All()
	var themeName utils.TmpThemeList

	for _, p := range allPlugins {
		themeFeatures := p.(*plugins.PluginApi)
		if themeFeatures.ThemesAPI.AdminTheme != nil || themeFeatures.ThemesAPI.PortalTheme != nil {
			themeName.ThemeAdmin = append(themeName.ThemeAdmin, struct {
				ThemeName string "json:\"theme_name\""
				ThemePkg  string "json:\"theme_pkg\""
			}{
				ThemeName: themeFeatures.Name(),
				ThemePkg:  themeFeatures.Pkg(),
			})
			themeName.ThemePortal = append(themeName.ThemePortal, struct {
				ThemeName string "json:\"theme_name\""
				ThemePkg  string "json:\"theme_pkg\""
			}{
				ThemeName: themeFeatures.Name(),
				ThemePkg:  themeFeatures.Pkg(),
			})
		}
	}
	return themeName
}
func RespondJsonThemes(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		themes := GetAvailableThemes(g)
		g.CoreAPI.HttpAPI.VueResponse().Json(w, themes, http.StatusOK)
	}
}
