package plugins

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/flarehotspot/core/internal/utils/crypt"
	"github.com/flarehotspot/core/internal/web/middlewares"
	"github.com/flarehotspot/core/internal/web/response"
	"github.com/flarehotspot/sdk/api/http"
	"github.com/flarehotspot/sdk/utils/strings"
	"github.com/gorilla/mux"
)

func NewVueRouteComponent(api *PluginApi, name string, path string, handler http.HandlerFunc, file string, permsReq []string, permsAny []string) *VueRouteComponent {

	compPath := filepath.Join(api.Utl.Resource("components/"), file)
	compHash, _ := crypt.SHA1Files(compPath)
	compHash = sdkstr.Sha1Hash(name, path, compPath, compHash)
	wrapperFile := api.CoreAPI.Utl.Resource("components/Wrapper.vue")

	if name == "" {
		name = "empty-route-name-" + compHash
	}

	return &VueRouteComponent{
		api:                 api,
		hash:                compHash,
		Path:                path,
		Name:                name,
		File:                file,
		WrapperFile:         wrapperFile,
		Handler:             handler,
		MuxDataRouteName:    api.HttpAPI.httpRouter.MuxRouteName(sdkhttp.PluginRouteName(name + ".data")),
		VueRouteName:        api.HttpAPI.vueRouter.MakeVueRouteName(name),
		VueRoutePath:        api.HttpAPI.vueRouter.MakeVueRoutePath(path),
		PermissionsRequired: permsReq,
		PermissionsAnyOf:    permsAny,
	}
}

type VueRouteComponent struct {
	api                 *PluginApi
	hash                string
	Path                string
	Name                string
	File                string
	WrapperFile         string
	Handler             http.HandlerFunc
	MuxDataRouteName    sdkhttp.MuxRouteName `json:"mux_data_route_name"`
	HttpDataFullPath    string               `json:"http_data_full_path"`
	HttpWrapperFullPath string               `json:"http_wrapper_full_path"`
	VueRoutePath        VueRoutePath         `json:"vue_route_path"`
	VueRouteName        string               `json:"vue_route_name"`
	PermissionsRequired []string             `json:"permissions_required"`
	PermissionsAnyOf    []string             `json:"permissions_any_of"`
}

func (self *VueRouteComponent) HttpWrapperRouteName() string {
	return fmt.Sprintf("%s.%s.%s", self.api.Pkg(), "wrapper", self.Name)
}

func (self *VueRouteComponent) HttpWrapperRoutePath() string {
	p := path.Join("/vue/components/wrapper", self.hash, "name", self.Name, "file", self.File)
	if !strings.HasSuffix(p, ".vue") {
		p = p + ".vue"
	}
	return p
}

func (self *VueRouteComponent) HttpComponentFullPath() string {
	if self.File == "" {
		return self.api.CoreAPI.HttpAPI.Helpers().VueComponentPath("Empty.vue")
	}
	return self.api.HttpAPI.Helpers().VueComponentPath(self.File)
}

func (self *VueRouteComponent) HttpDataPath() string {
	p := self.api.HttpAPI.vueRouter.VuePathToMuxPath(path.Join("/vue/data", self.hash, self.Path))
	return strings.TrimSuffix(p, "/")
}

func (self *VueRouteComponent) GetDataHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := NewVueResponse(self.api.HttpAPI.vueRouter)
		if self.Handler == nil {
			res.Json(w, nil, http.StatusOK)
			return
		}
		self.Handler(w, r)
	})
}

func (self *VueRouteComponent) GetComponentWrapperHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers := self.api.Http().Helpers()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		response.Text(w, self.WrapperFile, helpers, self)
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
