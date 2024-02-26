/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

import "net/http"

type VuePortalRoute struct {
	RouteName    string
	RoutePath    string
	Component    string
	HandlerFunc  http.HandlerFunc
	Middlewares  []func(http.Handler) http.Handler
}
