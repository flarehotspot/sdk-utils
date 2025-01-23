package themes

import (
	"net/http"
	sdkapi "sdk/api"

	"com.flarego.default-theme/resources/views/auth"
	"com.flarego.default-theme/resources/views/portal"
	"github.com/a-h/templ"
)

func SetPortalTheme(api sdkapi.IPluginApi) {

	api.Themes().NewPortalTheme(sdkapi.PortalThemeOpts{
		JsFile:  "theme.js",
		CssFile: "theme.css",
		LayoutFactory: func(w http.ResponseWriter, r *http.Request, data sdkapi.PortalLayoutData) templ.Component {
			layout := portal.PortalLayout(data)
			return layout
		},
		LoginPageFactory: func(w http.ResponseWriter, r *http.Request, data sdkapi.LoginPageData) sdkapi.ViewPage {
			csrfHtml := api.Http().Helpers().CsrfHtmlTag(r)
			page := auth.LoginPage(csrfHtml, data)
			return sdkapi.ViewPage{PageContent: page}
		},
		IndexPageFactory: func(w http.ResponseWriter, r *http.Request, data sdkapi.PortalIndexData) sdkapi.ViewPage {
			page := portal.PortalIndexPage()
			return sdkapi.ViewPage{PageContent: page}
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
