package themes

import (
	sdkhttp "sdk/api/http"
	sdkplugin "sdk/api/plugin"

	"com.flarego.default-theme/resources/views"
	"github.com/a-h/templ"
)

func SetPortalTheme(api sdkplugin.PluginApi) {

	api.Themes().NewPortalTheme(sdkhttp.PortalTheme{
		GlobalScripts:     []string{"test.js"},
		GlobalStylesheets: []string{"test.css"},
		LayoutFactory: func(data sdkhttp.PortalLayoutData) templ.Component {
			layout := views.PortalLayout(data)
			return layout
		},
	})

	// api.Themes().NewPortalTheme(sdkthemes.PortalTheme{
	// 	LayoutComponent: sdkthemes.ThemeComponent{
	// 		Component: "portal/ThemeLayout.vue",
	// 	},
	// 	IndexComponent: sdkthemes.ThemeComponent{
	// 		Component: "portal/ThemeIndex.vue",
	// 	},
	// 	ThemeAssets: &sdkthemes.ThemeAssets{
	// 		Styles: []string{
	// 			"vendor/bootstrap-4.6.1/bootstrap.min.css",
	//                "portal/style.css",
	// 		},
	// 	},
	// })
}
