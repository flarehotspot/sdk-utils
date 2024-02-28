/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

type MuxRouteName string
type PluginRouteName string

type HttpRouterApi interface {
	AdminRouter() HttpRouterInstance
	PluginRouter() HttpRouterInstance
	MuxRouteName(name PluginRouteName) (muxname MuxRouteName)
	UrlForMuxRoute(name MuxRouteName, pairs ...string) (url string)
	UrlForRoute(name PluginRouteName, pairs ...string) (url string)
}
