package controllers

import (
	"net/http"

	"core/internal/plugins"
	sse "core/internal/utils/sse"
)

func PortalIndexPage(g *plugins.CoreGlobals) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, themeApi, err := g.PluginMgr.GetPortalTheme()
		if err != nil {
			ErrorPage(w, err)
			return
		}

		route := themeApi.PortalTheme.IndexRoute
		url := p.HttpAPI.Helpers().UrlForRoute(route)
		http.Redirect(w, r, url, http.StatusSeeOther)
	})
}

func PortalSseHandler(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := sse.NewSocket(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		clnt, err := g.CoreAPI.HttpAPI.GetClientDevice(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sse.AddSocket(clnt.MacAddr(), s)
		s.Listen()
	}
}
