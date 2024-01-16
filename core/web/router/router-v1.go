package router

import (
	"github.com/gorilla/mux"
)

var (
	ApiRouterV1       *mux.Router
	AssetsApiRouterV1 *mux.Router
	AdminApiRouterV1  *mux.Router
	PortalApiRouterV1 *mux.Router
	AuthApiRouterV1   *mux.Router
)

func init() {
	RootRouter = mux.NewRouter().StrictSlash(true)
	BootingRouter = mux.NewRouter().StrictSlash(true)
	PluginRouter = RootRouter.PathPrefix("/plugin").Subrouter()

	ApiRouterV1 = RootRouter.PathPrefix("/api/v1").Subrouter()
	AuthApiRouterV1 = ApiRouterV1.PathPrefix("/auth").Subrouter()
	AssetsApiRouterV1 = ApiRouterV1.PathPrefix("/assets").Subrouter()
	AdminApiRouterV1 = ApiRouterV1.PathPrefix("/admin").Subrouter()
	PortalApiRouterV1 = ApiRouterV1.PathPrefix("/portal").Subrouter()
}
