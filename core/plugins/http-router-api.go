package plugins

import (
	"fmt"
	"log"

	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/sdk/api/http/router"
	routerI "github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/web/middlewares"
	coreR "github.com/flarehotspot/core/web/router"
)

type HttpRouterApi struct {
	api          *PluginApi
	adminRouter  *HttpRouter
	pluginRouter *HttpRouter
}

func NewRouterApi(api *PluginApi, db *db.Database, clnt *connmgr.ClientRegister) *HttpRouterApi {
	prefix := fmt.Sprintf("/%s/%s", api.Pkg(), api.Version())
	pluginMux := coreR.PluginRouter.PathPrefix(prefix).Subrouter()
	adminMux := pluginMux.PathPrefix("/admin").Subrouter()

	authMw := middlewares.AdminAuth
	adminMux.Use(authMw)

	pluginRouter := &HttpRouter{api, pluginMux}
	adminRouter := &HttpRouter{api, adminMux}

	return &HttpRouterApi{api, adminRouter, pluginRouter}
}

func (self *HttpRouterApi) AdminRouter() routerI.IHttpRouter {
	return self.adminRouter
}

func (self *HttpRouterApi) PluginRouter() routerI.IHttpRouter {
	return self.pluginRouter
}

func (self *HttpRouterApi) MuxRouteName(name router.PluginRouteName) router.MuxRouteName {
	muxname := fmt.Sprintf("%s.%s", self.api.Pkg(), string(name))
	return router.MuxRouteName(muxname)
}

func (util *HttpRouterApi) UrlForMuxRoute(muxname router.MuxRouteName, pairs ...string) string {
	route := coreR.RootRouter.Get(string(muxname))
	url, err := route.URL(pairs...)
	if err != nil {
		log.Println("Error: " + err.Error())
		return ""
	}

	return url.String()
}

func (util *HttpRouterApi) UrlForRoute(name router.PluginRouteName, pairs ...string) string {
	muxname := util.MuxRouteName(name)
	return util.UrlForMuxRoute(muxname, pairs...)
}
