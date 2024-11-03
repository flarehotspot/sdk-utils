/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

import "github.com/a-h/templ"

const (
	CssLibBootstrap4 CssLib = "bootstrap4"
)

type ThemesApi interface {
	NewAdminTheme(AdminTheme)
	NewPortalTheme(PortalTheme)
}

type CssLib string

type FlashMsg struct {
	Type    string
	Message string
}

type PageAssets struct {
	ThemeJsSrc   string
	ThemeCssHref string
	PageJsSrc    string
	PageCssHref  string
}

type LayoutData struct {
	Assets      PageAssets
	FlashMsg    FlashMsg
	PageContent templ.Component
}

type AdminLayoutData struct {
	Layout LayoutData
	Navs   AdminNavList
}

type PortalLayoutData struct {
	Layout LayoutData
	Navs   []PortalNavItem
}

type AdminTheme struct {
	IndexRoute    string
	JsFile        string
	CssFile       string
	CssLib        CssLib
	LayoutFactory func(data AdminLayoutData) templ.Component
}

type PortalTheme struct {
	IndexRoute    string
	JsFile        string
	CssFile       string
	CssLib        CssLib
	LayoutFactory func(data PortalLayoutData) templ.Component
}
