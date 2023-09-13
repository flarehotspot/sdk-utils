package apictrl

import (
	"net/http"

	"github.com/flarehotspot/core/accounts"
	"github.com/flarehotspot/core/web/helpers"
	"github.com/flarehotspot/core/sdk/utils/contexts"
	"github.com/flarehotspot/core/sdk/utils/sse"
)

type SseApiCtrl struct{}

func NewSseApiCtrl() *SseApiCtrl {
	return &SseApiCtrl{}
}

func (ctrl *SseApiCtrl) AdminEvents(w http.ResponseWriter, r *http.Request) {

	acctsym := r.Context().Value(contexts.SysAcctCtxKey)
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
	clnt, err := helpers.CurrentClient(r)
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
