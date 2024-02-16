package plugins

// import (
// 	"net/http"

// 	coreR "github.com/flarehotspot/flarehotspot/core/web/router"
// 	"github.com/flarehotspot/sdk/api/http/router"
// 	"github.com/flarehotspot/sdk/utils/translate"
// 	"github.com/gorilla/mux"
// )

// type NavAdminItem struct {
// 	r         *http.Request
// 	plugin    *PluginApi
// 	label     string
// 	translate bool
// 	route     router.PluginRouteName
// }

// func (nav *NavAdminItem) Href() string {
// 	if nav.plugin != nil {
// 		return nav.plugin.HttpApi().Router().UrlForRoute(nav.route)
// 	}
// 	url, err := coreR.UrlForRoute(router.MuxRouteName(nav.route))
// 	if err != nil {
// 		return ""
// 	}
// 	return url
// }

// func (nav *NavAdminItem) Text() string {
// 	if !nav.translate {
// 		return nav.label
// 	}
// 	if nav.plugin != nil {
// 		return nav.plugin.Translate(translate.Label, nav.label)
// 	}
// 	return translate.Core(translate.Label, nav.label)
// }

// func (nav *NavAdminItem) WhenActiveStr(str string) string {
// 	if nav.plugin != nil {
// 		muxRoute := nav.plugin.HttpAPI.Router().MuxRouteName(nav.route)
// 		if mux.CurrentRoute(nav.r).GetName() == string(muxRoute) {
// 			return str
// 		}
// 		return ""
// 	}
// 	if mux.CurrentRoute(nav.r).GetName() == string(nav.route) {
// 		return str
// 	}
// 	return ""
// }
