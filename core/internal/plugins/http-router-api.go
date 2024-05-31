package plugins

import (
	"fmt"
	"log"
	"net/http"

	"core/internal/connmgr"
	"core/internal/db"
	"core/internal/web/middlewares"
	"core/internal/web/router"
	sdkhttp "sdk/api/http"
	"github.com/gorilla/mux"
)

type HttpRouterApi struct {
	api          *PluginApi
	adminRouter  *HttpRouterInstance
	pluginRouter *HttpRouterInstance
}

func NewHttpRouterApi(api *PluginApi, db *db.Database, clnt *connmgr.ClientRegister) *HttpRouterApi {
	prefix := fmt.Sprintf("/%s/%s", api.Pkg(), api.Version())
	pluginMux := router.PluginRouter.PathPrefix(prefix).Subrouter()

	adminMux := pluginMux.PathPrefix("/admin").Subrouter()
	adminMux.Use(middlewares.AdminAuth)

	pluginRouter := &HttpRouterInstance{api, pluginMux}
	adminRouter := &HttpRouterInstance{api, adminMux}

	return &HttpRouterApi{api, adminRouter, pluginRouter}
}

func (self *HttpRouterApi) AdminRouter() sdkhttp.HttpRouterInstance {
	return self.adminRouter
}

func (self *HttpRouterApi) PluginRouter() sdkhttp.HttpRouterInstance {
	return self.pluginRouter
}

func (self *HttpRouterApi) UseMiddleware(middleware ...func(http.Handler) http.Handler) {
	for _, mw := range middleware {
		router.RootRouter.Use(mux.MiddlewareFunc(mw))
	}
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

func (self *HttpRouterApi) UrlForPkgRoute(pkg string, name string, pairs ...string) string {
	otherPkg, ok := self.api.PluginsMgrApi.FindByPkg(pkg)
	if !ok {
		return ""
	}
	return otherPkg.Http().Helpers().UrlForRoute(name, pairs...)
}
