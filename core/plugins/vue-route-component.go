package plugins

import (
	"net/http"
	"path/filepath"

	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/web/response"
)

func NewVueRouteComponent(api *PluginApi, name string, path string, handler sdkhttp.VueHandlerFn, comp string, auth bool, permsReq []string, permsAny []string) *VueRouteComponent {

	return &VueRouteComponent{
		api:                 api,
		handler:             handler,
		component:           comp,
		MuxCompRouteName:    api.HttpAPI.httpRouter.MuxRouteName(sdkhttp.PluginRouteName(name + ".component")),
		MuxDataRouteName:    api.HttpAPI.httpRouter.MuxRouteName(sdkhttp.PluginRouteName(name + ".data")),
		HttpComponentPath:   api.HttpAPI.vueRouter.HttpComponentPath(name),
		HttpDataPath:        api.HttpAPI.vueRouter.HttpDataPath(path),
		VueRouteName:        api.HttpAPI.vueRouter.VueRouteName(name),
		VueRoutePath:        api.HttpAPI.vueRouter.VueRoutePath(path),
		RequireAuth:         auth,
		PermissionsRequired: permsReq,
		PermissionsAnyOf:    permsAny,
	}
}

type VueRouteComponent struct {
	api                   *PluginApi           `json:"-"`
	handler               sdkhttp.VueHandlerFn `json:"-"`
	component             string               `json:"-"`
	MuxCompRouteName      sdkhttp.MuxRouteName `json:"mux_component_route_name"`
	MuxDataRouteName      sdkhttp.MuxRouteName `json:"mux_data_route_name"`
	HttpComponentPath     string               `json:"http_component_path"`
	HttpComponentFullPath string               `json:"http_component_full_path"`
	HttpDataPath          string               `json:"http_data_path"`
	HttpDataFullPath      string               `json:"http_data_full_path"`
	VueRoutePath          string               `json:"vue_route_path"`
	VueRouteName          string               `json:"vue_route_name"`
	RequireAuth           bool                 `json:"require_auth"`
	PermissionsRequired   []string             `json:"permissions_required"`
	PermissionsAnyOf      []string             `json:"permissions_any_of"`
}

func (self *VueRouteComponent) GetDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := NewVueResponse(self.api.HttpAPI.vueRouter, w, r)
		if self.handler == nil {
			response.Json(w, map[string]any{}, http.StatusOK)
			return
		}
		if err := self.handler(res, r); err != nil {
			response.ErrorJson(w, err.Error())
		}
	}
}

func (self *VueRouteComponent) GetComponentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers := self.api.HttpApi().Helpers()
		comp := self.component
		compfile := self.api.Utl.Resource(filepath.Join("components", comp))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		response.Text(w, compfile, helpers, nil)
	}
}
