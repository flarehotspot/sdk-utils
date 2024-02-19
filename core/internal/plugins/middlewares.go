package plugins

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/flarehotspot/core/internal/connmgr"
	"github.com/flarehotspot/core/internal/db/models"
	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/internal/web/helpers"
	"github.com/flarehotspot/core/internal/web/middlewares"
	routenames "github.com/flarehotspot/core/internal/web/routes/names"
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

func (mw *PluginMiddlewares) AdminAuth() sdkhttp.HttpMiddleware {
	return middlewares.AdminAuth
}

func (mw *PluginMiddlewares) Device() sdkhttp.HttpMiddleware {
	return middlewares.DeviceMiddleware(mw.api.db, mw.creg)
}

func (mw *PluginMiddlewares) CacheResponse(days int) sdkhttp.HttpMiddleware {
	return middlewares.CacheResponse(days)
}

func (mw *PluginMiddlewares) PendingPurchaseMw() sdkhttp.HttpMiddleware {
	return func(next http.Handler) http.Handler {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			errCode := http.StatusInternalServerError
			res := mw.api.CoreAPI.HttpAPI.VueResponse()

			client, err := helpers.CurrentClient(mw.api.ClntReg, r)
			if err != nil {
				res.FlashMsg("error", err.Error())
				res.Json(w, nil, errCode)
				return
			}

			mdls := mw.api.models
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

		deviceMw := mw.Device()
		return deviceMw(handler)
	}

}
