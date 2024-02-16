package plugins

import (
	"context"
	"net/http"

	"github.com/flarehotspot/flarehotspot/core/db/models"
	sdkpayments "github.com/flarehotspot/sdk/api/payments"
)

func NewPurchase(api *PluginApi, ctx context.Context, deviceId int64, p *models.Purchase) *Purchase {
	return &Purchase{
		api:      api,
		ctx:      ctx,
		deviceId: deviceId,
		purchase: p,
	}
}

type Purchase struct {
	api      *PluginApi
	ctx      context.Context
	deviceId int64
	purchase *models.Purchase
}

func (p *Purchase) Name() string {
	return p.purchase.Name()
}

func (p *Purchase) FixedPrice() (float64, bool) {
	return p.purchase.FixedPrice()
}

func (p *Purchase) CreatePayment(amount float64, optname string) error {
	mdls := p.api.models
	_, err := mdls.Payment().Create(p.ctx, p.purchase.Id(), amount, optname)
	return err
}

func (p *Purchase) PayWithWallet(dbt float64) error {
	err := p.purchase.Update(p.ctx, dbt, nil, p.purchase.CancelledAt(), p.purchase.ConfirmedAt(), nil)
	return err
}

func (p *Purchase) State() (sdkpayments.PurchaseState, error) {
	state := sdkpayments.PurchaseState{}

	device, err := p.api.models.Device().Find(p.ctx, p.deviceId)
	if err != nil {
		return state, err
	}

	wallet, err := device.Wallet(p.ctx)
	if err != nil {
		return state, err
	}

	total, err := p.purchase.TotalPayment(p.ctx)
	if err != nil {
		return state, err
	}

	walletDebit := p.purchase.WalletDebit()
	walletEndBal := wallet.Balance() - walletDebit

	state.TotalPayment = total
	state.WalletDebit = walletDebit
	state.WalletEndingBal = walletEndBal
	state.WalletRealBal = wallet.Balance()

	return state, nil
}

func (p *Purchase) Execute(w http.ResponseWriter) {
	res := p.api.HttpAPI.VueResponse()
	pmgr := p.api.PluginsMgr()
	callbackPkg, ok := pmgr.FindByPkg(p.purchase.CallbackPluginPkg())
	if !ok {
		res.Error(w, "Unable to find plugin to receive the payment.", 500)
		return
	}

	callbackPkg.Http().VueResponse().Redirect(w, p.purchase.CallbackVueRouteName())
}

func (p *Purchase) Confirm() error {
	return p.purchase.Confirm(p.ctx)
}

func (p *Purchase) Cancel() error {
	return p.purchase.Cancel(p.ctx)
}
