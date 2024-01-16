package plugins

import (
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/payments"
	"github.com/flarehotspot/core/sdk/api/http/middlewares"
	"github.com/flarehotspot/core/sdk/api/http/response"
	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/sdk/api/http/views"
	"github.com/gorilla/mux"
)

type HttpApi struct {
	api         *PluginApi
	httpRouter  *HttpRouterApi
	vueRouter   *VueRouterApi
	response    *HttpResponse
	middlewares *PluginMiddlewares
}

func NewHttpApi(api *PluginApi, mdls *models.Models, dmgr *connmgr.ClientRegister, pmgr *payments.PaymentsMgr) *HttpApi {
	httpRouter := NewRouterApi(api)
	vueRouter := NewVueRouterApi(api)
	response := NewHttpResponse(api)
	middlewares := NewPluginMiddlewares(api.db, mdls, dmgr, pmgr)
	return &HttpApi{
		api:         api,
		httpRouter:  httpRouter,
		vueRouter:   vueRouter,
		response:    response,
		middlewares: middlewares,
	}
}

func (self *HttpApi) HttpRouter() router.IHttpRouterApi {
	return self.httpRouter
}

func (self *HttpApi) VueRouter() router.IVueRouterApi {
	return self.vueRouter
}

func (self *HttpApi) Helpers(w http.ResponseWriter, r *http.Request) views.IViewHelpers {
	return NewViewHelpers(self.api, w, r)
}

func (self *HttpApi) AssetPath(path string) string {
	return filepath.Join("/plugin", self.api.Pkg(), self.api.Version(), "assets", path)
}

func (self *HttpApi) Middlewares() middlewares.Middlewares {
	return self.middlewares
}

func (self *HttpApi) Respond() response.IHttpResponse {
	return self.response
}

func (self *HttpApi) MuxVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}
