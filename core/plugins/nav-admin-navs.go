package plugins

// import (
// 	"net/http"

// 	"github.com/flarehotspot/sdk/api/http/router"
// 	"github.com/flarehotspot/sdk/utils/translate"
// )

// func GetAdminNavs(pmgr *PluginsMgr, r *http.Request) []*router.AdminNavList {
// 	systemNavs := router.NewAdminList(translate.Core(translate.Label, "system"), []string{})
// 	networkNavs := router.NewAdminList(translate.Core(translate.Label, "network"), []string{})
// 	paymentNavs := router.NewAdminList(translate.Core(translate.Label, "payments"), []string{})
// 	adminThemeNavs := router.NewAdminList(translate.Core(translate.Label, "themes"), []string{})
// 	toolsNavs := router.NewAdminList(translate.Core(translate.Label, "tools"), []string{})
// 	navs := []*router.AdminNavList{systemNavs, networkNavs, paymentNavs, adminThemeNavs, toolsNavs}

// 	systemNavs.AddNav(&router.AdminNavJson{
// 		Category: router.CategorySystem,
// 		Label:    translate.Core(translate.Label, "dashboard"),
// 		Route:    "/system",
// 	})

// 	for _, p := range pmgr.All() {
// 		navApi := p.HttpApi().VueRouter().(*VueRouter)
// 		for _, nav := range navApi.GetAdminNavs(r) {
// 			switch nav.Category {
// 			case router.CategorySystem:
// 				systemNavs.AddNav(nav)
// 			case router.CategoryNetwork:
// 				networkNavs.AddNav(nav)
// 			case router.CategoryPayments:
// 				paymentNavs.AddNav(nav)
// 			case router.CategoryThemes:
// 				adminThemeNavs.AddNav(nav)
// 			case router.CategoryTools:
// 				toolsNavs.AddNav(nav)
// 			}
// 		}
// 	}

// 	// systemNavs.AddNav(cnav.NewAdminNavItem(navig.CategorySystem, translate.Core(translate.Label, "plugins"), names.RouteAdminPluginsIndex, []string{}))
// 	// networkNavs.AddNav(cnav.NewAdminNavItem(navig.CategoryNetwork, translate.Core(translate.Label, "bandwidth"), names.RouteAdminBandwidthIndex, []string{}))

// 	return navs
// }
