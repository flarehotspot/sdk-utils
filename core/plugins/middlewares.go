package plugins

import (
	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/web/middlewares"
)

type PluginMiddlewares struct {
	db     *db.Database
	models *models.Models
	creg   *connmgr.ClientRegister
	pmgr   *PaymentsMgr
}

func (mw *PluginMiddlewares) AdminAuth() sdkhttp.HttpMiddleware {
	return middlewares.AdminAuth
}

func (mw *PluginMiddlewares) Device() sdkhttp.HttpMiddleware {
	return middlewares.DeviceMiddleware(mw.db, mw.creg)
}

func (mw *PluginMiddlewares) CacheResponse(days int) sdkhttp.HttpMiddleware {
	return middlewares.CacheResponse(days)
}

func NewPluginMiddlewares(dtb *db.Database, mdls *models.Models, dmgr *connmgr.ClientRegister, pmgr *PaymentsMgr) *PluginMiddlewares {
	return &PluginMiddlewares{dtb, mdls, dmgr, pmgr}
}
