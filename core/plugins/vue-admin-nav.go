package plugins

import (
	"net/http"

	"github.com/flarehotspot/core/sdk/api/http"
	translate "github.com/flarehotspot/core/sdk/utils/translate"
	"github.com/flarehotspot/core/web/helpers"
)

func NewVueAdminNav(api *PluginApi, r *http.Request, nav sdkhttp.VueAdminNav) VueAdminNav {
	vueRouter := api.HttpAPI.vueRouter
	path := sdkhttp.VueNotFoundPath
	if route, ok := vueRouter.FindAdminRoute(nav.RouteName); ok {
		path = route.HttpDataPath
	}

	label := api.Translate(translate.Label, nav.TranslateLabel)
	return VueAdminNav{
		Category: nav.Category,
		PermitFn: nav.PermitFn,
		Label:    label,
		Path:     path,
	}
}

type VueAdminNav struct {
	Category sdkhttp.INavCategory      `json:"-"`
	PermitFn func(perms []string) bool `json:"-"`
	Label    string                    `json:"label"`
	Path     string                    `json:"path"`
}

func (nav *VueAdminNav) Permit(r *http.Request) bool {
	if nav.PermitFn == nil {
		return true
	}

	acct, err := helpers.CurrentAdmin(r)
	if err != nil {
		return false
	}

	return (nav.PermitFn)(acct.Perms)
}

type VueAdminNavList struct {
	MenuHead string        `json:"menu_head"`
	Navs     []VueAdminNav `json:"navs"`
}

func (navList *VueAdminNavList) AddNav(nav VueAdminNav) {
	navList.Navs = append(navList.Navs, nav)
}
