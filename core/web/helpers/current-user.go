package helpers

import (
	"errors"
	"net/http"

	"github.com/flarehotspot/core/accounts"
	"github.com/flarehotspot/core/sdk/api/http"
)

func CurrentAdmin(r *http.Request) (*accounts.Account, error) {
	sym := r.Context().Value(sdkhttp.SysAcctCtxKey)
	acct, ok := sym.(*accounts.Account)
	if !ok {
		return nil, errors.New("Can't determine current admin account.")
	}

	return acct, nil
}
