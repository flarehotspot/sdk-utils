package themes

import (
	"sdk/api/plugin"
	"sdk/api/themes"
)

func SetPortalTheme(api sdkplugin.PluginApi) {

	api.Themes().NewPortalTheme(sdkthemes.PortalTheme{
		LayoutComponent: sdkthemes.ThemeComponent{
			Component: "portal/ThemeLayout.vue",
		},
		IndexComponent: sdkthemes.ThemeComponent{
			Component: "portal/ThemeIndex.vue",
		},
		ThemeAssets: &sdkthemes.ThemeAssets{
			Styles: []string{
				"vendor/bootstrap-4.6.1/bootstrap.min.css",
                "portal/style.css",
			},
		},
	})
}
