package plugins

import (
	sdkhttp "sdk/api/http"
)

func NewThemesApi(api *PluginApi) *ThemesApi {
	return &ThemesApi{api: api}
}

type ThemesApi struct {
	api         *PluginApi
	AdminTheme  *sdkhttp.AdminTheme
	PortalTheme *sdkhttp.PortalTheme
}

func (self *ThemesApi) NewAdminTheme(theme sdkhttp.AdminTheme) {
	self.AdminTheme = &theme
}

func (self *ThemesApi) NewPortalTheme(theme sdkhttp.PortalTheme) {
	self.PortalTheme = &theme
}
