package plugins

import (
	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/payments"
	"github.com/flarehotspot/core/web/middlewares"
	mwI "github.com/flarehotspot/core/sdk/api/http/middlewares"
)

type PluginMiddlewares struct {
	db     *db.Database
	models *models.Models
	creg   *connmgr.ClientRegister
	pmgr   *payments.PaymentsMgr
}

func (mw *PluginMiddlewares) AdminAuth() mwI.HttpMiddleware {
	return middlewares.AdminAuth
}

func (mw *PluginMiddlewares) Device() mwI.HttpMiddleware {
	return middlewares.DeviceMiddleware(mw.db, mw.creg)
}

func (mw *PluginMiddlewares) PendingPurchase() mwI.HttpMiddleware {
	return middlewares.PendingPurchaseMw(mw.db, mw.models, mw.pmgr)
}

func NewPluginMiddlewares(dtb *db.Database, mdls *models.Models, dmgr *connmgr.ClientRegister, pmgr *payments.PaymentsMgr) *PluginMiddlewares {
	return &PluginMiddlewares{dtb, mdls, dmgr, pmgr}
}
