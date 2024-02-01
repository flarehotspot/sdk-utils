package helpers

import (
	"errors"
	"net/http"

	connmgr "github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/api/http"
)

func CurrentClient(r *http.Request) (connmgr.ClientDevice, error) {
	clntSym := r.Context().Value(sdkhttp.ClientCtxKey)
	clnt, ok := clntSym.(connmgr.ClientDevice)
	if !ok {
		return nil, errors.New("Cannot convert nil to client device.")
	}
	return clnt, nil
}
