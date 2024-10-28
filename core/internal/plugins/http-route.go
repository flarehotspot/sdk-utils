package plugins

import (
	"sdk/api/http"
	"github.com/gorilla/mux"
)

func NewHttpRoute(api *PluginApi, r *mux.Route) *HttpRoute {
	return &HttpRoute{api, r}
}

type HttpRoute struct {
	api   *PluginApi
	route *mux.Route
}

func (self *HttpRoute) Name(name sdkhttp.PluginRouteName) {
	muxname := self.api.HttpAPI.httpRouter.MuxRouteName(name)
	self.route.Name(string(muxname))
}
