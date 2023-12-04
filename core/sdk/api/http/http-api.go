package http

import (
	"net/http"

	"github.com/flarehotspot/core/sdk/api/http/middlewares"
	"github.com/flarehotspot/core/sdk/api/http/response"
	"github.com/flarehotspot/core/sdk/api/http/router"
)

// IHttpApi is used to process and respond to http requests.
type IHttpApi interface {

	// Returns the router API.
	Router() router.IRouterApi

	// Returns the middlewares API.
	Middlewares() middlewares.Middlewares

	// Returns the http response API.
	Respond() response.IHttpResponse

	// Returns the http variables in your routes. For example, if your route path is "/some/path/{varname}",
	// then you can get the value of "varname" by calling GetMuxVars(r)["varname"].
	MuxVars(r *http.Request) map[string]string

	// Set a global template.FuncMap{} available for your view templates.
	// Note that these custom function maps are not available in the layout templates.
	ViewFuncMap(fmap map[string]func())
}
