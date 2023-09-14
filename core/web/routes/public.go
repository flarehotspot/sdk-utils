package routes

import (
	"net/http"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/sdk/utils/paths"
	"github.com/gorilla/mux"
)

func PublicAssets(r *mux.Router, g *globals.CoreGlobals) {
	fs := http.FileServer(http.Dir(paths.PublicDir))
	r.PathPrefix("/public/").
		Handler(middlewares.AssetPath(http.StripPrefix("/public/", fs)))
}
