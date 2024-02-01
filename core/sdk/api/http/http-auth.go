package sdkhttp

import (
	"net/http"

	sdkacct "github.com/flarehotspot/core/sdk/api/accounts"
)

type IAuth interface {

	// Get the current admin user from the http request.
	CurrentAdmin(r *http.Request) (sdkacct.Account, error)

	AuthenticateAdmin(username string, password string) (sdkacct.Account, error)

	// Sets the auth-token cookie in response header
	SignInAdmin(w http.ResponseWriter, acct sdkacct.Account) error

	// Sets empty auth-token cooke response header
	SignOutAdmin(w http.ResponseWriter) error
}
