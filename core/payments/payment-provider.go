package payments

import (
	"github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/api/payments"
	"github.com/flarehotspot/core/sdk/api/plugin"
)

type PaymentProvider struct {
	api      plugin.IPluginApi
	provider payments.IPaymentProvider
}

func (self *PaymentProvider) IProvider() payments.IPaymentProvider {
	return self.provider
}

func (self *PaymentProvider) PaymentOpts(clnt connmgr.IClientDevice) []payments.IPaymentOpt {
	return self.provider.PaymentOpts(clnt)
}

func NewPaymentProvider(api plugin.IPluginApi, provider payments.IPaymentProvider) *PaymentProvider {
	return &PaymentProvider{api, provider}
}
