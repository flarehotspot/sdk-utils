package plugins

import (
	"net/http"

	"github.com/flarehotspot/core/web/response"
)

func NewVueResponse(vr *VueRouterApi, w http.ResponseWriter, r *http.Request) *VueResponse {
	return &VueResponse{w, r, vr}
}

type VueResponse struct {
	w      http.ResponseWriter
	r      *http.Request
	router *VueRouterApi
}

func (self *VueResponse) JsonData(data any) {
	data = map[string]any{
		"redirect": false,
		"data":     data,
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

	data := map[string]any{
		"redirect": true,
		"name":     route.VueRouteName,
		"params":   params,
	}

	response.Json(res.w, data, http.StatusOK)
}
