package sdkhttp

import (
	"net/http"

	sdkconnmgr "github.com/flarehotspot/flarehotspot/core/sdk/api/connmgr"
)

// HttpApi is used to process and respond to http requests.
type HttpApi interface {

	// Returns the auth API.
	Auth() HttpAuth

	// Returns helper methods for views and handlers.
	Helpers() HttpHelpers

	// Returns the router API.
	HttpRouter() HttpRouter

	// Returns the router API for vue requests.
	VueRouter() VueRouter

	// Returns the middlewares API.
	Middlewares() Middlewares

	// Returns the http response writer API.
	HttpResponse() HttpResponse

	// Returns the http response writer API for vue requests
	VueResponse() VueResponse

	// Returns the current client device from http request.
	GetDevice(r *http.Request) (sdkconnmgr.ClientDevice, error)

	// Returns the http variables in your routes. For example, if your route path is "/some/path/{varname}",
	// then you can get the value of "varname" by calling GetMuxVars(r)["varname"].
	MuxVars(r *http.Request) map[string]string

	// Returns the consolidated vue navigation list from all plugins for the admin dashboard.
	GetAdminNavs(r *http.Request) []AdminNavList

	// Returns the consolidated vue navigation list from all plugins for the portal.
	GetPortalItems(r *http.Request) []PortalItem
}
