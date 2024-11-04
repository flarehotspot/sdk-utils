package adminctrl

import (
	"core/internal/plugins"
	"core/internal/web/router"
	"net/http"
	sdkhttp "sdk/api/http"

	"github.com/gorilla/csrf"
)

func AdminLoginCtrl(g *plugins.CoreGlobals) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.HttpResponse()
		_, t, err := g.PluginMgr.GetAdminTheme()
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		authRoute := router.RootRouter.Get("admin:authenticate")
		authUrl, _ := authRoute.URL()
		csrf := csrf.TemplateField(r)

		data := sdkhttp.LoginPageData{
			CsrfHTML: string(csrf),
			LoginUrl: authUrl.String(),
		}

		page := t.AdminTheme.LoginPageFactory(w, r, data)
		g.CoreAPI.HttpAPI.HttpResponse().AdminView(w, r, page)
	})
}

func AdminAuthenticateCtrl(g *plugins.CoreGlobals) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			// TODO: Handle error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		acct, err := g.CoreAPI.HttpAPI.Auth().Authenticate(username, password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		g.CoreAPI.HttpAPI.Auth().SignIn(w, acct)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	})
}
