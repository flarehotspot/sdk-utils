package plugins

import (
	"context"
	"net/http"

	"github.com/flarehotspot/core/internal/db/models"
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

func (self *Purchase) Name() string {
	return self.purchase.Name()
}

func (self *Purchase) FixedPrice() (float64, bool) {
	return self.purchase.FixedPrice()
}

func (self *Purchase) CreatePayment(amount float64, optname string) error {
	mdls := self.api.models
	_, err := mdls.Payment().Create(self.ctx, self.purchase.Id(), amount, optname)
	return err
}

func (self *Purchase) PayWithWallet(dbt float64) error {
	err := self.purchase.Update(self.ctx, dbt, nil, self.purchase.CancelledAt(), self.purchase.ConfirmedAt(), nil)
	return err
}

func (self *Purchase) State() (sdkpayments.PurchaseState, error) {
	state := sdkpayments.PurchaseState{}

	device, err := self.api.models.Device().Find(self.ctx, self.deviceId)
	if err != nil {
		return state, err
	}

	wallet, err := device.Wallet(self.ctx)
	if err != nil {
		return state, err
	}

	total, err := self.purchase.TotalPayment(self.ctx)
	if err != nil {
		return state, err
	}

	walletDebit := self.purchase.WalletDebit()
	walletEndBal := wallet.Balance() - walletDebit

	state.TotalPayment = total
	state.WalletDebit = walletDebit
	state.WalletEndingBal = walletEndBal
	state.WalletRealBal = wallet.Balance()

	return state, nil
}

func (self *Purchase) Execute(w http.ResponseWriter) {
	res := self.api.HttpAPI.VueResponse()
	pmgr := self.api.PluginsMgr()
	callbackPkg, ok := pmgr.FindByPkg(self.purchase.CallbackPluginPkg())
	if !ok {
		res.Error(w, "Unable to find plugin to receive the payment.", 500)
		return
	}

	callbackPkg.Http().VueResponse().Redirect(w, self.purchase.CallbackVueRouteName())
}

func (self *Purchase) Confirm() error {
	return self.purchase.Confirm(self.ctx)
}

func (self *Purchase) Cancel() error {
	return self.purchase.Cancel(self.ctx)
}
