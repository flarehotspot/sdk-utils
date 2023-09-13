package plugins

import (
	"log"

	coreRouter "github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/sdk/api/http/router"
	IRouter "github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/gorilla/mux"
)

type RouterApi struct {
	api          *PluginApi
	adminRouter  *PluginRouter
	pluginRouter *PluginRouter
}

func NewRouterApi(api *PluginApi) *RouterApi {
	pluginMux := coreRouter.PluginRouter().PathPrefix("/" + api.slug).Subrouter()
	pluginRouter := PluginRouter{api, pluginMux}

	adminMux := coreRouter.AdminRouter().PathPrefix("/plugin/" + api.slug).Subrouter()
	adminRouter := PluginRouter{api, adminMux}
	return &RouterApi{api, &adminRouter, &pluginRouter}
}

func (self *RouterApi) AdminRouter() IRouter.IRouter {
	return self.adminRouter
}

func (self *RouterApi) PluginRouter() IRouter.IRouter {
	return self.pluginRouter
}

func (self *RouterApi) MuxRouteName(name router.PluginRouteName) router.MuxRouteName {
	muxname := self.api.slug + "::" + string(name)
	return router.MuxRouteName(muxname)
}

func (util *RouterApi) UrlForMuxRoute(muxname router.MuxRouteName, pairs ...string) string {
	var url string
	coreRouter.RootRouter().Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		if url == "" && route.GetName() == string(muxname) {
			result, err := route.URL(pairs...)
			if err != nil {
				log.Println(err)
				return nil
			}
			url = result.String()
		}
		return nil
	})
	return url
}

func (util *RouterApi) UrlForRoute(name router.PluginRouteName, pairs ...string) string {
	muxname := util.MuxRouteName(name)
	return util.UrlForMuxRoute(muxname, pairs...)
}
