package sdkhttp

import (
	"net/http"
)

// IHttpApi is used to process and respond to http requests.
type IHttpApi interface {
	Auth() IAuthApi

	// Returns the router API.
	HttpRouter() IHttpRouterApi

	VueRouter() IVueRouterApi

	Helpers() IHelpers

	// Returns the middlewares API.
	Middlewares() Middlewares

	// Returns the http response writer API.
	HttpResponse() IHttpResponse

	// Returns the http response writer API for vue requests
	VueResponse() IVueResponse

	// Returns the http variables in your routes. For example, if your route path is "/some/path/{varname}",
	// then you can get the value of "varname" by calling GetMuxVars(r)["varname"].
	MuxVars(r *http.Request) map[string]string

	// Returns the consolidated vue navigation list from all plugins for the admin dashboard.
	GetAdminNavs(r *http.Request) []AdminNavCategory
}
