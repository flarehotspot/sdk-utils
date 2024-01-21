package plugins

import (
	"net/http"

	"github.com/flarehotspot/core/web/response"
)

const (
	rootjson = "$$response$$"
)

func NewVueResponse(vr *VueRouterApi, w http.ResponseWriter, r *http.Request) *VueResponse {
	return &VueResponse{w, r, vr, map[string]any{
		rootjson: map[string]any{},
	}}
}

type VueResponse struct {
	w      http.ResponseWriter
	r      *http.Request
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

func (self *VueResponse) JsonData(data any) {
	newdata := self.data[rootjson].(map[string]any)
	newdata["data"] = data
	data = map[string]any{
		rootjson: newdata,
	}
	response.Json(self.w, data, http.StatusOK)
}

func (res *VueResponse) Redirect(routename string, pairs ...string) {
	route, ok := res.router.FindVueRoute(routename)
	if !ok {
		response.ErrorJson(res.w, "Vue route \""+routename+"\" not found")
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

	response.Json(res.w, data, http.StatusOK)
}
