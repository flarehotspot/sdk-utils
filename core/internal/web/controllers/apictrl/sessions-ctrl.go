package apictrl

// import (
// 	"net/http"

// 	"github.com/flarehotspot/core/internal/plugins"
// 	sdkmdls "github.com/flarehotspot/sdk/api/models"
// 	"github.com/flarehotspot/core/internal/web/helpers"
// 	"github.com/flarehotspot/core/internal/web/response"
// )

// type SessionsApiCtrl struct {
// 	g *plugins.CoreGlobals
// }

// func NewSessionsApiCtrl(g *plugins.CoreGlobals) *SessionsApiCtrl {
// 	return &SessionsApiCtrl{g}
// }

// func (self *SessionsApiCtrl) Index(w http.ResponseWriter, r *http.Request) {
// 	clnt, err := helpers.CurrentClient(r)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response.Json(w, map[string]any{"error": err.Error()}, 500)
// 	}

// 	sessions, err := self.g.Models.Session().SessionsForDev(r.Context(), clnt.Id())
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response.Json(w, map[string]any{"error": err.Error()}, 500)
// 	}

// 	var currmap map[string]any
// 	curr, ok := self.g.ClientMgr.CurrSession(clnt)
// 	if ok {
// 		currmap = sdkmdls.SessionToMap(curr.SessionModel())
// 	}

// 	smap := []map[string]any{}
// 	for _, s := range sessions {
// 		m := sdkmdls.SessionToMap(s)
// 		smap = append(smap, m)
// 	}

// 	data := map[string]any{
// 		"current":  currmap,
// 		"sessions": smap,
// 	}

// 	response.Json(w, data, 200)
// }
