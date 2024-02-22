package helpers

import (
	"errors"
	"net/http"

	"github.com/flarehotspot/core/internal/accounts"
	"github.com/flarehotspot/sdk/api/http"
)

func CurrentAcct(r *http.Request) (*accounts.Account, error) {
	sym := r.Context().Value(sdkhttp.SysAcctCtxKey)
	acct, ok := sym.(*accounts.Account)
	if !ok {
		return nil, errors.New("Can't determine current admin account.")
	}

	return acct, nil
}
