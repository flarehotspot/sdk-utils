package plugins

import (
	"errors"
	"net/http"

	"github.com/flarehotspot/core/accounts"
	"github.com/flarehotspot/core/config"
	acct "github.com/flarehotspot/core/sdk/api/accounts"
	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/utils/jsonwebtoken"
	"github.com/flarehotspot/core/web/helpers"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/response"
)

func NewAuthApi(api *PluginApi) *AuthApi {
	return &AuthApi{
		api: api,
	}
}

type AuthApi struct {
	api *PluginApi
}

func (self *AuthApi) CurrentAdmin(r *http.Request) (acct.IAccount, error) {
	return helpers.CurrentAdmin(r)
}

func (self *AuthApi) AuthenticateAdmin(username string, password string) (acct.IAccount, error) {
	acct, err := accounts.Find(username)
	if err != nil {
		return nil, err
	}

	if !acct.Auth(password) {
		return nil, errors.New(self.api.coreApi.Utl.Translate("error", "invalid_login"))
	}

	return acct, nil
}

func (self *AuthApi) SignInAdmin(w http.ResponseWriter, acct acct.IAccount) {
	appcfg, err := config.ReadApplicationConfig()
	if err != nil {
		msg := self.api.coreApi.Utl.Translate("error", "app_cfg_error")
		response.ErrorJson(w, msg, 500)
		return
	}

	payload := map[string]string{"username": acct.Username()}
	token, err := jsonwebtoken.GenerateToken(payload, appcfg.Secret)
	if err != nil {
		response.ErrorJson(w, err.Error(), 500)
		return
	}

	sdkhttp.SetCookie(w, middlewares.AuthTokenCookie, token)
	data := map[string]string{"token": token}
	self.api.HttpAPI.VueResponse().Json(w, data, 200)
}

func (self *AuthApi) SignOutAdmin(w http.ResponseWriter) {
	sdkhttp.SetCookie(w, middlewares.AuthTokenCookie, "")
	data := map[string]string{"message": "Logout success"}
	self.api.HttpApi().HttpResponse().Json(w, data, http.StatusOK)
}
