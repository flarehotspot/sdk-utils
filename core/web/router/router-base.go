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
	RootRouter    *mux.Router
	BootingRouter *mux.Router
	PluginRouter  *mux.Router
)

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
