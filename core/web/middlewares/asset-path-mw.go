package middlewares

import (
	"errors"
	"github.com/flarehotspot/core/web/helpers"
	"github.com/flarehotspot/core/web/response"
	"net/http"
)

func AssetPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if helpers.IsAssetPath(r.URL.Path) {
			next.ServeHTTP(w, r)
		} else {
			response.Error(w, errors.New("Invalid asset path: "+r.URL.Path))
		}
	})
}
