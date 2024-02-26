/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

type INavCategory string

// List of admin navigation menu categories.
const (
	NavCategorySystem   INavCategory = "system"
	NavCategoryPayments INavCategory = "payments"
	NavCategoryNetwork  INavCategory = "network"
	NavCategoryThemes   INavCategory = "themes"
	NavCategoryTools    INavCategory = "tools"
)

// VueAdminNav represents an admin navigation menu item.
type VueAdminNav struct {
	Category    INavCategory
	Label       string
	RouteName   string
	RouteParams map[string]string
}
