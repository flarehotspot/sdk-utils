package helpers

import (
	"net"
	"net/http"

	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/utils/hostfinder"
)

func CurrentClient(clntMgr *connmgr.ClientRegister, r *http.Request) (sdkconnmgr.ClientDevice, error) {
	clntSym := r.Context().Value(sdkhttp.ClientCtxKey)
	if clntSym != nil {
        clnt, ok := clntSym.(sdkconnmgr.ClientDevice)
        if ok {
            return clnt, nil
        }
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return nil, err
	}

	h, err := hostfinder.FindByIp(ip)
	if err != nil {
		return nil, err
	}

	clnt, err := clntMgr.Register(r.Context(), h.MacAddr, h.IpAddr, h.Hostname)
	if err != nil {
		return nil, err
	}

	return clnt, nil
}
