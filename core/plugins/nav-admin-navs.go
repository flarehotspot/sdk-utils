package plugins

import (
	"net/http"

	cnav "github.com/flarehotspot/core/web/navigation"
	"github.com/flarehotspot/core/web/routes/names"
	navig "github.com/flarehotspot/core/sdk/api/http/navigation"
	"github.com/flarehotspot/core/sdk/utils/translate"
)

func GetAdminNavs(pmgr *PluginsMgr, r *http.Request) []navig.IAdminNavList {
	systemNavs := cnav.NewAdminListItem(translate.Core(translate.Label, "system"), []string{})
	networkNavs := cnav.NewAdminListItem(translate.Core(translate.Label, "network"), []string{})
	paymentNavs := cnav.NewAdminListItem(translate.Core(translate.Label, "payments"), []string{})
	adminThemeNavs := cnav.NewAdminListItem(translate.Core(translate.Label, "themes"), []string{})
	toolsNavs := cnav.NewAdminListItem(translate.Core(translate.Label, "tools"), []string{})
	navs := []navig.IAdminNavList{systemNavs, networkNavs, paymentNavs, adminThemeNavs, toolsNavs}

	systemNavs.AddNav(cnav.NewAdminNavItem(navig.CategorySystem, translate.Core(translate.Label, "dashboard"), names.RouteAdminDashboardIndex, []string{}))

	for _, p := range pmgr.All() {
		navApi := p.NavApi().(*NavApi)
		for _, nav := range navApi.GetAdminNavs(r) {
			switch nav.Category() {
			case navig.CategorySystem:
				systemNavs.AddNav(nav)
			case navig.CategoryNetwork:
				networkNavs.AddNav(nav)
			case navig.CategoryPayments:
				paymentNavs.AddNav(nav)
			case navig.CategoryThemes:
				adminThemeNavs.AddNav(nav)
			case navig.CategoryTools:
				toolsNavs.AddNav(nav)
			}
		}
	}

	systemNavs.AddNav(cnav.NewAdminNavItem(navig.CategorySystem, translate.Core(translate.Label, "plugins"), names.RouteAdminPluginsIndex, []string{}))
	networkNavs.AddNav(cnav.NewAdminNavItem(navig.CategoryNetwork, translate.Core(translate.Label, "bandwidth"), names.RouteAdminBandwidthIndex, []string{}))

	return navs
}
