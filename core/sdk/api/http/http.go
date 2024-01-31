package sdkhttp

import (
	"net/http"

	sdkconnmgr "github.com/flarehotspot/core/sdk/api/connmgr"
)

// IHttp is used to process and respond to http requests.
type IHttp interface {
	Auth() IAuth

	// Returns the router API.
	HttpRouter() IHttpRouter

	VueRouter() IVueRouter

	Helpers() IHelpers

	// Returns the middlewares API.
	Middlewares() Middlewares

	// Returns the http response writer API.
	HttpResponse() IHttpResponse

	// Returns the http response writer API for vue requests
	VueResponse() IVueResponse

	// Returns the current client device from http request.
	GetDevice(r *http.Request) (sdkconnmgr.IClientDevice, error)

	// Returns the http variables in your routes. For example, if your route path is "/some/path/{varname}",
	// then you can get the value of "varname" by calling GetMuxVars(r)["varname"].
	MuxVars(r *http.Request) map[string]string

	// Returns the consolidated vue navigation list from all plugins for the admin dashboard.
	GetAdminNavs(r *http.Request) []AdminNavList

	// Returns the consolidated vue navigation list from all plugins for the portal.
	GetPortalItems(r *http.Request) []PortalItem
}
