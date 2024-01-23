package plugins

import (
	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/payments"
	"github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/web/middlewares"
)

type PluginMiddlewares struct {
	db     *db.Database
	models *models.Models
	creg   *connmgr.ClientRegister
	pmgr   *payments.PaymentsMgr
}

func (mw *PluginMiddlewares) AdminAuth() sdkhttp.HttpMiddleware {
	return middlewares.AdminAuth
}

func (mw *PluginMiddlewares) Device() sdkhttp.HttpMiddleware {
	return middlewares.DeviceMiddleware(mw.db, mw.creg)
}

func (mw *PluginMiddlewares) PendingPurchase() sdkhttp.HttpMiddleware {
	return middlewares.PendingPurchaseMw(mw.db, mw.models, mw.pmgr)
}

func (mw *PluginMiddlewares) CacheResponse(days int) sdkhttp.HttpMiddleware {
	return middlewares.CacheResponse(days)
}

func NewPluginMiddlewares(dtb *db.Database, mdls *models.Models, dmgr *connmgr.ClientRegister, pmgr *payments.PaymentsMgr) *PluginMiddlewares {
	return &PluginMiddlewares{dtb, mdls, dmgr, pmgr}
}
