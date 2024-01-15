package themes

type IThemesApi interface {
	PortalThemeComponent(*PortalTheme)
}

type PortalTheme struct {
	ThemeComponentPath string
	IndexComponentPath string
	ThemeAssets        *ThemeAssets
}

type ThemeAssets struct {
	Scripts *[]string
	Styles  *[]string
}
