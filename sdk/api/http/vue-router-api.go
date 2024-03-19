/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

import (
	sdkacct "github.com/flarehotspot/sdk/api/accounts"
	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
)

const (
	VueNotFoundPath = "/404"
)

type VueAdminNavsFunc func(acct sdkacct.Account) []VueAdminNav
type VuePortalItemsFunc func(clnt sdkconnmgr.ClientDevice) []VuePortalItem

// VueRouterApi is used to create navigation items in the application.
type VueRouterApi interface {

	// Set portal vue routes
	RegisterPortalRoutes(routes ...VuePortalRoute)

	// Set admin vue route
	RegisterAdminRoutes(routes ...VueAdminRoute)

	// Used to register a function that returns a slice of admin navs.
	// Items returned from this function is added to the admin navigation menu.
	AdminNavsFunc(VueAdminNavsFunc)

	// Used to register a function that returns a slice of *PortalNavItem.
	// Items returned from this function is added to the captive portal navigation items.
	PortalItemsFunc(VuePortalItemsFunc)

	// Returns the vue route name for a named route which can be used in vue router, e.g.
	//   $this.push({name: '<% .Helpers.VueRouteName "login" %>'})
	VueRouteName(name string) string

	// Returns the vue route path for a named route
	VueRoutePath(name string, pairs ...string) string

	// Returns the vue route from another plugin
	VuePkgRoutePath(pkg string, name string, pairs ...string) string

	// Reloads the captive portal navigation items for the given client device
	ReloadPortalItems(clnt sdkconnmgr.ClientDevice)
}
