package plugins

import (
	"github.com/flarehotspot/core/sdk/api/http"
	"github.com/gorilla/mux"
)

type HttpRoute struct {
	api   *PluginApi
	route *mux.Route
}

func (r *HttpRoute) Name(name sdkhttp.PluginRouteName) {
	muxname := r.api.HttpApi().HttpRouter().MuxRouteName(name)
	r.route.Name(string(muxname))
}

func NewPluginRoute(api *PluginApi, r *mux.Route) *HttpRoute {
	return &HttpRoute{api, r}
}
