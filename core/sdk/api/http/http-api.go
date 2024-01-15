package http

import (
	"net/http"

	"github.com/flarehotspot/core/sdk/api/http/middlewares"
	"github.com/flarehotspot/core/sdk/api/http/response"
	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/sdk/api/http/views"
)

// IHttpApi is used to process and respond to http requests.
type IHttpApi interface {

	// Returns the router API.
	HttpRouter() router.IHttpRouterApi

	VueRouter() router.IVueRouterApi

	Helpers(w http.ResponseWriter, r *http.Request) views.IViewHelpers

	// Returns the http uri path for the given asset.
	AssetPath(assetPath string) string

	// Returns the middlewares API.
	Middlewares() middlewares.Middlewares

	// Returns the http response API.
	Respond() response.IHttpResponse

	// Returns the http variables in your routes. For example, if your route path is "/some/path/{varname}",
	// then you can get the value of "varname" by calling GetMuxVars(r)["varname"].
	MuxVars(r *http.Request) map[string]string
}
