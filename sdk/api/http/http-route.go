/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

// HttpRoute represents a single route in the router.
type HttpRoute interface {
	// Returns the name of the route.
	Name(name PluginRouteName)
}
