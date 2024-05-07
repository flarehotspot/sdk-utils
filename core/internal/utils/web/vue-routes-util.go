package webutil

import (
	"errors"
	"log"

	"github.com/flarehotspot/core/internal/config"
	"github.com/flarehotspot/core/internal/plugins"
)

type ChildRoutes struct {
	Path      string `json:"path"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Meta      struct {
		AuthRequired bool `json:"auth_required"`
	} `json:"meta"`
}

type ThemeComponent struct {
	Path      string `json:"path"`
	Name      string `json:"name"`
	Component string `json:"component"`
}

type portalRoutesData struct {
	ThemeComponent `json:"theme_component"`
	ChildRoutes    []ChildRoutes `json:"child_routes"`
}

func GetPortalRoutesData(g *plugins.CoreGlobals) (portalRoutesData, error) {
	var data portalRoutesData

	routes := []*plugins.VueRouteComponent{}
	for _, p := range g.PluginMgr.All() {
		vueR := p.Http().VueRouter().(*plugins.VueRouterApi)
		portalRoutes := vueR.PortalRoutes
		routes = append(routes, portalRoutes...)
	}

	for _, r := range routes {
		data.ChildRoutes = append(data.ChildRoutes, ChildRoutes{
			Path:      string(r.VueRoutePath),
			Name:      r.VueRouteName,
			Component: r.HttpComponentPath,
		})
	}

	themecfg, err := config.ReadThemesConfig()
	if err != nil {
		log.Println("Error reading themes config: ", err)
        return data, err
	}

	themePkg, ok := g.PluginMgr.FindByPkg(themecfg.Portal)
	if !ok {
		log.Println("Invalid portal theme: ", themecfg.Portal)
		return data, errors.New("Invalid portal theme")
	}

	themeApi := themePkg.Themes().(*plugins.ThemesApi)

	return data, nil
}
