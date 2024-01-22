package sdktheme

type IThemesApi interface {
	NewAdminTheme(AdminTheme)
	NewPortalTheme(PortalTheme)
}

type AdminTheme struct {
	LayoutComponent    ThemeComponent
	LoginComponentPath ThemeComponent
	DashboardRoute     string
	ThemeAssets        *ThemeAssets
}

type PortalTheme struct {
	LayoutComponent ThemeComponent
	IndexComponent  ThemeComponent
	ThemeAssets     *ThemeAssets
}

type ThemeComponent struct {
	Data          any
	ComponentPath string
}

type ThemeAssets struct {
	Scripts []string
	Styles  []string
}
