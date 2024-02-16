package apictrl

import (
	"net/http"

	"github.com/flarehotspot/flarehotspot/core/accounts"
	"github.com/flarehotspot/flarehotspot/core/connmgr"
	"github.com/flarehotspot/sdk/api/http"
	sse "github.com/flarehotspot/sdk/utils/sse"
	"github.com/flarehotspot/flarehotspot/core/web/helpers"
)

type SseApiCtrl struct {
	reg *connmgr.ClientRegister
}

func NewSseApiCtrl(reg *connmgr.ClientRegister) *SseApiCtrl {
	return &SseApiCtrl{reg}
}

func (ctrl *SseApiCtrl) AdminEvents(w http.ResponseWriter, r *http.Request) {

	acctsym := r.Context().Value(sdkhttp.SysAcctCtxKey)
	acct := acctsym.(*accounts.Account)

	socket, err := sse.NewSocket(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sse.AddSocket(acct.Username(), socket)
	socket.Listen()
}

func (ctrl *SseApiCtrl) PortalEvents(w http.ResponseWriter, r *http.Request) {
	clnt, err := helpers.CurrentClient(ctrl.reg, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s, err := sse.NewSocket(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sse.AddSocket(clnt.MacAddr(), s)
	s.Listen()
}
