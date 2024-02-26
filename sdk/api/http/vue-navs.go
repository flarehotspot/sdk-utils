/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

type AdminNavList struct {
	Label string         `json:"label"`
	Items []AdminNavItem `json:"items"`
}

type AdminNavItem struct {
	Category       INavCategory      `json:"category"`
	Label          string            `json:"label"`
	VueRouteName   string            `json:"route_name"`
	VueRoutePath   string            `json:"route_path"`
	VueRouteParams map[string]string `json:"route_params"`
}

type PortalItem struct {
	IconUri        string            `json:"icon_uri"`
	Label          string            `json:"label"`
	VueRouteName   string            `json:"route_name"`
	VueRoutePath   string            `json:"route_path"`
	VueRouteParams map[string]string `json:"route_params"`
}
