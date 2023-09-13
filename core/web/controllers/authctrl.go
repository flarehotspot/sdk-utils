package controllers

import (
	"errors"
	"net/http"

	"github.com/flarehotspot/core/accounts"
	"github.com/flarehotspot/core/config/appcfg"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/themes"
	"github.com/flarehotspot/core/utils/jsonwebtoken"
	"github.com/flarehotspot/core/web/response"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes/names"
	"github.com/flarehotspot/core/sdk/utils/cookie"
	"github.com/flarehotspot/core/sdk/utils/flash"
	"github.com/flarehotspot/core/sdk/utils/translate"
)

const (
	authTokenCookie = "auth-token"
)

type AuthCtrl struct {
	g *globals.CoreGlobals
}

func NewAuthCtrl(g *globals.CoreGlobals) *AuthCtrl {
	return &AuthCtrl{g}
}

func (ctrl *AuthCtrl) GetLogin(w http.ResponseWriter, r *http.Request) {
	api := ctrl.g.PluginMgr.AdminPluginApi()
	api.HttpApi().Respond().View(w, r, themes.LoginHtml, nil)
}

func (ctrl *AuthCtrl) PostLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	loginUrl, _ := router.UrlForRoute(names.RouteAuthLogin)
	acct, err := accounts.Find(username)

	if err != nil {
		ctrl.Error(w, r, err)
		return
	}
	if !acct.Auth(password) {
		err = errors.New(translate.Core(translate.Error, "invalid_login"))
		ctrl.Error(w, r, err)
		return
	}

	cfg, err := appcfg.ReadConfig()
	if err != nil {
		err = errors.New(translate.Core(translate.Error, "invalid_login"))
		ctrl.Error(w, r, err)
		return
	}

	payload := map[string]string{"username": username}
	token, err := jsonwebtoken.GenerateToken(payload, cfg.Secret)
	if err != nil {
		flash.SetFlashMsg(w, flash.Error, err.Error())
		http.Redirect(w, r, loginUrl, http.StatusFound)
		return
	}
	cookie.SetCookie(w, authTokenCookie, token)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (ctrl *AuthCtrl) GetLogout(w http.ResponseWriter, r *http.Request) {
	flash.SetFlashMsg(w, flash.Info, translate.Core(translate.Info, "logout_msg"))
	cookie.SetCookie(w, authTokenCookie, "")
	loginUrl, err := router.UrlForRoute(names.RouteAuthLogin)
	if err != nil {
		ctrl.Error(w, r, err)
		return
	}
	http.Redirect(w, r, loginUrl, http.StatusSeeOther)
}

func (ctrl *AuthCtrl) Error(w http.ResponseWriter, r *http.Request, err error) {
	errRoute := response.NewErrRoute(names.RouteAuthLogin)
	errRoute.Redirect(w, r, err)
}
