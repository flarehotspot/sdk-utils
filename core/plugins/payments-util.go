package plugins

import (
	"net/http"

	payments "github.com/flarehotspot/flarehotspot/core/sdk/api/payments"
	qs "github.com/flarehotspot/flarehotspot/core/sdk/libs/urlquery"
)

func ParsePurchaseReq(r *http.Request) (*payments.PurchaseRequest, error) {
	var params payments.PurchaseRequest
	query := r.URL.Query().Encode()
	err := qs.Unmarshal([]byte(query), &params)
	if err != nil {
		return nil, err
	}
	return &params, nil
}

// func ParsePaymentInfo(dtb *db.Database, mdls *models.Models, r *http.Request) (*payments.PaymentInfo, error) {
// 	ctx := r.Context()
// 	tx, err := dtb.SqlDB().BeginTx(ctx, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer tx.Rollback()

// 	token := r.URL.Query().Get("token")
// 	if token == "" {
// 		return nil, errors.New("Invalid or missing purchase token string.")
// 	}

// 	clntSym := r.Context().Value(contexts.ClientCtxKey)
// 	clnt, ok := clntSym.(connmgr.ClientDevice)
// 	if !ok {
// 		return nil, errors.New("Unable to determine client device.")
// 	}

// 	purchase, err := mdls.Purchase().FindByTokenTx(tx, ctx, token)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if purchase.DeviceId() != clnt.Id() {
// 		return nil, errors.New("Purchase request does not belong to current client device.")
// 	}

// 	if purchase.IsProcessed() {
// 		return nil, errors.New("Purchase request has already been processed.")
// 	}

// 	pmts, err := purchase.PaymentsTx(tx, ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &payments.PaymentInfo{
// 		Client:   clnt,
// 		Purchase: purchase,
// 		Payments: pmts,
// 	}, nil
// }
