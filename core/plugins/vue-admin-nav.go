package plugins

import (
	"log"
	"net/http"

	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/web/helpers"
)

func NewVueAdminNav(api *PluginApi, r *http.Request, nav sdkhttp.VueAdminNav) (sdkhttp.AdminNavItem, bool) {
	var adminNav sdkhttp.AdminNavItem

	if nav.PermitFn != nil {
		acct, err := helpers.CurrentAdmin(r)
		if err != nil {
			log.Println("Warning: helpers.CurrentAdmin() failed: ", err)
			return adminNav, false
		}

		if !nav.PermitFn(acct.Permissions()) {
			return adminNav, false
		}
	}

	var routename, routepath string
	routepath = sdkhttp.VueNotFoundPath

	if vueRoute, ok := api.HttpAPI.vueRouter.FindVueRoute(nav.RouteName); ok {
		// nav.RouteParams is a map[string]string
		pairs := []string{}
		for k, v := range nav.RouteParams {
			pairs = append(pairs, k, v)
		}
		routename = vueRoute.VueRouteName
		routepath = vueRoute.VueRoutePath.URL(pairs...)
	}

	label := api.Translate("label", nav.Label)

	return sdkhttp.AdminNavItem{
		Category:     nav.Category,
		Label:        label,
		VueRouteName: routename,
		VueRoutePath: routepath,
	}, true
}
