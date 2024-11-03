package themes

import (
	// "net/http"

	"net/http"
	sdkhttp "sdk/api/http"
	plugin "sdk/api/plugin"

	"com.flarego.default-theme/resources/views"
	"com.flarego.default-theme/resources/views/admin"
	"com.flarego.default-theme/resources/views/auth"
	"github.com/a-h/templ"
)

func SetAdminTheme(api plugin.PluginApi) {

	api.Themes().NewAdminTheme(sdkhttp.AdminThemeOpts{
		LayoutFactory: func(w http.ResponseWriter, r *http.Request, data sdkhttp.AdminLayoutData) templ.Component {
			layout := views.AdminLayout(data)
			return layout
		},
		LoginPageFactory: func(w http.ResponseWriter, r *http.Request, data sdkhttp.LoginPageData) sdkhttp.ViewPage {
			page := auth.LoginPage()
			return sdkhttp.ViewPage{PageContent: page}
		},
		IndexPageFactory: func(w http.ResponseWriter, r *http.Request) sdkhttp.ViewPage {
			page := admin.AdminIndexPage()
			return sdkhttp.ViewPage{PageContent: page}
		},
	})

	// api.Http().VueRouter().AdminNavsFunc(func(acct sdkacct.Account) []sdkhttp.VueAdminNav {
	// 	return []sdkhttp.VueAdminNav{
	// 		{
	// 			Category:  sdkhttp.NavCategorySystem,
	// 			Label:     "Dashboard",
	// 			RouteName: "dashboard",
	// 		},
	// 	}
	// })

	// api.Themes().NewAdminTheme(themes.AdminTheme{
	// 	CssLib: themes.CssLibBootstrap4,
	// 	DashboardComponent: themes.ThemeComponent{
	// 		RouteName: "dashboard",
	// 		Component: "admin/Dashboard.vue",
	// 	},
	// 	LayoutComponent: themes.ThemeComponent{
	// 		Component: "admin/ThemeLayout.vue",
	// 	},
	// 	LoginComponent: themes.ThemeComponent{
	// 		RouteName: "login",
	// 		Component: "admin/ThemeLogin.vue",
	// 	},
	// 	ThemeAssets: &themes.ThemeAssets{
	// 		Scripts: []string{
	// 			"vendor/polyfills/intersection-observer.js",
	// 			"vendor/polyfills/intersection-observer-enable-polling.js",
	// 			"vendor/bootstrap-vue/bootstrap-vue-2.23.1.umd.min.js",
	// 			"vendor/bootstrap-vue/bootstrap-vue-icons-2.23.1.umd.min.js",
	// 		},
	// 		Styles: []string{
	// 			"vendor/bootstrap-4.6.1/bootstrap.min.css",
	// 		},
	// 	},
	// })
}
