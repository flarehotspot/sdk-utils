package helpers

import (
	"errors"
	"net/http"

	"github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/utils/contexts"
)

func CurrentClient(r *http.Request) (connmgr.IClientDevice, error) {
	clntSym := r.Context().Value(contexts.ClientCtxKey)
	clnt, ok := clntSym.(connmgr.IClientDevice)
	if !ok {
		return nil, errors.New("Cannot convert nil to client device.")
	}
	return clnt, nil
}
