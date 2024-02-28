/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

import "net/http"

const (
	VueNotFoundPath = "/404"
)

type VueAdminNavsFunc func(r *http.Request) []VueAdminNav
type VuePortalItemsFunc func(r *http.Request) []VuePortalItem

// VueRouterApi is used to create navigation items in the application.
type VueRouterApi interface {
	// Set admin vue route
	RegisterAdminRoutes(routes ...VueAdminRoute)

	// Set portal vue routes
	RegisterPortalRoutes(routes ...VuePortalRoute)

	// Used to register a function that returns a slice of admin navs.
	// Items returned from this function is added to the admin navigation menu.
	AdminNavsFunc(VueAdminNavsFunc)

	// Used to register a function that returns a slice of *PortalNavItem.
	// Items returned from this function is added to the captive portal navigation items.
	PortalItemsFunc(VuePortalItemsFunc)
}
