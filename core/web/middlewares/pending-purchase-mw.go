package middlewares

// import (
// 	"database/sql"
// 	"errors"
// 	"net/http"

// 	"github.com/flarehotspot/core/connmgr"
// 	"github.com/flarehotspot/core/db"
// 	"github.com/flarehotspot/core/db/models"
// 	// Ipmt "github.com/flarehotspot/core/sdk/api/payments"
// 	"github.com/flarehotspot/core/sdk/api/http"
// 	// "github.com/flarehotspot/core/web/router"
// 	// "github.com/flarehotspot/core/web/routes/names"
// )

// func PendingPurchaseMw(dtb *db.Database, mdls *models.Models, paymgr *pmt.PaymentsMgr) func(next http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			ctx := r.Context()
// 			errCode := http.StatusInternalServerError

// 			sym := ctx.Value(sdkhttp.ClientCtxKey)
// 			if sym == nil {
// 				http.Error(w, "Cannot identify device.", errCode)
// 				return
// 			}

// 			client := sym.(*connmgr.ClientDevice)
// 			device, err := mdls.Device().Find(ctx, client.Id())

// 			if err != nil {
// 				http.Error(w, err.Error(), errCode)
// 				return
// 			}

// 			purchase, err := mdls.Purchase().PendingPurchase(ctx, device.Id())
// 			if err != nil && !errors.Is(err, sql.ErrNoRows) {
// 				http.Error(w, err.Error(), errCode)
// 				return
// 			}

// 			if purchase != nil {
// 				// paymentUrl, err := router.UrlForRoute(routenames.RoutePaymentOptions)
// 				// if err != nil {
// 				// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 				// 	return
// 				// }

// 				// pr, err := Ipmt.FromPurchase(ctx, purchase)
// 				// if err != nil {
// 				// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 				// 	return
// 				// }

// 				// params, err := pr.ToQueryParams()
// 				// if err != nil {
// 				// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 				// 	return
// 				// }

// 				// http.Redirect(w, r, paymentUrl+"?"+params, http.StatusSeeOther)
// 				// return
// 			}

// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }
