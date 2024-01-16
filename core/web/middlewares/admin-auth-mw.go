package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/flarehotspot/core/accounts"
	"github.com/flarehotspot/core/config/appcfg"
	"github.com/flarehotspot/core/sdk/libs/jwt"
	"github.com/flarehotspot/core/sdk/utils/contexts"
	"github.com/flarehotspot/core/sdk/utils/cookie"
	"github.com/flarehotspot/core/sdk/utils/translate"
	"github.com/flarehotspot/core/utils/jsonwebtoken"
)

const (
	AuthTokenCookie = "auth-token"
)

func AdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acct, err := IsAdminAuthenticated(w, r)
		if err != nil {
			log.Println("Invalid login: ", err)
			autherr := translate.Core(translate.Error, "unauthorized")
			ErrUnauthorized(w, autherr)
			return
		}

		ctx := context.WithValue(r.Context(), contexts.SysAcctCtxKey, acct)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func IsAdminAuthenticated(w http.ResponseWriter, r *http.Request) (*accounts.Account, error) {
	authtoken, err := cookie.GetCookie(r, AuthTokenCookie)
	if err != nil {
		bearer := r.Header.Get("Authorization")
		splitToken := strings.Split(bearer, "Bearer ")
		if len(splitToken) != 2 {
			return nil, errors.New("invalid auth token")
		}

		authtoken = splitToken[1]
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

	return accounts.Find(username)
}

func ErrUnauthorized(w http.ResponseWriter, msg string) {
	data := map[string]string{"error": msg}
	jsonData, _ := json.Marshal(data)
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
