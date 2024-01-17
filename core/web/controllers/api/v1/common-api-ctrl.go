package apiv1

import (
	"net/http"

	"github.com/flarehotspot/core/globals"
	"github.com/gorilla/mux"
)

func NewCommonApiCtrl(g *globals.CoreGlobals) CommonApiCtrl {
	return CommonApiCtrl{g}
}

type CommonApiCtrl struct {
	g *globals.CoreGlobals
}

// Render api-v1.js
func (c *CommonApiCtrl) ApiJs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkg := vars["pkg"]
	pluginApi := c.g.PluginMgr.FindByPkg(pkg)
	if pluginApi == nil {
		http.Error(w, "Plugin not found: "+pkg, 404)
		return
	}

	vdata := map[string]any{
		"CoreApi": c.g.CoreApi,
		"Plugin":  pluginApi,
	}

	c.g.CoreApi.HttpApi().Respond().Text(w, r, "views/js/api-v1.tpl.js", vdata)
}
