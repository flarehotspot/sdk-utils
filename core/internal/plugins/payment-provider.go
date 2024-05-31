package plugins

import (
	connmgr "sdk/api/connmgr"
	payments "sdk/api/payments"
	plugin "sdk/api/plugin"
)


func NewPaymentProvider(api plugin.PluginApi, provider payments.PaymentProvider) *PaymentProvider {
    prv := &PaymentProvider{api, provider}
    return prv
}

type PaymentProvider struct {
	api      plugin.PluginApi
	provider payments.PaymentProvider
}

func (self *PaymentProvider) IProvider() payments.PaymentProvider {
	return self.provider
}

func (self *PaymentProvider) PaymentOpts(clnt connmgr.ClientDevice) []payments.PaymentOpt {
	return self.provider.PaymentOpts(clnt)
}
