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
	response.Json(self.w, data, http.StatusOK)
}

func (res *VueResponse) Redirect(routename string, params any) {
	route, ok := res.router.FindVueRoute(routename)
	if !ok {
		response.ErrorJson(res.w, "Vue route \""+routename+"\" not found")
		return
	}

	data := map[string]any{
		"redirect":  true,
		"routename": route.VueRouteName,
		"params":    params,
	}

	response.Json(res.w, data, http.StatusOK)
}
