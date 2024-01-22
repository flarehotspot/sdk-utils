package plugins

import (
	"net/http"

	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/payments"
	"github.com/flarehotspot/core/sdk/api/http"
	"github.com/gorilla/mux"
)

type HttpApi struct {
	api         *PluginApi
	httpRouter  *HttpRouterApi
	vueRouter   *VueRouterApi
	response    *HttpResponse
	middlewares *PluginMiddlewares
}

func NewHttpApi(api *PluginApi, db *db.Database, clnt *connmgr.ClientRegister, mdls *models.Models, dmgr *connmgr.ClientRegister, pmgr *payments.PaymentsMgr) *HttpApi {
	httpRouter := NewRouterApi(api, db, clnt)
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

func (self *HttpApi) HttpRouter() sdkhttp.IHttpRouterApi {
	return self.httpRouter
}

func (self *HttpApi) VueRouter() sdkhttp.IVueRouterApi {
	return self.vueRouter
}

func (self *HttpApi) Helpers() sdkhttp.IHelpers {
	return NewViewHelpers(self.api)
}

func (self *HttpApi) Middlewares() sdkhttp.Middlewares {
	return self.middlewares
}

func (self *HttpApi) HttpResponse() sdkhttp.IHttpResponse {
	return self.response
}

func (self *HttpApi) MuxVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func (self *HttpApi) VueRespones(w http.ResponseWriter, r *http.Request) *VueResponse {
	return NewVueResponse(self.api.HttpAPI.vueRouter, w, r)
}
