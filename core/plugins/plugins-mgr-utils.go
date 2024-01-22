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
		adminRoutes := vueR.GetAdminRoutes()
		routes = append(routes, adminRoutes...)
	}

	children := []map[string]any{}
	for _, r := range routes {
		children = append(children, map[string]any{
			"path":      r.VueRoutePath,
			"name":      r.VueRouteName,
			"component": r.HttpComponentFullPath,
			"meta": map[string]any{
				"data_path": r.HttpDataFullPath,
			},
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
	// dashboardRoute, _ := themesApi.GetDashboardVueRoute()
	children = append(children, map[string]any{
		"path":     "*",
		"redirect": themesApi.AdminDashVuePath,
	})

	routesMap := []map[string]any{
		{
			"path":      "",
			"name":      "layout",
			"component": themesApi.AdminLayoutComponentFullPath,
			"children":  children,
			"meta": map[string]any{
				"requireAuth": true,
				"data_path":   themesApi.AdminLayoutDataFullPath,
			},
		},
		{
			"path":      "/login",
			"name":      "login",
			"component": themesApi.AdminLoginComponentFullPath,
			"meta": map[string]any{
				"requireNoAuth": true,
				"data_path":     themesApi.AdminLoginDataFullPath,
			},
		},
	}

	return routesMap
}

func (utils *PluginsMgrUtils) GetPortalRoutes() []*VueRouteComponent {
	routes := []*VueRouteComponent{}
	for _, p := range utils.pmgr.All() {
		vueR := p.HttpApi().VueRouter().(*VueRouterApi)
		portalRoutes := vueR.GetPortalRoutes()
		routes = append(routes, portalRoutes...)
	}
	return routes
}

func (utils *PluginsMgrUtils) GetAdminNavs(r *http.Request) []sdkhttp.AdminNavCategory {
	navs := []sdkhttp.AdminNavCategory{}
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

		navs = append(navs, sdkhttp.AdminNavCategory{
			Category: utils.coreApi.Utl.Translate("label", string(category)),
			Items:    navItems,
		})
	}

	return navs
}
