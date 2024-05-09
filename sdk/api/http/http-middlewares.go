/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

import "net/http"

// Middlewares contains http middlewares for admin authentication, mobile device details, etc.
type Middlewares interface {
	// Returns the middleware for admin authentication.
	AdminAuth() func(next http.Handler) http.Handler

	// Returns the middleware for caching the response. It forces browsers to cache the response for n number of days.
	CacheResponse(days int) func(next http.Handler) http.Handler

	// Returns the middleware that checks pending purchases
	PendingPurchase() func(next http.Handler) http.Handler
}
