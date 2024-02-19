package middlewares

import (
	"github.com/flarehotspot/core/internal/web/helpers"
	"github.com/flarehotspot/core/internal/web/response"
	"net/http"
)

func AssetPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if helpers.IsAssetPath(r.URL.Path) {
			next.ServeHTTP(w, r)
		} else {
			response.ErrorJson(w, "Invalid asset path: "+r.URL.Path, 404)
		}
	})
}
