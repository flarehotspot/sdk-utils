package router

import (
	"errors"
	"fmt"
	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/gorilla/mux"
)

const (
	NotFoundRoute string = "/404"
)

var (
	bootingRouter *mux.Router
	rootRouter    *mux.Router
	adminRouter   *mux.Router
	pluginRouter  *mux.Router
)

func init() {
	bootingRouter = mux.NewRouter().StrictSlash(true)
	rootRouter = mux.NewRouter().StrictSlash(true)
	pluginRouter = rootRouter.PathPrefix("/plugin").Subrouter()
	adminRouter = rootRouter.PathPrefix("/admin").Subrouter()
}

func BootingRrouter() *mux.Router {
	return bootingRouter
}

func RootRouter() *mux.Router {
	return rootRouter
}

func AdminRouter() *mux.Router {
	return adminRouter
}

func PluginRouter() *mux.Router {
	return pluginRouter
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
	return rootRouter.Get(string(muxname))
}
