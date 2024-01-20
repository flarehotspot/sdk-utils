package apiv1

import (
	"net/http"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/plugins"
)

func NewPortalApiCtrl(g *globals.CoreGlobals) PortalApiCtrl {
	return PortalApiCtrl{g}
}

type PortalApiCtrl struct {
	g *globals.CoreGlobals
}

func (c *PortalApiCtrl) PortalNavs(w http.ResponseWriter, r *http.Request) {
	portalItems := []plugins.VuePortalItem{}
	allPlugins := c.g.PluginMgr.All()

	for _, p := range allPlugins {
		vueRouter := p.HttpApi().VueRouter().(*plugins.VueRouterApi)
		portalItems = append(portalItems, vueRouter.GetPortalItems(r)...)
	}

	c.g.CoreApi.HttpApi().HttpResponse().Json(w, portalItems, http.StatusOK)
}
