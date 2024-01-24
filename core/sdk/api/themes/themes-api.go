package sdktheme

import (
	"net/http"
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
}

type PortalTheme struct {
	LayoutComponent ThemeComponent
	IndexComponent  ThemeComponent
	ThemeAssets     *ThemeAssets
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
