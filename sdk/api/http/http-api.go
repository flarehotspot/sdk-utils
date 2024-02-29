/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

import (
	"net/http"

	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
)

// HttpApi is used to process and respond to http requests.
type HttpApi interface {

	// Returns the auth API.
	Auth() HttpAuth

	// Returns helper methods for views and handlers.
	Helpers() HttpHelpers

	// Returns the router API.
	HttpRouter() HttpRouterApi

	// Returns the middlewares API.
	Middlewares() Middlewares

	// Returns the http response writer API.
	HttpResponse() HttpResponse

	// Returns the router API for vue requests.
	VueRouter() VueRouterApi

	// Returns the http response writer API for vue requests
	VueResponse() VueResponse

	// Returns the current client device from http request.
	GetClientDevice(r *http.Request) (clnt sdkconnmgr.ClientDevice, err error)

	// Returns the http variables in your routes. For example, if your route path is "/some/path/{varname}",
	// then you can get the value of "varname" by calling GetMuxVars(r)["varname"].
	MuxVars(r *http.Request) map[string]string

	// Returns the consolidated vue navigation list from all plugins for the admin dashboard.
	GetAdminNavs(r *http.Request) []AdminNavList

	// Returns the consolidated vue navigation list from all plugins for the portal.
	GetPortalItems(r *http.Request) []PortalItem
}
