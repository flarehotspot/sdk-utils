package plugins

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkstr "github.com/flarehotspot/sdk/utils/strings"
	"github.com/flarehotspot/flarehotspot/core/utils/crypt"
	"github.com/flarehotspot/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/flarehotspot/core/web/response"
	"github.com/gorilla/mux"
)

func NewVueRouteComponent(api *PluginApi, name string, p string, handler http.HandlerFunc, file string, permsReq []string, permsAny []string) *VueRouteComponent {

	compPath := filepath.Join(api.Utl.Resource("components/" + file))
	compHash, _ := crypt.SHA1Files(compPath)
	compHash = sdkstr.Sha1Hash(name, p, compPath, compHash)

	if name == "" {
		name = "empty-route-name-" + compHash
	}

	return &VueRouteComponent{
		api:                 api,
		path:                p,
		name:                name,
		file:                file,
		hash:                compHash,
		handler:             handler,
		MuxDataRouteName:    api.HttpAPI.httpRouter.MuxRouteName(sdkhttp.PluginRouteName(name + ".data")),
		VueRouteName:        api.HttpAPI.vueRouter.VueRouteName(name),
		VueRoutePath:        api.HttpAPI.vueRouter.VueRoutePath(p),
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
	VueRoutePath        VueRoutePath         `json:"vue_route_path"`
	VueRouteName        string               `json:"vue_route_name"`
	PermissionsRequired []string             `json:"permissions_required"`
	PermissionsAnyOf    []string             `json:"permissions_any_of"`
}

func (self *VueRouteComponent) HttpWrapperRouteName() string {
	return fmt.Sprintf("%s.%s.%s", self.api.Pkg(), "wrapper", self.name)
}

func (self *VueRouteComponent) HttpWrapperRoutePath() string {
	p := path.Join("/vue/components/wrapper", self.hash, "name", self.name, "file", self.file)
	if !strings.HasSuffix(p, ".vue") {
		p = p + ".vue"
	}
	return p
}

func (self *VueRouteComponent) HttpComponentFullPath() string {
	if self.file == "" {
		return self.api.CoreAPI.HttpAPI.Helpers().VueComponentPath("Empty.vue")
	}
	return self.api.HttpAPI.Helpers().VueComponentPath(self.file)
}

func (self *VueRouteComponent) HttpDataPath() string {
	p := self.api.HttpAPI.vueRouter.VuePathToMuxPath(path.Join("/vue/data", self.hash, self.path))
	return strings.TrimSuffix(p, "/")
}

func (self *VueRouteComponent) GetDataHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := NewVueResponse(self.api.HttpAPI.vueRouter)
		if self.handler == nil {
			res.Json(w, map[string]any{}, http.StatusOK)
			return
		}
		self.handler(w, r)
	})
}

func (self *VueRouteComponent) GetComponentWrapperHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wrapperFile := self.api.CoreAPI.Utl.Resource("components/Wrapper.vue")
		helpers := self.api.Http().Helpers()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		response.Text(w, wrapperFile, helpers, self)
	}
}

func (self *VueRouteComponent) MountRoute(dataRouter *mux.Router, mws ...func(http.Handler) http.Handler) {
	compRouter := self.api.HttpAPI.httpRouter.pluginRouter.mux
	cacheMw := middlewares.CacheResponse(365)

	// mount wrapper handler
	wrapHandler := cacheMw(self.GetComponentWrapperHandler())
	wrapperR := compRouter.
		Handle(self.HttpWrapperRoutePath(), wrapHandler).
		Methods("GET").
		Name(self.HttpWrapperRouteName())

		// attache middlewares
	finalHandler := self.GetDataHandler()
	for i := len(mws) - 1; i >= 0; i-- {
		finalHandler = mws[i](finalHandler)
	}

	// mount vue data path
	dataR := dataRouter.
		Handle(self.HttpDataPath(), finalHandler).
		Methods("GET").
		Name(string(self.MuxDataRouteName))

	// wrapperR := compRouter.Get(self.HttpWrapperRouteName())
	// dataR := dataRouter.Get(string(self.MuxDataRouteName))

	wrapperpath, _ := wrapperR.GetPathTemplate()
	datapath, _ := dataR.GetPathTemplate()

	self.HttpWrapperFullPath = wrapperpath
	self.HttpDataFullPath = datapath
}
