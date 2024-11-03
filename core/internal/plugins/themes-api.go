package plugins

import (
	"core/internal/utils/pkg"
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

func (self *ThemesApi) GetAdminAssets() (jsSrc string, cssHref string) {
	manifest := pkg.GetAssetManifest(self.api.dir)

	if self.AdminTheme != nil {
		scriptFile, ok := manifest.AdminAssets.Scripts[self.AdminTheme.JsFile]
		if ok {
			jsSrc = self.api.HttpAPI.Helpers().AssetPath(scriptFile)
		}

		cssFile, ok := manifest.AdminAssets.Styles[self.AdminTheme.CssFile]
		if ok {
			cssHref = self.api.HttpAPI.Helpers().AssetPath(cssFile)
		}
	}

	return
}

func (self *ThemesApi) GetPortalAssets() (jsSrc string, cssHref string) {
	manifest := self.api.AssetsManifest

	if self.PortalTheme != nil {
		scriptFile, ok := manifest.PortalAssets.Scripts[self.PortalTheme.JsFile]
		if ok {
			jsSrc = self.api.HttpAPI.Helpers().AssetPath(scriptFile)
		}

		cssFile, ok := manifest.PortalAssets.Styles[self.PortalTheme.CssFile]
		if ok {
			cssHref = self.api.HttpAPI.Helpers().AssetPath(cssFile)
		}
	}

	return
}
