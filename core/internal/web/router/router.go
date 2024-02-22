package router

import (
	"errors"
	"fmt"

	"github.com/flarehotspot/sdk/api/http"
	"github.com/gorilla/mux"
)

const (
	NotFoundRoute string = "/404"
)

var (
	RootRouter    *mux.Router
	BootingRouter *mux.Router
	PluginRouter  *mux.Router
	AssetsRouter  *mux.Router
)

var (
	ApiRouter       *mux.Router
	AdminApiRouter  *mux.Router
	PortalApiRouter *mux.Router
	AuthApiRouter   *mux.Router
)

func init() {
	RootRouter = mux.NewRouter().StrictSlash(true)
	BootingRouter = mux.NewRouter().StrictSlash(true)
	PluginRouter = RootRouter.PathPrefix("/plugin").Subrouter()
	AssetsRouter = RootRouter.PathPrefix("/assets").Subrouter()

	ApiRouter = RootRouter.PathPrefix("/api").Subrouter()
	AdminApiRouter = ApiRouter.PathPrefix("/admin").Subrouter()
	PortalApiRouter = ApiRouter.PathPrefix("/portal").Subrouter()
	AuthApiRouter = ApiRouter.PathPrefix("/auth").Subrouter()
}

func UrlForRoute(muxname sdkhttp.MuxRouteName, pairs ...string) (string, error) {
	route := FindRoute(muxname)
	if route != nil {
		if url, err := route.URL(pairs...); err == nil {
			return url.EscapedPath(), nil
		}
	}
	return "", errors.New(fmt.Sprintf("Route name not found: \"%s\"", muxname))
}

func FindRoute(muxname sdkhttp.MuxRouteName) *mux.Route {
	return RootRouter.Get(string(muxname))
}
