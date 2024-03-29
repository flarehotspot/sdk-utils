package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/flarehotspot/core/internal/config"
	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/utils"
)

func SaveThemeSettings(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// save the settings
		var thm utils.SavedThemeData
		err := json.NewDecoder(r.Body).Decode(&thm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cfg := config.ThemesConfig{
			//registered plugins with theme here
			Portal: thm.ThemePortal,
			Admin:  thm.ThemeAdmin,
		}
		err = config.WriteThemesConfig(cfg)
		if err != nil {
			g.CoreAPI.HttpAPI.VueResponse().SetFlashMsg("error", err.Error())
			return
		}
		g.CoreAPI.HttpAPI.VueResponse().Json(w, cfg, 200)
	}
}
