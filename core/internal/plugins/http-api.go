package plugins

import (
	"net/http"

	"github.com/flarehotspot/core/internal/connmgr"
	"github.com/flarehotspot/core/internal/db"
	"github.com/flarehotspot/core/internal/db/models"
	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	"github.com/flarehotspot/core/internal/web/helpers"
	"github.com/gorilla/mux"
)

func NewHttpApi(api *PluginApi, db *db.Database, clnt *connmgr.ClientRegister, mdls *models.Models, dmgr *connmgr.ClientRegister, pmgr *PaymentsMgr) *HttpApi {
	auth := NewAuthApi(api)
	httpRouter := NewHttpRouterApi(api, db, clnt)
	vueRouter := NewVueRouterApi(api)
	httpResp := NewHttpResponse(api)
	middlewares := NewPluginMiddlewares(api, mdls, dmgr, pmgr)

	return &HttpApi{
		api:         api,
		auth:        auth,
		httpRouter:  httpRouter,
		vueRouter:   vueRouter,
		httpResp:    httpResp,
		middlewares: middlewares,
	}
}

type HttpApi struct {
	api         *PluginApi
	auth        *AuthApi
	httpRouter  *HttpRouterApi
	vueRouter   *VueRouterApi
	httpResp    *HttpResponse
	middlewares *PluginMiddlewares
}

func (self *HttpApi) Auth() sdkhttp.HttpAuth {
	return self.auth
}

func (self *HttpApi) GetDevice(r *http.Request) (sdkconnmgr.ClientDevice, error) {
	return helpers.CurrentClient(self.api.ClntReg, r)
}

func (self *HttpApi) HttpRouter() sdkhttp.HttpRouter {
	return self.httpRouter
}

func (self *HttpApi) VueRouter() sdkhttp.VueRouterApi {
	return self.vueRouter
}

func (self *HttpApi) Helpers() sdkhttp.HttpHelpers {
	return NewViewHelpers(self.api)
}

func (self *HttpApi) Middlewares() sdkhttp.Middlewares {
	return self.middlewares
}

func (self *HttpApi) HttpResponse() sdkhttp.HttpResponse {
	return self.httpResp
}

func (self *HttpApi) VueResponse() sdkhttp.VueResponse {
	return NewVueResponse(self.api.HttpAPI.vueRouter)
}

func (self *HttpApi) MuxVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func (self *HttpApi) GetAdminNavs(r *http.Request) []sdkhttp.AdminNavList {
	return self.api.PluginsMgrApi.Utils().GetAdminNavs(r)
}

func (self *HttpApi) GetPortalItems(r *http.Request) []sdkhttp.PortalItem {
	return self.api.PluginsMgrApi.Utils().GetPortalItems(r)
}
