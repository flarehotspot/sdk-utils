package plugins

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/sdk/libs/slug"
	sdkstr "github.com/flarehotspot/core/sdk/utils/strings"
	"github.com/flarehotspot/core/utils/crypt"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/response"
	"github.com/gorilla/mux"
)

func NewVueRouteComponent(api *PluginApi, name string, path string, handler http.HandlerFunc, file string, permsReq []string, permsAny []string) *VueRouteComponent {

	compPath := filepath.Join(api.Utl.Resource("components/" + file))
	compHash, _ := crypt.SHA1Files(compPath)
	compHash = sdkstr.Sha1Hash(name, path, compPath, compHash)

	if name == "" {
		name = "empty-route-name-" + compHash
	}

	return &VueRouteComponent{
		api:                 api,
		path:                path,
		name:                name,
		file:                file,
		hash:                compHash,
		handler:             handler,
		MuxDataRouteName:    api.HttpAPI.httpRouter.MuxRouteName(sdkhttp.PluginRouteName(name + ".data")),
		VueRouteName:        api.HttpAPI.vueRouter.VueRouteName(name),
		VueRoutePath:        api.HttpAPI.vueRouter.VueRoutePath(path),
		PermissionsRequired: permsReq,
		PermissionsAnyOf:    permsAny,
	}
}

type VueRouteComponent struct {
	api                 *PluginApi       `json:"-"`
	handler             http.HandlerFunc `json:"-"`
	path                string
	name                string
	file                string               `json:"-"`
	hash                string               `json:"-"`
	MuxDataRouteName    sdkhttp.MuxRouteName `json:"mux_data_route_name"`
	HttpDataFullPath    string               `json:"http_data_full_path"`
	HttpWrapperFullPath string               `json:"http_wrapper_full_path"`
	VueRoutePath        string               `json:"vue_route_path"`
	VueRouteName        string               `json:"vue_route_name"`
	PermissionsRequired []string             `json:"permissions_required"`
	PermissionsAnyOf    []string             `json:"permissions_any_of"`
}

func (self *VueRouteComponent) HttpWrapperRouteName() string {
	return fmt.Sprintf("%s.%s.%s", self.api.Pkg(), "wrapper", self.name)
}

func (self *VueRouteComponent) HttpWrapperRoutePath() string {
	path := filepath.Join("/vue/components/wrapper", self.hash, "name", slug.Make(self.name), "file", self.file)
	if !strings.HasSuffix(path, ".vue") {
		path = path + ".vue"
	}
	return path
}

func (self *VueRouteComponent) HttpComponentFullPath() string {
	if self.file == "" {
		return self.api.CoreAPI.HttpAPI.Helpers().VueComponentPath("Empty.vue")
	}
	return self.api.HttpAPI.Helpers().VueComponentPath(self.file)
}

func (self *VueRouteComponent) HttpDataPath() string {
	path := self.api.HttpAPI.vueRouter.VuePathToMuxPath(filepath.Join("/vue/data", self.hash, self.path))
	return strings.TrimSuffix(path, "/")
}

func (self *VueRouteComponent) GetDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := NewVueResponse(self.api.HttpAPI.vueRouter)
		if self.handler == nil {
			res.Json(w, map[string]any{}, http.StatusOK)
			return
		}
		self.handler(w, r)
	}
}

func (self *VueRouteComponent) GetComponentWrapperHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wrapperFile := self.api.CoreAPI.Utl.Resource("components/Wrapper.vue")
		helpers := self.api.HttpApi().Helpers()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		response.Text(w, wrapperFile, helpers, self)
	}
}

func (self *VueRouteComponent) MountRoute(dataRouter *mux.Router, mw ...func(http.Handler) http.Handler) {
	compRouter := self.api.HttpAPI.httpRouter.pluginRouter.mux
	cacheMw := middlewares.CacheResponse(365)

	// mount wrapper handler
	wrapHandler := cacheMw(self.GetComponentWrapperHandler())
	compRouter.
		Handle(self.HttpWrapperRoutePath(), wrapHandler).
		Methods("GET").
		Name(self.HttpWrapperRouteName())

		// attache middlewares
	handler := http.Handler(self.GetDataHandler())
	if mw != nil {
		for i := len(mw) - 1; i >= 0; i-- {
			m := mw[i]
			handler = m(handler)
		}
	}

	// mount vue data path
	dataRouter.
		Handle(self.HttpDataPath(), handler).
		Methods("GET").
		Name(string(self.MuxDataRouteName))

	wrapperR := compRouter.Get(self.HttpWrapperRouteName())
	dataR := compRouter.Get(string(self.MuxDataRouteName))

	wrapperpath, _ := wrapperR.GetPathTemplate()
	datapath, _ := dataR.GetPathTemplate()

	self.HttpWrapperFullPath = wrapperpath
	self.HttpDataFullPath = datapath
}
