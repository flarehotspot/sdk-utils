package sdktheme

import (
	"net/http"
)

const (
	CssLibBootstrap4 CssLib = "bootstrap4"
)

type IThemesApi interface {
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
	CssLib             CssLib
}

type ThemeComponent struct {
	RouteName     string
	HandlerFunc   http.HandlerFunc
	ComponentPath string
}

type ThemeAssets struct {
	Scripts []string
	Styles  []string
}

type CssLib string
