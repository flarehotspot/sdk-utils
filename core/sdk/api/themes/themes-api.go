package themes

type IThemesApi interface {
	AdminThemeComponent(AdminTheme)
	PortalThemeComponent(PortalTheme)
}

type PortalTheme struct {
	LayoutComponentPath string
	IndexComponentPath  string
	ThemeAssets         *ThemeAssets
}

type AdminTheme struct {
	LayoutComponentPath string
	IndexComponentPath  string
	LoginComponentPath  string
	ThemeAssets         *ThemeAssets
}

type ThemeAssets struct {
	Scripts []string
	Styles  []string
}
