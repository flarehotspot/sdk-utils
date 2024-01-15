package plugins

import (
	"log"

	"github.com/flarehotspot/core/sdk/api/http/router"
	IRouter "github.com/flarehotspot/core/sdk/api/http/router"
	coreRouter "github.com/flarehotspot/core/web/router"
	"github.com/gorilla/mux"
)

type HttpRouterApi struct {
	api          *PluginApi
	adminRouter  *HttpRouter
	pluginRouter *HttpRouter
}

func NewRouterApi(api *PluginApi) *HttpRouterApi {
	pluginMux := coreRouter.PluginRouter.PathPrefix("/" + api.slug).Subrouter()
	pluginRouter := &HttpRouter{api, pluginMux}

	adminMux := coreRouter.AdminApiRouterV1.PathPrefix("/plugin/" + api.slug).Subrouter()
	adminRouter := &HttpRouter{api, adminMux}

	return &HttpRouterApi{api, adminRouter, pluginRouter}
}

func (self *HttpRouterApi) AdminRouter() IRouter.IHttpRouter {
	return self.adminRouter
}

func (self *HttpRouterApi) PluginRouter() IRouter.IHttpRouter {
	return self.pluginRouter
}

func (self *HttpRouterApi) MuxRouteName(name router.PluginRouteName) router.MuxRouteName {
	muxname := self.api.slug + "::" + string(name)
	return router.MuxRouteName(muxname)
}

func (util *HttpRouterApi) UrlForMuxRoute(muxname router.MuxRouteName, pairs ...string) string {
	var url string
	coreRouter.RootRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
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

func (util *HttpRouterApi) UrlForRoute(name router.PluginRouteName, pairs ...string) string {
	muxname := util.MuxRouteName(name)
	return util.UrlForMuxRoute(muxname, pairs...)
}
