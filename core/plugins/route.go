package plugins

import (
	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/gorilla/mux"
)

type PluginRoute struct {
  api *PluginApi
	route *mux.Route
}

func (r *PluginRoute) Name(name router.PluginRouteName) {
  muxname := r.api.HttpApi().Router().MuxRouteName(name)
	r.route.Name(string(muxname))
}

func NewPluginRoute(api *PluginApi, r *mux.Route) *PluginRoute {
	return &PluginRoute{api, r}
}
