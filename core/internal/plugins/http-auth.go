package plugins

import (
	"errors"
	"net/http"

	"core/internal/accounts"
	"core/internal/config"
	"core/internal/utils/jsonwebtoken"
	webutil "core/internal/utils/web"
	"core/internal/web/helpers"
	sdkacct "sdk/api/accounts"
	sdkhttp "sdk/api/http"
)

func NewHttpAuth(api *PluginApi) *HttpAuth {
	return &HttpAuth{
		api: api,
	}
}

type HttpAuth struct {
	api *PluginApi
}

func (self *HttpAuth) CurrentAcct(r *http.Request) (sdkacct.Account, error) {
	return helpers.CurrentAcct(r)
}

func (self *HttpAuth) Authenticate(username string, password string) (sdkacct.Account, error) {
	acct, err := accounts.Find(username)
	if err != nil {
		return nil, err
	}

	if !acct.Auth(password) {
		return nil, errors.New(self.api.CoreAPI.Utl.Translate("error", "invalid_login"))
	}

	return acct, nil
}

func (self *HttpAuth) SignIn(w http.ResponseWriter, acct sdkacct.Account) error {
	appcfg, err := config.ReadApplicationConfig()
	if err != nil {
		return err
	}

	payload := map[string]string{"username": acct.Username()}
	token, err := jsonwebtoken.GenerateToken(payload, appcfg.Secret)
	if err != nil {
		return err
	}

	sdkhttp.SetCookie(w, webutil.AuthTokenCookie, token)
	return nil
}

func (self *HttpAuth) SignOut(w http.ResponseWriter) error {
	sdkhttp.SetCookie(w, webutil.AuthTokenCookie, "")
	return nil
}
