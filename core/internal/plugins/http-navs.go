package plugins

import (
	"net/http"
	sdkhttp "sdk/api/http"
)

func NewNavsApi(api *PluginApi) *HttpNavsApi {
	return &HttpNavsApi{api: api}
}

type HttpNavsApi struct {
	api          *PluginApi
	adminNavsFn  func(r *http.Request) []sdkhttp.AdminNavItem
	portalNavsFn func(r *http.Request) []sdkhttp.PortalNavItem
}

func (self *HttpNavsApi) AdminNavsFactory(fn func(r *http.Request) []sdkhttp.AdminNavItem) {
	self.adminNavsFn = fn
}

func (self *HttpNavsApi) PortalNavsFactory(fn func(r *http.Request) []sdkhttp.PortalNavItem) {
	self.portalNavsFn = fn
}

func (self *HttpNavsApi) GetAdminNavs(r *http.Request) []sdkhttp.AdminNavList {
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

		for _, p := range self.api.PluginsMgrApi.All() {
			navapi := p.Http().Navs().(*HttpNavsApi)
			adminNavs := navapi.adminNavsFn(r)
			for _, nav := range adminNavs {
				if nav.Category == category {
					navItems = append(navItems, nav)
				}
			}
		}

		navs = append(navs, sdkhttp.AdminNavList{
			Label: self.api.CoreAPI.Utl.Translate("label", string(category)),
			Items: navItems,
		})
	}

	return navs
}

func (self *HttpNavsApi) GetPortalItems(r *http.Request) []sdkhttp.PortalNavItem {
	items := []sdkhttp.PortalNavItem{}
	for _, p := range self.api.PluginsMgrApi.All() {
		navsapi := p.Http().Navs().(*HttpNavsApi)
		portalItems := navsapi.portalNavsFn(r)
		for _, item := range portalItems {
			items = append(items, item)
		}
	}
	return items
}
