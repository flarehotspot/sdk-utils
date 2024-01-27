package plugins

import (
	"log"
	"net/http"

	"github.com/flarehotspot/core/config"
	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
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

func (util *PluginsMgrUtils) GetAdminRoutes() []map[string]any {
	routes := []*VueRouteComponent{}
	for _, p := range util.pmgr.All() {
		vueR := p.HttpApi().VueRouter().(*VueRouterApi)
		adminRoutes := vueR.adminRoutes
		routes = append(routes, adminRoutes...)
	}

	children := []map[string]any{}
	for _, r := range routes {
		children = append(children, map[string]any{
			"path":      r.VueRoutePath,
			"name":      r.VueRouteName,
			"component": r.HttpWrapperFullPath,
		})
	}

	themecfg, err := config.ReadThemesConfig()
	if err != nil {
		log.Println("Error reading themes config: ", err)
	}

	themesPlugin, ok := util.pmgr.FindByPkg(themecfg.Admin)
	if !ok {
		log.Println("Invalid admin theme: ", themecfg.Admin)
	}

	themesApi := themesPlugin.ThemesApi().(*ThemesApi)
	children = append(children, map[string]any{
		"path":     "*",
		"redirect": themesApi.AdminDashboardRoute.VueRoutePath,
	})

	routesMap := []map[string]any{
		{
			"path":      "/",
			"name":      themesApi.AdminLayoutRoute.VueRouteName,
			"component": themesApi.AdminLayoutRoute.HttpWrapperFullPath,
			"children":  children,
			"meta": map[string]any{
				"requireAuth": true,
			},
		},
		{
			"path":      themesApi.AdminLoginRoute.VueRoutePath,
			"name":      themesApi.AdminLoginRoute.VueRouteName,
			"component": themesApi.AdminLoginRoute.HttpWrapperFullPath,
			"meta": map[string]any{
				"requireNoAuth": true,
			},
		},
	}

	return routesMap
}

func (util *PluginsMgrUtils) GetPortalRoutes() []map[string]any {
	routes := []*VueRouteComponent{}
	for _, p := range util.pmgr.All() {
		vueR := p.HttpApi().VueRouter().(*VueRouterApi)
		portalRoutes := vueR.portalRoutes
		routes = append(routes, portalRoutes...)
	}

	children := []map[string]any{}
	for _, r := range routes {
		children = append(children, map[string]any{
			"path":      r.VueRoutePath,
			"name":      r.VueRouteName,
			"component": r.HttpWrapperFullPath,
		})
	}

	themecfg, err := config.ReadThemesConfig()
	if err != nil {
		log.Println("Error reading themes config: ", err)
	}

	themesPlugin, ok := util.pmgr.FindByPkg(themecfg.Portal)
	if !ok {
		log.Println("Invalid portal theme: ", themecfg.Portal)
	}

	themesApi := themesPlugin.ThemesApi().(*ThemesApi)
	children = append(children, map[string]any{
		"path":     "*",
		"redirect": themesApi.PortalIndexRoute.VueRoutePath,
	})

	routesMap := []map[string]any{
		{
			"path":      "/",
			"name":      themesApi.PortalLayoutRoute.VueRouteName,
			"component": themesApi.PortalLayoutRoute.HttpWrapperFullPath,
			"children":  children,
		},
	}

	return routesMap
}

func (utils *PluginsMgrUtils) GetAdminNavs(r *http.Request) []sdkhttp.AdminNavList {
	navs := []sdkhttp.AdminNavList{}
	categories := []sdkhttp.INavCategory{
		sdkhttp.NavCategorySystem,
		// sdkhttp.NavCategoryNetwork,
	}

	for _, category := range categories {
		navItems := []sdkhttp.AdminNavItem{}

		for _, p := range utils.pmgr.All() {
			vueR := p.HttpApi().VueRouter().(*VueRouterApi)
			adminNavs := vueR.GetAdminNavs(r)
			for _, nav := range adminNavs {
				if nav.Category == category {
					navItems = append(navItems, nav)
				}
			}
		}

		navs = append(navs, sdkhttp.AdminNavList{
			Label: utils.coreApi.Utl.Translate("label", string(category)),
			Items: navItems,
		})
	}

	return navs
}

func (utils *PluginsMgrUtils) GetPortalItems(r *http.Request) []sdkhttp.PortalItem {
	items := []sdkhttp.PortalItem{}
	for _, p := range utils.pmgr.All() {
		vueR := p.HttpApi().VueRouter().(*VueRouterApi)
		portalItems := vueR.GetPortalItems(r)
		for _, item := range portalItems {
			items = append(items, item)
		}
	}
	return items
}
