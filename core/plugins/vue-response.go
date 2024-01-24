package plugins

import (
	"net/http"

	"github.com/flarehotspot/core/web/response"
)

const (
	rootjson = "$$response"
)

func NewVueResponse(vr *VueRouterApi) *VueResponse {
	return &VueResponse{vr, map[string]any{
		rootjson: map[string]any{},
	}}
}

type VueResponse struct {
	router *VueRouterApi
	data   map[string]any
}

func (self *VueResponse) FlashMsg(msgType string, msg string) {
	newdata := self.data[rootjson].(map[string]any)
	newdata["flash"] = map[string]string{
		"type": msgType, // "success", "error", "warning", "info
		"msg":  msg,
	}
	self.data[rootjson] = newdata
}

func (self *VueResponse) Json(w http.ResponseWriter, data any, status int) {
	newdata := self.data[rootjson].(map[string]any)
	newdata["data"] = data
	data = map[string]any{
		rootjson: newdata,
	}
	response.Json(w, data, status)
}

func (res *VueResponse) Redirect(w http.ResponseWriter, routename string, pairs ...string) {
	route, ok := res.router.FindVueRoute(routename)
	if !ok {
		response.ErrorJson(w, "Vue route \""+routename+"\" not found", 500)
		return
	}

	params := map[string]string{}
	for i := 0; i < len(pairs); i += 2 {
		params[pairs[i]] = pairs[i+1]
	}

	newdata := res.data[rootjson].(map[string]any)
	newdata["redirect"] = true
	newdata["route_name"] = route.VueRouteName
	newdata["params"] = params
	data := map[string]any{rootjson: newdata}

	response.Json(w, data, http.StatusOK)
}
