package plugins

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/flarehotspot/core/internal/connmgr"
	"github.com/flarehotspot/core/internal/db/models"
	"github.com/flarehotspot/core/internal/web/helpers"
	"github.com/flarehotspot/core/internal/web/middlewares"
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

func (self *PluginMiddlewares) AdminAuth() func(http.Handler) http.Handler {
	return middlewares.AdminAuth
}

func (self *PluginMiddlewares) Device() func(http.Handler) http.Handler {
	return middlewares.DeviceMiddleware(self.api.db, self.creg)
}

func (self *PluginMiddlewares) CacheResponse(days int) func(http.Handler) http.Handler {
	return middlewares.CacheResponse(days)
}

func (self *PluginMiddlewares) PendingPurchase() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			errCode := http.StatusInternalServerError
			res := self.api.CoreAPI.HttpAPI.VueResponse()

			client, err := helpers.CurrentClient(self.api.ClntReg, r)
			if err != nil {
				res.SendFlashMsg(w, "error", err.Error(), errCode)
				return
			}

			mdls := self.api.models
			device, err := mdls.Device().Find(ctx, client.Id())
			if err != nil {
				res.SendFlashMsg(w, "error", err.Error(), errCode)
				return
			}

			purchase, err := mdls.Purchase().PendingPurchase(ctx, device.Id())
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				res.SendFlashMsg(w, "error", err.Error(), errCode)
				return
			}

			if purchase != nil {
                res.Redirect(w, "payments:customer:options")
				return
			}

			next.ServeHTTP(w, r)

		})

		deviceMw := self.Device()
		return deviceMw(handler)
	}

}
