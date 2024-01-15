package plugins

import "github.com/flarehotspot/core/sdk/api/themes"

func NewThemesApi(api *PluginApi) *ThemesApi {
	return &ThemesApi{api: api}
}

type ThemesApi struct {
	api         *PluginApi
	portalTheme *themes.PortalTheme
	adminTheme  *themes.AdminTheme
}

func (t *ThemesApi) AdminThemeComponent(adminTheme *themes.AdminTheme) {
	t.adminTheme = adminTheme
}

func (t *ThemesApi) GetAdminThemeComponent() (adminTheme *themes.AdminTheme, ok bool) {
	if t.adminTheme != nil {
		return t.adminTheme, true
	}
	return nil, false
}

func (t *ThemesApi) PortalThemeComponent(portalTheme *themes.PortalTheme) {
	t.portalTheme = portalTheme
}

func (t *ThemesApi) GetPortalComponent() (portalTheme *themes.PortalTheme, ok bool) {
	if t.portalTheme != nil {
		return t.portalTheme, true
	}
	return nil, false
}
