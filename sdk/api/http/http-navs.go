/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

import (
	"net/http"
)

type INavCategory string

// List of admin navigation menu categories.
const (
	NavCategorySystem   INavCategory = "system"
	NavCategoryPayments INavCategory = "payments"
	NavCategoryNetwork  INavCategory = "network"
	NavCategoryThemes   INavCategory = "themes"
	NavCategoryTools    INavCategory = "tools"
)

// AdminNavItem represents an admin navigation menu item.
type AdminNavItem struct {
	Category    INavCategory
	Label       string
	RouteName   string
	RouteParams map[string]string
}

type AdminNavList struct {
	Label string
	Items []AdminNavItem
}

type PortalNavItem struct {
	Label       string
	IconUrl     string
	RouteName   string
	RouteParams map[string]string
}

type NavsApi interface {
	AdminNavsFactory(func(r *http.Request) []AdminNavItem)

	PortalNavsFactory(func(r *http.Request) []PortalNavItem)

	// Returns the consolidated vue navigation list from all plugins for the admin dashboard.
	GetAdminNavs(r *http.Request) []AdminNavList

	// Returns the consolidated vue navigation list from all plugins for the portal.
	GetPortalItems(r *http.Request) []PortalNavItem
}
