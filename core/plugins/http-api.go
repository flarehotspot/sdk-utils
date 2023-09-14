package plugins

import (
	"html/template"
	"log"
	"net/http"

	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/payments"
	"github.com/flarehotspot/core/sdk/api/http/middlewares"
	"github.com/flarehotspot/core/sdk/api/http/response"
	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/gorilla/mux"
)

type HttpApi struct {
	api      *PluginApi
	router   *RouterApi
	response *HttpResponse
	mw       *PluginMiddlewares
}

func NewHttpApi(api *PluginApi, mdls *models.Models, dmgr *connmgr.ClientRegister, pmgr *payments.PaymentsMgr) *HttpApi {
	prouter := NewRouterApi(api)
	response := NewHttpResponse(api)
	mw := NewPluginMiddlewares(api.db, mdls, dmgr, pmgr)
	return &HttpApi{
		api:      api,
		router:   prouter,
		response: response,
		mw:       mw,
	}
}

func (self *HttpApi) Router() router.IRouterApi {
	return self.router
}

func (self *HttpApi) Middlewares() middlewares.Middlewares {
	return self.mw
}

func (self *HttpApi) Respond() response.IHttpResponse {
	return self.response
}

func (self *HttpApi) MuxVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func (self *HttpApi) ViewFuncMap(fmap template.FuncMap) {
	if self.api.vfmap != nil {
		log.Println("Warning: api.HttpApi().ViewFuncMap() can only be set once, discarding subsequent calls.")
		return
	}
	self.api.vfmap = fmap
}
