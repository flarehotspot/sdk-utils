package adminctrl

import (
	"encoding/json"
	"net/http"

	"github.com/flarehotspot/core/internal/config"
	"github.com/flarehotspot/core/internal/plugins"
)

func GetAvailableThemes(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()
		allPlugins := g.PluginMgr.All()
		adminThemes := []map[string]string{}
		portalThemes := []map[string]string{}

		for _, p := range allPlugins {
			features := p.Features()
			for _, f := range features {
				if f == "theme:admin" {
					adminThemes = append(adminThemes, map[string]string{
						"name": p.Name(),
						"pkg":  p.Pkg(),
					})
				}

				if f == "theme:portal" {
					portalThemes = append(portalThemes, map[string]string{
						"name": p.Name(),
						"pkg":  p.Pkg(),
					})
				}
			}
		}

		cfg, err := config.ReadThemesConfig()
		if err != nil {
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"admin_themes":  adminThemes,
			"portal_themes": portalThemes,
			"themes_config": cfg,
		}

		res.Json(w, data, http.StatusOK)
	}
}

func SaveThemeSettings(g *plugins.CoreGlobals) http.HandlerFunc {

	type ThemeSettings struct {
		AdminTheme  string `json:"admin"`
		PortalTheme string `json:"portal"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()
		var thm ThemeSettings
		err := json.NewDecoder(r.Body).Decode(&thm)
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cfg := config.ThemesConfig{
			Portal: thm.PortalTheme,
			Admin:  thm.AdminTheme,
		}

		err = config.WriteThemesConfig(cfg)
		if err != nil {
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, cfg, http.StatusOK)
	}
}
