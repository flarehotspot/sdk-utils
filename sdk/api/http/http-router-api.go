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

	// Returns the muxnmame for the route name in your plugin.
	// "muxname" is a route name that can be used for the UrlForMuxRoute() method.
	MuxRouteName(name PluginRouteName) (muxname MuxRouteName)

	// Returns the url for the mux route.
	// The difference between UrlForMuxRoute() vs UrlForRoute() is that UrlForRoute() only accepts route names specific to your plugin.
	UrlForMuxRoute(name MuxRouteName, pairs ...string) (url string)

	// Returns the url for the route.
	// The difference between UrlForMuxRoute() vs UrlForRoute() is that UrlForMuxRoute() only accepts route names built-in to the core system.
	UrlForRoute(name PluginRouteName, pairs ...string) (url string)
}
