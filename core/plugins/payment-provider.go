package plugins

import (
	connmgr "github.com/flarehotspot/core/sdk/api/connmgr"
	payments "github.com/flarehotspot/core/sdk/api/payments"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"
)


func NewPaymentProvider(api plugin.IPluginApi, provider payments.IPaymentProvider) *PaymentProvider {
    prv := &PaymentProvider{api, provider}
    return prv
}

type PaymentProvider struct {
	api      plugin.IPluginApi
	provider payments.IPaymentProvider
}

func (self *PaymentProvider) IProvider() payments.IPaymentProvider {
	return self.provider
}

func (self *PaymentProvider) PaymentOpts(clnt connmgr.IClientDevice) []payments.PaymentOpt {
	return self.provider.PaymentOpts(clnt)
}
