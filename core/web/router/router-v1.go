package router

import (
	"errors"
	"fmt"
	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/gorilla/mux"
)

var (
	ApiRouterV1       *mux.Router
	AssetsApiRouterV1 *mux.Router
	AdminApiRouterV1  *mux.Router
	PortalApiRouterV1 *mux.Router
)

func init() {
	RootRouter = mux.NewRouter().StrictSlash(true)
	BootingRouter = mux.NewRouter().StrictSlash(true)
	PluginRouter = RootRouter.PathPrefix("/plugin").Subrouter()

	ApiRouterV1 = RootRouter.PathPrefix("/api/v1").Subrouter()
	AssetsApiRouterV1 = ApiRouterV1.PathPrefix("/assets").Subrouter()
	AdminApiRouterV1 = ApiRouterV1.PathPrefix("/admin").Subrouter()
	PortalApiRouterV1 = ApiRouterV1.PathPrefix("/portal").Subrouter()
}

func UrlForRoute(muxname router.MuxRouteName, pairs ...string) (string, error) {
	route := FindRoute(muxname)
	if route != nil {
		if url, err := route.URL(pairs...); err == nil {
			return url.EscapedPath(), nil
		}
	}
	return "", errors.New(fmt.Sprintf("Route name not found: \"%s\"", muxname))
}

func FindRoute(muxname router.MuxRouteName) *mux.Route {
	return RootRouter.Get(string(muxname))
}
