package routes

import (
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/flarehotspot/core/web/middlewares"
	"github.com/gorilla/mux"
)

func PublicAssets(prefix string, dir string, r *mux.Router) {
	fs := http.FileServer(http.Dir(dir))
	h := middlewares.AssetPath(StripPrefix(prefix, fs))
	pubprefix := "/assets/{version}/" + prefix
	r.PathPrefix(pubprefix).Handler(h)
}

// Taken from http.StripPrefix function
func StripPrefix(prefix string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		version := vars["version"]
		pubprefix := filepath.Join("/assets", version, prefix)

		p := strings.TrimPrefix(r.URL.Path, pubprefix)
		rp := strings.TrimPrefix(r.URL.RawPath, pubprefix)
		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			r2.URL.RawPath = rp
			h.ServeHTTP(w, r2)
		} else {
			http.NotFound(w, r)
		}
	})
}
