package sdkhttp

import (
	"net/http"

	sdkacct "github.com/flarehotspot/sdk/api/accounts"
)

type HttpAuth interface {

	// Get the current admin user from the http request.
	CurrentAcct(r *http.Request) (sdkacct.Account, error)

	Authenticate(username string, password string) (sdkacct.Account, error)

	// Sets the auth-token cookie in response header
	SignIn(w http.ResponseWriter, acct sdkacct.Account) error

	// Sets empty auth-token cooke response header
	SignOut(w http.ResponseWriter) error
}
