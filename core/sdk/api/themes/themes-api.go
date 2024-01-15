package themes

type IThemesApi interface {
	AdminThemeComponent(*AdminTheme)
	PortalThemeComponent(*PortalTheme)
}

type PortalTheme struct {
	ThemeComponentPath string
	IndexComponentPath string
	ThemeAssets        *ThemeAssets
}

type AdminTheme struct {
	ThemeComponentPath string
	IndexComponentPath string
	ThemeAssets        *ThemeAssets
}

type ThemeAssets struct {
	Scripts *[]string
	Styles  *[]string
}
