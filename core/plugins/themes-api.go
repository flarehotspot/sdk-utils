package plugins

import "github.com/flarehotspot/core/sdk/api/themes"

func NewThemesApi(api *PluginApi) *ThemesApi {
	return &ThemesApi{api: api}
}

type ThemesApi struct {
	api         *PluginApi
	portalTheme *themes.PortalTheme
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
