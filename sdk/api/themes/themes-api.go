/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkthemes

import (
	"net/http"
)

const (
	CssLibBootstrap4 CssLib = "bootstrap4"
)

type ThemesApi interface {
	NewAdminTheme(AdminTheme)
	NewPortalTheme(PortalTheme)
}

type AdminTheme struct {
	LayoutComponent    ThemeComponent
	LoginComponent     ThemeComponent
	DashboardComponent ThemeComponent
	ThemeAssets        *ThemeAssets
	CssLib             CssLib
}

type PortalTheme struct {
	LayoutComponent ThemeComponent
	IndexComponent  ThemeComponent
	ThemeAssets     *ThemeAssets
	CssLib          CssLib
}

type ThemeComponent struct {
	RouteName   string
	HandlerFunc http.HandlerFunc
	Component   string
}

type ThemeAssets struct {
	Scripts []string
	Styles  []string
}

type CssLib string
