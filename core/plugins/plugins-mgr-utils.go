package plugins

import (
	"log"

	"github.com/flarehotspot/core/config"
)

func NewPluginsMgrUtil(pmgr *PluginsMgr) *PluginsMgrUtils {
	return &PluginsMgrUtils{
		pmgr: pmgr,
	}
}

type PluginsMgrUtils struct {
	pmgr *PluginsMgr
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
	dashboardRoute, _ := themesApi.GetDashboardVueRoute()
	children = append(children, map[string]any{
		"path":     "*",
		"redirect": dashboardRoute.VueRoutePath,
	})

	routesMap := []map[string]any{
		{
			"path":      "",
			"name":      "layout",
			"component": themesApi.AdminLayoutComponentFullPath,
			"children":  children,
			"meta": map[string]any{
				"requireAuth": true,
			},
		},
		{
			"path":      "/login",
			"name":      "login",
			"component": themesApi.AdminLoginComponentFullPath,
			"meta": map[string]any{
				"requireNoAuth": true,
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
