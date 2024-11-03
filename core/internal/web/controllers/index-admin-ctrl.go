package controllers

import (
	"log"
	"net/http"

	"core/internal/plugins"
	sse "core/internal/utils/sse"
)

func AdminIndexPage(g *plugins.CoreGlobals) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, themeApi, err := g.PluginMgr.GetAdminTheme()
		if err != nil {
			ErrorPage(w, err)
			return
		}

		route := themeApi.AdminTheme.IndexRoute
		log.Println("Theme:", p.Pkg())
		log.Println("IndexRoute: ", route)
		log.Printf("AdminTheme: %+v\n", themeApi.AdminTheme)

		url := p.HttpAPI.Helpers().UrlForRoute(route)
		http.Redirect(w, r, url, http.StatusSeeOther)
	})
}

func AdminSseHandler(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := sse.NewSocket(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		acct, err := g.CoreAPI.HttpAPI.Auth().CurrentAcct(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sse.AddSocket(acct.Username(), s)
		s.Listen()
	}
}

func ErrorPage(w http.ResponseWriter, err error) {
	// TODO: show error page
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}
