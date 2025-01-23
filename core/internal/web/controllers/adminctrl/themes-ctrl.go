package adminctrl

import (
	"errors"
	"fmt"
	"net/http"
	sdkapi "sdk/api"

	"core/internal/config"
	"core/internal/plugins"
	coreforms "core/internal/web/forms"
	"core/resources/views/admin/themes"
)

func GetAvailableThemes(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.HttpResponse()
		httpForm, ok := g.CoreAPI.HttpAPI.Forms().GetForm(coreforms.ThemesFormName)
		if !ok {
			err := errors.New("form not found: themes")
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		page := themes.AdminThemesIndex(httpForm.GetTemplate(r))
		res.AdminView(w, r, sdkapi.ViewPage{PageContent: page})
	}
}

func SaveThemeSettings(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.HttpResponse()
		httpForm, ok := g.CoreAPI.HttpAPI.Forms().GetForm(coreforms.ThemesFormName)
		if !ok {
			err := errors.New("form not found: themes")
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		err := httpForm.ParseForm(r)
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		mfdata, err := httpForm.GetMultiField("themes", "multi_field")
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		fmt.Printf("mfdata: %+v\n", mfdata)

		portalTheme, err := httpForm.GetStringValue("themes", "portal_theme")
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		adminTheme, err := httpForm.GetStringValue("themes", "admin_theme")
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		err = config.WriteThemesConfig(config.ThemesConfig{
			AdminThemePkg:  adminTheme,
			PortalThemePkg: portalTheme,
		})
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		themesIndexUrl := g.CoreAPI.HttpAPI.Helpers().UrlForRoute("admin:themes:index")
		http.Redirect(w, r, themesIndexUrl, http.StatusSeeOther)
	}
}
