package plugins

import (
	"context"

	"github.com/flarehotspot/core/db/models"
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

func (p *Purchase) TotalPayments() (float64, error) {
	return p.purchase.TotalPayments(p.ctx)
}

func (p *Purchase) WalletDebit() float64 {
	return p.purchase.WalletDebit()
}

func (p *Purchase) WalletBalance() (float64, error) {
	wallet, err := p.api.models.Wallet().FindByDevice(p.ctx, p.deviceId)
	if err != nil {
		return 0, err
	}
	return wallet.Balance() - p.WalletDebit(), nil
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
