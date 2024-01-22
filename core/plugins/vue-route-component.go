package plugins

import (
	"net/http"
	"path/filepath"

	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/sdk/libs/slug"
	sdkstr "github.com/flarehotspot/core/sdk/utils/strings"
	"github.com/flarehotspot/core/web/response"
	"github.com/gorilla/mux"
)

func NewVueRouteComponent(api *PluginApi, name string, path string, handler sdkhttp.VueHandlerFn, comp string, permsReq []string, permsAny []string) *VueRouteComponent {
	if name == "" {
		name = sdkstr.Rand(8) + "-" + slug.Make(comp)
	}

	return &VueRouteComponent{
		api:                  api,
		handler:              handler,
		component:            comp,
		MuxCompRouteName:     api.HttpAPI.httpRouter.MuxRouteName(sdkhttp.PluginRouteName(name + ".component")),
		MuxDataRouteName:     api.HttpAPI.httpRouter.MuxRouteName(sdkhttp.PluginRouteName(name + ".data")),
		HttpWrapperRouteName: api.HttpAPI.vueRouter.HttpWrapperRouteName(name),
		HttpWrapperPath:      api.HttpAPI.vueRouter.HttpWrapperRoutePath(name),
		HttpComponentPath:    api.HttpAPI.vueRouter.HttpComponentPath(name),
		HttpDataPath:         api.HttpAPI.vueRouter.HttpDataPath(path),
		VueRouteName:         api.HttpAPI.vueRouter.VueRouteName(name),
		VueRoutePath:         api.HttpAPI.vueRouter.VueRoutePath(path),
		PermissionsRequired:  permsReq,
		PermissionsAnyOf:     permsAny,
	}
}

type VueRouteComponent struct {
	api                   *PluginApi           `json:"-"`
	handler               sdkhttp.VueHandlerFn `json:"-"`
	component             string               `json:"-"`
	MuxCompRouteName      sdkhttp.MuxRouteName `json:"mux_component_route_name"`
	MuxDataRouteName      sdkhttp.MuxRouteName `json:"mux_data_route_name"`
	HttpWrapperRouteName  string               `json:"http_wrapper_route_name"`
	HttpComponentPath     string               `json:"http_component_path"`
	HttpComponentFullPath string               `json:"http_component_full_path"`
	HttpDataPath          string               `json:"http_data_path"`
	HttpDataFullPath      string               `json:"http_data_full_path"`
	HttpWrapperPath       string               `json:"http_wrapper_path"`
	HttpWrapperFullPath   string               `json:"http_wrapper_full_path"`
	VueRoutePath          string               `json:"vue_route_path"`
	VueRouteName          string               `json:"vue_route_name"`
	PermissionsRequired   []string             `json:"permissions_required"`
	PermissionsAnyOf      []string             `json:"permissions_any_of"`
}

func (self *VueRouteComponent) GetDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := NewVueResponse(self.api.HttpAPI.vueRouter, w, r)
		if self.handler == nil {
			res.JsonData(map[string]any{})
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

func (self *VueRouteComponent) GetComponentWrapperHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wrapperFile := self.api.coreApi.Utl.Resource("views/vue/component-wrapper.html")
		helpers := self.api.HttpApi().Helpers()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		response.Text(w, wrapperFile, helpers, self)
	}
}

func (self *VueRouteComponent) MountRoute(dataRouter *mux.Router, middlewares ...func(http.Handler) http.Handler) {
	rand := sdkstr.Rand(8)
	compRouter := self.api.HttpAPI.httpRouter.pluginRouter.mux.PathPrefix("/vue/components/" + rand).Subrouter()
	dataRouter = dataRouter.PathPrefix("/vue/data/" + rand).Subrouter()

	// mount vue component path
	compRouter.
		HandleFunc(self.HttpComponentPath, self.GetComponentHandler()).
		Methods("GET").
		Name(string(self.MuxCompRouteName))

	// mount vue data path

	if middlewares != nil {
		for _, m := range middlewares {
			dataRouter.Use(m)
		}
	}

	dataRouter.
		HandleFunc(self.HttpDataPath, self.GetDataHandler()).
		Methods("GET").
		Name(string(self.MuxDataRouteName))

	// mount wrapper handler
	wrapHandler := self.GetComponentWrapperHandler()
	compRouter.
		Handle(self.HttpWrapperPath, wrapHandler).
		Methods("GET").
		Name(self.HttpWrapperRouteName)

	wrapperR := compRouter.Get(self.HttpWrapperRouteName)
	compR := compRouter.Get(string(self.MuxCompRouteName))
	dataR := compRouter.Get(string(self.MuxDataRouteName))

	wrapperpath, _ := wrapperR.GetPathTemplate()
	comppath, _ := compR.GetPathTemplate()
	datapath, _ := dataR.GetPathTemplate()

	self.HttpWrapperFullPath = wrapperpath
	self.HttpComponentFullPath = comppath
	self.HttpDataFullPath = datapath
}
