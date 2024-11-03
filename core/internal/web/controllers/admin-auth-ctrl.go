package controllers

import (
	"core/internal/plugins"
	"net/http"
	sdkhttp "sdk/api/http"
)

func AdminLoginCtrl(g *plugins.CoreGlobals) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.HttpResponse()
		_, t, err := g.PluginMgr.GetAdminTheme()
		if err != nil {
			res.ErrorPage(w, r, err, http.StatusInternalServerError)
			return
		}

		data := sdkhttp.LoginPageData{
			LoginUrl: g.CoreAPI.HttpAPI.Helpers().UrlForRoute("admin:authenticate"),
		}

		page := t.AdminTheme.LoginPageFactory(w, r, data)
		g.CoreAPI.HttpAPI.HttpResponse().AdminView(w, r, page)
	})
}
