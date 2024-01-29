package middlewares

import (
	"context"
	"net"
	"net/http"

	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/utils/hostfinder"
	"github.com/flarehotspot/core/web/response"
)

func DeviceMiddleware(dtb *db.Database, clntMgr *connmgr.ClientRegister) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clntSym := r.Context().Value(sdkhttp.ClientCtxKey)
			if clntSym != nil {
				next.ServeHTTP(w, r)
				return
			}

			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				response.ErrorJson(w, err.Error(), 500)
				return
			}

			h, err := hostfinder.FindByIp(ip)
			if err != nil {
				response.ErrorJson(w, err.Error(), 500)
				return
			}

			clnt, err := clntMgr.Register(r.Context(), h.MacAddr, h.IpAddr, h.Hostname)
			if err != nil {
				response.ErrorJson(w, err.Error(), 500)
				return
			}

			ctx := context.WithValue(r.Context(), sdkhttp.ClientCtxKey, clnt)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
