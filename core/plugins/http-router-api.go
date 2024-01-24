package plugins

import (
	"fmt"
	"log"

	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
)

type HttpRouterApi struct {
	api          *PluginApi
	adminRouter  *HttpRouter
	pluginRouter *HttpRouter
}

func NewHttpRouterApi(api *PluginApi, db *db.Database, clnt *connmgr.ClientRegister) *HttpRouterApi {
	prefix := fmt.Sprintf("/%s/%s", api.Pkg(), api.Version())
	pluginMux := router.PluginRouter.PathPrefix(prefix).Subrouter()

	adminMux := pluginMux.PathPrefix("/admin").Subrouter()
	adminMux.Use(middlewares.AdminAuth)

	pluginRouter := &HttpRouter{api, pluginMux}
	adminRouter := &HttpRouter{api, adminMux}

	return &HttpRouterApi{api, adminRouter, pluginRouter}
}

func (self *HttpRouterApi) AdminRouter() sdkhttp.IHttpRouter {
	return self.adminRouter
}

func (self *HttpRouterApi) PluginRouter() sdkhttp.IHttpRouter {
	return self.pluginRouter
}

func (self *HttpRouterApi) MuxRouteName(name sdkhttp.PluginRouteName) sdkhttp.MuxRouteName {
	muxname := fmt.Sprintf("%s.%s", self.api.Pkg(), string(name))
	return sdkhttp.MuxRouteName(muxname)
}

func (self *HttpRouterApi) UrlForMuxRoute(muxname sdkhttp.MuxRouteName, pairs ...string) string {
	route := router.RootRouter.Get(string(muxname))
	if route == nil {
		log.Println("Error: route not found for " + string(muxname))
		return "Error: route not found for " + string(muxname)
	}

	url, err := route.URL(pairs...)
	if err != nil {
		log.Println("Error: " + err.Error())
		return ""
	}

	return url.String()
}

func (self *HttpRouterApi) UrlForRoute(name sdkhttp.PluginRouteName, pairs ...string) string {
	muxname := self.MuxRouteName(name)
	return self.UrlForMuxRoute(muxname, pairs...)
}
