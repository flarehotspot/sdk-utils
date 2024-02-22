package plugins

import (
	"errors"
	"net/http"

	"github.com/flarehotspot/core/internal/accounts"
	"github.com/flarehotspot/core/internal/config"
	acct "github.com/flarehotspot/sdk/api/accounts"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	"github.com/flarehotspot/core/internal/utils/jsonwebtoken"
	"github.com/flarehotspot/core/internal/web/helpers"
	"github.com/flarehotspot/core/internal/web/middlewares"
)

func NewAuthApi(api *PluginApi) *AuthApi {
	return &AuthApi{
		api: api,
	}
}

type AuthApi struct {
	api *PluginApi
}

func (self *AuthApi) CurrentAcct(r *http.Request) (acct.Account, error) {
	return helpers.CurrentAcct(r)
}

func (self *AuthApi) Authenticate(username string, password string) (acct.Account, error) {
	acct, err := accounts.Find(username)
	if err != nil {
		return nil, err
	}

	if !acct.Auth(password) {
		return nil, errors.New(self.api.CoreAPI.Utl.Translate("error", "invalid_login"))
	}

	return acct, nil
}

func (self *AuthApi) SignIn(w http.ResponseWriter, acct acct.Account) error {
	appcfg, err := config.ReadApplicationConfig()
	if err != nil {
		return err
	}

	payload := map[string]string{"username": acct.Username()}
	token, err := jsonwebtoken.GenerateToken(payload, appcfg.Secret)
	if err != nil {
		return err
	}

	sdkhttp.SetCookie(w, middlewares.AuthTokenCookie, token)
	return nil
}

func (self *AuthApi) SignOut(w http.ResponseWriter) error {
	sdkhttp.SetCookie(w, middlewares.AuthTokenCookie, "")
	return nil
}
