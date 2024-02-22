package plugins

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/flarehotspot/core/internal/connmgr"
	"github.com/flarehotspot/core/internal/db/models"
	"github.com/flarehotspot/core/internal/web/helpers"
	"github.com/flarehotspot/core/internal/web/middlewares"
	routenames "github.com/flarehotspot/core/internal/web/routes/names"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
)

func NewPluginMiddlewares(api *PluginApi, mdls *models.Models, dmgr *connmgr.ClientRegister, pmgr *PaymentsMgr) *PluginMiddlewares {
	return &PluginMiddlewares{api, mdls, dmgr, pmgr}
}

type PluginMiddlewares struct {
	api    *PluginApi
	models *models.Models
	creg   *connmgr.ClientRegister
	pmgr   *PaymentsMgr
}

func (self *PluginMiddlewares) AdminAuth() sdkhttp.HttpMiddleware {
	return middlewares.AdminAuth
}

func (self *PluginMiddlewares) Device() sdkhttp.HttpMiddleware {
	return middlewares.DeviceMiddleware(self.api.db, self.creg)
}

func (self *PluginMiddlewares) CacheResponse(days int) sdkhttp.HttpMiddleware {
	return middlewares.CacheResponse(days)
}

func (self *PluginMiddlewares) PendingPurchaseMw() sdkhttp.HttpMiddleware {
	return func(next http.Handler) http.Handler {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			errCode := http.StatusInternalServerError
			res := self.api.CoreAPI.HttpAPI.VueResponse()

			client, err := helpers.CurrentClient(self.api.ClntReg, r)
			if err != nil {
				res.FlashMsg("error", err.Error())
				res.Json(w, nil, errCode)
				return
			}

			mdls := self.api.models
			device, err := mdls.Device().Find(ctx, client.Id())
			if err != nil {
				res.FlashMsg("error", err.Error())
				res.Json(w, nil, errCode)
				return
			}

			purchase, err := mdls.Purchase().PendingPurchase(ctx, device.Id())
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				res.FlashMsg("error", err.Error())
				res.Json(w, nil, errCode)
				return
			}

			if purchase != nil {
				res.Redirect(w, routenames.RoutePaymentOptions)
				return
			}

			next.ServeHTTP(w, r)

		})

		deviceMw := self.Device()
		return deviceMw(handler)
	}

}
