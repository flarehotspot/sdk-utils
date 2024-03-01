/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

type MuxRouteName string
type PluginRouteName string

type HttpRouterApi interface {

    // Returns a plugin router with authentication middleware.
	AdminRouter() HttpRouterInstance

    // Returns a generic plugin router.
	PluginRouter() HttpRouterInstance

	// Returns the url for the given route name.
	UrlForRoute(name PluginRouteName, pairs ...string) (url string)

    // Returns the url for the route from third-party plugins.
    // This is used create links to routes from other plugins.
    UrlForPkgRoute(pkg string, name string, pairs ...string) (url string)
}
