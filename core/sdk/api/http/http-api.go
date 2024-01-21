package sdkhttp

import (
	"net/http"
)

// IHttpApi is used to process and respond to http requests.
type IHttpApi interface {

	// Returns the router API.
	HttpRouter() IHttpRouterApi

	VueRouter() IVueRouterApi

	Helpers() IHelpers

	// Returns the middlewares API.
	Middlewares() Middlewares

	// Returns the http response API.
	HttpResponse() IHttpResponse

	// Returns the http variables in your routes. For example, if your route path is "/some/path/{varname}",
	// then you can get the value of "varname" by calling GetMuxVars(r)["varname"].
	MuxVars(r *http.Request) map[string]string
}
