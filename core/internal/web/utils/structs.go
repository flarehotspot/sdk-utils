package utils

type TmpThemeList struct {
	ThemeAdmin []struct {
		ThemeName string `json:"theme_name"`
		ThemePkg  string `json:"theme_pkg"`
	} `json:"theme_admin"`

	ThemePortal []struct {
		ThemeName string `json:"theme_name"`
		ThemePkg  string `json:"theme_pkg"`
	} `json:"theme_portal"`
}
type SavedThemeData struct {
	ThemeAdmin  string `json:"admin"`
	ThemePortal string `json:"portal"`
}
