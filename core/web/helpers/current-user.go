package helpers

import (
	"errors"
	"net/http"

	"github.com/flarehotspot/core/accounts"
	"github.com/flarehotspot/core/sdk/utils/contexts"
)

func CurrentAdmin(r *http.Request) (*accounts.Account, error) {
	sym := r.Context().Value(contexts.SysAcctCtxKey)
	acct, ok := sym.(*accounts.Account)
	if !ok {
		return nil, errors.New("Can't determine current admin account.")
	}

	return acct, nil
}
