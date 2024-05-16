package plugins

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/flarehotspot/core/internal/config"
	"github.com/flarehotspot/core/internal/web/response"
)

func NewVueResponse(vr *VueRouterApi) *VueResponse {
	data := VueResponseData{
		Response: VueResponseJson{},
	}
	return &VueResponse{vr, data}
}

type VueResponse struct {
	router *VueRouterApi
	data   VueResponseData
}

func (self *VueResponse) SetFlashMsg(msgType string, msg string) {
	self.data.Response.Flash = &VueResponseFlash{
		Type:    msgType,
		Message: msg,
	}
}

func (self *VueResponse) SendFlashMsg(w http.ResponseWriter, msgType string, msg string, status int) {
	self.SetFlashMsg(msgType, msg)
	self.Json(w, nil, status)
}

func (self *VueResponse) Json(w http.ResponseWriter, data any, status int) {
	self.data.Response.Data = data
	response.Json(w, self.data, status)
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

	self.data.Response.Redirect = &VueResponseRedirect{
		RouteName: route.VueRouteName,
		Params:    params,
		Query:     query,
	}

	response.Json(w, self.data, http.StatusOK)
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
	self.data.Response.Redirect = &VueResponseRedirect{
		RouteName: portalIndexRoute.VueRouteName,
	}
	response.Json(w, self.data, http.StatusOK)
}

func (self *VueResponse) Error(w http.ResponseWriter, err string, status int) {
	self.SendFlashMsg(w, "error", err, status)
}

type VueResponseRedirect struct {
	RouteName string            `json:"routename"`
	Params    map[string]string `json:"params"`
	Query     map[string]string `json:"query"`
}

type VueResponseFlash struct {
	Type    string `json:"type"`
	Message string `json:"msg"`
}

type VueResponseJson struct {
	Flash    *VueResponseFlash    `json:"flash"`
	Redirect *VueResponseRedirect `json:"redirect"`
	Data     interface{}          `json:"data"`
}

type VueResponseData struct {
	Response VueResponseJson `json:"$vue_response"`
}
