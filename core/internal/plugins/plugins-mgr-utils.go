package plugins

import (
	"sdk/api/accounts"
	"sdk/api/connmgr"
	"sdk/api/http"
)

func NewPluginsMgrUtil(pmgr *PluginsMgr, coreApi *PluginApi) *PluginsMgrUtils {
	return &PluginsMgrUtils{
		pmgr:    pmgr,
		coreApi: coreApi,
	}
}

type PluginsMgrUtils struct {
	pmgr    *PluginsMgr
	coreApi *PluginApi
}

func (self *PluginsMgrUtils) GetAdminNavs(acct sdkacct.Account) []sdkhttp.AdminNavList {
	navs := []sdkhttp.AdminNavList{}
	categories := []sdkhttp.INavCategory{
		sdkhttp.NavCategorySystem,
		sdkhttp.NavCategoryPayments,
		sdkhttp.NavCategoryNetwork,
		sdkhttp.NavCategoryThemes,
		sdkhttp.NavCategoryTools,
	}

	for _, category := range categories {
		navItems := []sdkhttp.AdminNavItem{}

		for _, p := range self.pmgr.All() {
			vueR := p.Http().VueRouter().(*VueRouterApi)
			adminNavs := vueR.GetAdminNavs(acct)
			for _, nav := range adminNavs {
				if nav.Category == category {
					navItems = append(navItems, nav)
				}
			}
		}

		navs = append(navs, sdkhttp.AdminNavList{
			Label: self.coreApi.Utl.Translate("label", string(category)),
			Items: navItems,
		})
	}

	return navs
}

func (self *PluginsMgrUtils) GetPortalItems(clnt sdkconnmgr.ClientDevice) []sdkhttp.PortalItem {
	items := []sdkhttp.PortalItem{}
	for _, p := range self.pmgr.All() {
		vueR := p.Http().VueRouter().(*VueRouterApi)
		portalItems := vueR.GetPortalItems(clnt)
		for _, item := range portalItems {
			items = append(items, item)
		}
	}
	return items
}
