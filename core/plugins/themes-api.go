package plugins

import "github.com/flarehotspot/core/sdk/api/themes"

func NewThemesApi(api *PluginApi) *ThemesApi {
	return &ThemesApi{api: api}
}

type ThemesApi struct {
	api         *PluginApi
	adminTheme  themes.AdminTheme
	portalTheme themes.PortalTheme
}

func (t *ThemesApi) AdminThemeComponent(adminTheme themes.AdminTheme) {
	t.adminTheme = adminTheme
}

func (t *ThemesApi) GetAdminLayoutComponents() (adminTheme themes.AdminTheme, ok bool) {
	if t.adminTheme.IndexComponentPath != "" {
		adminTheme = themes.AdminTheme{
			LayoutComponentPath: t.api.HttpApi().AssetPath(t.adminTheme.LayoutComponentPath),
			IndexComponentPath:  t.api.HttpApi().AssetPath(t.adminTheme.IndexComponentPath),
			LoginComponentPath:  t.api.HttpApi().AssetPath(t.adminTheme.LoginComponentPath),
			ThemeAssets:         adminTheme.ThemeAssets,
		}
		return adminTheme, true
	}
	return themes.AdminTheme{}, false
}

func (t *ThemesApi) PortalThemeComponent(portalTheme themes.PortalTheme) {
	t.portalTheme = portalTheme
}

func (t *ThemesApi) GetPortalThemeComponents() (portalTheme themes.PortalTheme, ok bool) {
	if t.portalTheme.IndexComponentPath != "" {
		portalTheme = themes.PortalTheme{
			LayoutComponentPath: t.api.HttpApi().AssetPath(t.portalTheme.LayoutComponentPath),
			IndexComponentPath:  t.api.HttpApi().AssetPath(t.portalTheme.IndexComponentPath),
			ThemeAssets:         portalTheme.ThemeAssets,
		}
		return portalTheme, true
	}
	return themes.PortalTheme{}, false
}
