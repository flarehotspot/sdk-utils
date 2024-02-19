package middlewares

import (
	"net/http"

	acct "github.com/flarehotspot/core/internal/accounts"
	"github.com/flarehotspot/core/sdk/api/http"
)

func PermsMiddleware(perms []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			acctsym := r.Context().Value(sdkhttp.SysAcctCtxKey)
			if acctsym == nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			acct := acctsym.(acct.Account)
			pcount := 0

			for _, perm := range perms {
				for _, acctperm := range acct.Permissions() {
					if perm == acctperm {
						pcount++
					}
				}
			}

			if pcount < len(perms) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
