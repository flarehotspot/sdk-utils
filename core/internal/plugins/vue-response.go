package plugins

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/flarehotspot/core/internal/config"
	"github.com/flarehotspot/core/internal/web/response"
)

const (
	rootjson = "$vue_response"
)

func NewVueResponse(vr *VueRouterApi) *VueResponse {
	data := map[string]any{
		rootjson: map[string]any{},
	}
	return &VueResponse{vr, data}
}

type VueResponse struct {
	router *VueRouterApi
	data   map[string]any
}

func (self *VueResponse) FlashMsg(msgType string, msg string) {
	newdata := self.data[rootjson].(map[string]any)
	newdata["flash"] = map[string]string{
		"type": msgType, // "success", "error", "warning", "info
		"msg":  msg,
	}
	self.data[rootjson] = newdata
}

func (self *VueResponse) Json(w http.ResponseWriter, data any, status int) {
	newdata := self.data[rootjson].(map[string]any)
	newdata["data"] = data
	data = map[string]any{
		rootjson: newdata,
	}
	response.Json(w, data, status)
}

func (self *VueResponse) Component(w http.ResponseWriter, vuefile string, data any) {
	vuefile = self.router.api.Utl.Resource(filepath.Join("components", vuefile))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	helpers := self.router.api.HttpAPI.Helpers()
	response.Text(w, vuefile, helpers, data)
}

func (self *VueResponse) Redirect(w http.ResponseWriter, routename string, pairs ...string) {
	route, ok := self.router.FindVueRoute(routename)
	if !ok {
		response.ErrorJson(w, "Vue route \""+routename+"\" not found", 500)
		return
	}

	paramKeys := []string{}
	pathsegs := strings.Split(route.VueRoutePath.GetTemplate(), "/")
	for _, seg := range pathsegs {
		if strings.HasPrefix(seg, ":") {
			paramKeys = append(paramKeys, seg[1:])
		}
	}

	paramsMap := map[string]string{}
	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		paramsMap[key] = pairs[i+1]
	}

	params := map[string]string{}
	for _, key := range paramKeys {
		params[key] = paramsMap[key]
	}

	query := map[string]string{}
	for key, val := range paramsMap {
		if _, ok := params[key]; !ok {
			query[key] = val
		}
	}

	newdata := self.data[rootjson].(map[string]any)
	newdata["redirect"] = true
	newdata["routename"] = route.VueRouteName
	newdata["params"] = params
	newdata["query"] = query
	data := map[string]any{rootjson: newdata}

	response.Json(w, data, http.StatusOK)
}

func (self *VueResponse) RedirectToPortal(w http.ResponseWriter) {
    themecfg, err := config.ReadThemesConfig()
    if err != nil {
        self.Error(w, err.Error(), 500)
        return
    }
    themePlugin, ok := self.router.api.PluginsMgr().FindByPkg(themecfg.Portal)
    if !ok {
        self.Error(w, "Theme plugin not found", 500)
        return
    }
	portalIndexRoute := themePlugin.(*PluginApi).ThemesAPI.PortalIndexRoute
	newdata := self.data[rootjson].(map[string]any)
	newdata["redirect"] = true
	newdata["routename"] = portalIndexRoute.VueRouteName
	data := map[string]any{rootjson: newdata}
	response.Json(w, data, http.StatusOK)
}

func (self *VueResponse) Error(w http.ResponseWriter, err string, status int) {
	self.FlashMsg("error", err)
	self.Json(w, nil, status)
}
