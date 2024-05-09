package plugins

import (
	"net/http"

	"github.com/flarehotspot/core/internal/connmgr"
	"github.com/flarehotspot/core/internal/db"
	"github.com/flarehotspot/core/internal/db/models"
	"github.com/flarehotspot/core/internal/web/helpers"
	sdkacct "github.com/flarehotspot/sdk/api/accounts"
	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	"github.com/gorilla/mux"
)

func NewHttpApi(api *PluginApi, db *db.Database, clnt *connmgr.ClientRegister, mdls *models.Models, dmgr *connmgr.ClientRegister, pmgr *PaymentsMgr) *HttpApi {
	auth := NewHttpAuth(api)
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
	auth        *HttpAuth
	httpRouter  *HttpRouterApi
	vueRouter   *VueRouterApi
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

func (self *HttpApi) VueRouter() sdkhttp.VueRouterApi {
	return self.vueRouter
}

func (self *HttpApi) Helpers() sdkhttp.HttpHelpers {
	return NewHttpHelpers(self.api)
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

func (self *HttpApi) GetAdminNavs(acct sdkacct.Account) []sdkhttp.AdminNavList {
	return self.api.PluginsMgrApi.Utils().GetAdminNavs(acct)
}

func (self *HttpApi) GetPortalItems(clnt sdkconnmgr.ClientDevice) []sdkhttp.PortalItem {
	return self.api.PluginsMgrApi.Utils().GetPortalItems(clnt)
}
