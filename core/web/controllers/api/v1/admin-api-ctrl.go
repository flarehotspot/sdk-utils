package apiv1

import (
	"net/http"

	"github.com/flarehotspot/core/globals"
)

func NewAdminApiCtrl(g *globals.CoreGlobals) *AdminApiCtrl {
	return &AdminApiCtrl{g}
}

type AdminApiCtrl struct {
	g *globals.CoreGlobals
}

func (c *AdminApiCtrl) GetAdminNavs(w http.ResponseWriter, r *http.Request) {
	// allPlugins := c.g.PluginMgr.All()
	// navs := []*plugins.VueAdminNavList{}

	// systemNavs := &plugins.VueAdminNavList{
	// 	MenuHead: c.g.CoreApi.Utl.Translate(translate.Label, "system"),
	// 	Navs:     []plugins.VueAdminNav{},
	// }

	// navs = append(navs, systemNavs)

	// for _, p := range allPlugins {
	// 	vueR := p.HttpApi().VueRouter().(*plugins.VueRouterApi)
	// 	adminNavs := vueR.GetAdminNavs(r)
	// 	for _, nav := range adminNavs {
	// 		if nav.Permit(r) {
	// 			systemNavs.AddNav(nav)
	// 		}
	// 	}
	// }

	c.g.CoreApi.HttpApi().HttpResponse().Json(w, nil, http.StatusOK)
}
