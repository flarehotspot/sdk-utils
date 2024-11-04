package controllers

import (
	"core/internal/plugins"
	"core/internal/web/router"
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

		authRoute := router.RootRouter.Get("admin:authenticate")
		authUrl, _ := authRoute.URL()

		data := sdkhttp.LoginPageData{
			LoginUrl: authUrl.String(),
		}

		page := t.AdminTheme.LoginPageFactory(w, r, data)
		g.CoreAPI.HttpAPI.HttpResponse().AdminView(w, r, page)
	})
}

func AdminAuthenticateCtrl(g *plugins.CoreGlobals) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
