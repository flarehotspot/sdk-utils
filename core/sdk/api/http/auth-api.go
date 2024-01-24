package sdkhttp

import (
	"net/http"

	sdkacct "github.com/flarehotspot/core/sdk/api/accounts"
)

type IAuthApi interface {

	// Get the current admin user from the http request.
	CurrentAdmin(r *http.Request) (sdkacct.IAccount, error)

	AuthenticateAdmin(username string, password string) (sdkacct.IAccount, error)

	SignInAdmin(w http.ResponseWriter, acct sdkacct.IAccount)

	SignOutAdmin(w http.ResponseWriter)
}
