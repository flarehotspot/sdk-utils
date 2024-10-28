package middlewares

import (
	"context"
	"net"
	"net/http"

	"core/internal/connmgr"
	"core/internal/db"
	"sdk/api/http"
	"core/internal/utils/hostfinder"
	"core/internal/web/response"
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

			clnt, err := clntMgr.Register(r, h.MacAddr, h.IpAddr, h.Hostname)
			if err != nil {
				response.ErrorJson(w, err.Error(), 500)
				return
			}

			ctx := context.WithValue(r.Context(), sdkhttp.ClientCtxKey, clnt)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
