package middlewares

import (
	"context"
	"errors"
	"net/http"

	coreAcct "github.com/flarehotspot/core/accounts"
	"github.com/flarehotspot/core/config/appcfg"
	"github.com/flarehotspot/core/sdk/libs/jwt"
	"github.com/flarehotspot/core/sdk/utils/contexts"
	"github.com/flarehotspot/core/sdk/utils/cookie"
	"github.com/flarehotspot/core/sdk/utils/flash"
	"github.com/flarehotspot/core/sdk/utils/translate"
	"github.com/flarehotspot/core/utils/jsonwebtoken"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes/names"
)

func AdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acct, err := IsAdminAuthenticated(w, r)
		if err != nil {
			autherr := translate.Core(translate.Error, "unauthorized")
			flash.SetFlashMsg(w, flash.Error, autherr)
			loginUrl, _ := router.UrlForRoute(names.RouteAuthLogin)
			http.Redirect(w, r, loginUrl, http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), contexts.SysAcctCtxKey, acct)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func MustBeLoggedOut(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := IsAdminAuthenticated(w, r)
		if err != nil {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		}
	})
}

func IsAdminAuthenticated(w http.ResponseWriter, r *http.Request) (*coreAcct.Account, error) {
	authtoken, err := cookie.GetCookie(r, "auth-token")
	if err != nil {
		return nil, err
	}

	cfg, err := appcfg.Read()
	if err != nil {
		return nil, err
	}

	token, err := jsonwebtoken.VerifyToken(authtoken, cfg.Secret)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid jwt claims")
	}

	username := claims["username"].(string)

	return coreAcct.Find(username)
}
