package routes

import (
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/web/middlewares"
	"github.com/gorilla/mux"
)

func PublicAssets(prefix string, dir string, r *mux.Router) {
	pubprefix := filepath.Join("/public", prefix)
	fs := http.FileServer(http.Dir(dir))

	h := middlewares.AssetPath(http.StripPrefix(pubprefix, fs))
	r.PathPrefix(pubprefix).Handler(h)
}
