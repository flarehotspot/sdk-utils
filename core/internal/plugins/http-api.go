package plugins

import (
	"net/http"

	"core/internal/connmgr"
	"core/internal/db"
	"core/internal/db/models"
	"core/internal/web/helpers"
	sdkconnmgr "sdk/api/connmgr"
	sdkhttp "sdk/api/http"

	"github.com/gorilla/mux"
)

func NewHttpApi(api *PluginApi, db *db.Database, clnt *connmgr.ClientRegister, mdls *models.Models, dmgr *connmgr.ClientRegister, pmgr *PaymentsMgr) {
	navs := NewNavsApi(api)
	auth := NewHttpAuth(api)
	httpResp := NewHttpResponse(api)
	middlewares := NewPluginMiddlewares(api, mdls, dmgr, pmgr)
	httpRouter := NewHttpRouterApi(api, db, clnt)
	httpForm := NewHttpFormApi(api)

	httpApi := &HttpApi{
		api:         api,
		auth:        auth,
		httpRouter:  httpRouter,
		navsApi:     navs,
		httpResp:    httpResp,
		middlewares: middlewares,
		formsApi:    httpForm,
	}

	api.HttpAPI = httpApi
	httpRouter.Init()
}

type HttpApi struct {
	api         *PluginApi
	auth        *HttpAuth
	httpRouter  *HttpRouterApi
	navsApi     *HttpNavsApi
	formsApi    *HttpFormApi
	httpResp    *HttpResponse
	middlewares *PluginMiddlewares
}

func (self *HttpApi) GetClientDevice(r *http.Request) (sdkconnmgr.ClientDevice, error) {
	return helpers.CurrentClient(self.api.ClntReg, r)
}

func (self *HttpApi) Auth() sdkhttp.HttpAuth {
	return self.auth
}

func (self *HttpApi) HttpRouter() sdkhttp.HttpRouterApi {
	return self.httpRouter
}

func (self *HttpApi) Forms() sdkhttp.HttpFormApi {
	return self.formsApi
}

func (self *HttpApi) Helpers() sdkhttp.HttpHelpers {
	return NewHttpHelpers(self.api)
}

func (self *HttpApi) Middlewares() sdkhttp.HttpMiddlewares {
	return self.middlewares
}

func (self *HttpApi) HttpResponse() sdkhttp.HttpResponse {
	return self.httpResp
}

func (self *HttpApi) MuxVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func (self *HttpApi) Navs() sdkhttp.NavsApi {
	return self.navsApi
}
