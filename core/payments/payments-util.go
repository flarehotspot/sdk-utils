package payments

import (
	"errors"
	"net/http"

	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/api/models"
	"github.com/flarehotspot/core/sdk/api/payments"
	qs "github.com/flarehotspot/core/sdk/libs/urlquery"
	"github.com/flarehotspot/core/sdk/utils/contexts"
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

func ParsePaymentInfo(dtb *db.Database, mdls models.IModelsApi, r *http.Request) (*payments.PaymentInfo, error) {
	ctx := r.Context()
	tx, err := dtb.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	token := r.URL.Query().Get("token")
	if token == "" {
		return nil, errors.New("Invalid or missing purchase token string.")
	}

	clntSym := r.Context().Value(contexts.ClientCtxKey)
	clnt, ok := clntSym.(connmgr.IClientDevice)
	if !ok {
		return nil, errors.New("Unable to determine client device.")
	}

	purchase, err := mdls.Purchase().FindByTokenTx(tx, ctx, token)
	if err != nil {
		return nil, err
	}

	if purchase.DeviceId() != clnt.Id() {
		return nil, errors.New("Purchase request does not belong to current client device.")
	}

	if purchase.IsProcessed() {
		return nil, errors.New("Purchase request has already been processed.")
	}

	pmts, err := purchase.PaymentsTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &payments.PaymentInfo{
		Client:   clnt,
		Purchase: purchase,
		Payments: pmts,
	}, nil
}
