package apiv1

import (
	"errors"
	"net/http"

	"github.com/flarehotspot/core/accounts"
	"github.com/flarehotspot/core/config/appcfg"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/sdk/utils/cookie"
	"github.com/flarehotspot/core/sdk/utils/translate"
	"github.com/flarehotspot/core/utils/jsonwebtoken"
	"github.com/flarehotspot/core/web/middlewares"
)

func NewAdminAuthCtrl(g *globals.CoreGlobals) *AdminAuthCtrl {
	return &AdminAuthCtrl{g}
}

type AdminAuthCtrl struct {
	g *globals.CoreGlobals
}

func (c *AdminAuthCtrl) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	acct, err := accounts.Find(username)
	if err != nil {
		c.ErrorUnauthorized(w, err.Error())
		return
	}

	if !acct.Auth(password) {
		err = errors.New(translate.Core(translate.Error, "invalid_login"))
		c.ErrorUnauthorized(w, err.Error())
		return
	}

	cfg, err := appcfg.Read()
	if err != nil {
		err = errors.New(translate.Core(translate.Error, "invalid_login"))
		c.ErrorUnauthorized(w, err.Error())
		return
	}

	payload := map[string]string{"username": username}
	token, err := jsonwebtoken.GenerateToken(payload, cfg.Secret)
	if err != nil {
		c.ErrorUnauthorized(w, err.Error())
		return
	}

	cookie.SetCookie(w, middlewares.AuthTokenCookie, token)
	data := map[string]string{"token": token}
	c.g.CoreApi.HttpApi().HttpResponse().Json(w, data, http.StatusOK)
}

func (c *AdminAuthCtrl) Logout(w http.ResponseWriter, r *http.Request) {
	cookie.SetCookie(w, middlewares.AuthTokenCookie, "")
	data := map[string]string{"message": "Logout success"}
	c.g.CoreApi.HttpApi().HttpResponse().Json(w, data, http.StatusOK)
}

func (c *AdminAuthCtrl) IsAuthenticated(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"message": "Success"}
	c.g.CoreApi.HttpApi().HttpResponse().Json(w, data, http.StatusOK)
}

func (c *AdminAuthCtrl) ErrorUnauthorized(w http.ResponseWriter, msg string) {
	data := map[string]string{"error": msg}
	c.g.CoreApi.HttpApi().HttpResponse().Json(w, data, http.StatusUnauthorized)
}
