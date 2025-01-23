package plugins

import (
	sdkapi "sdk/api"
)

func NewPaymentProvider(api sdkapi.IPluginApi, provider sdkapi.IPaymentProvider) *PaymentProvider {
	prv := &PaymentProvider{api, provider}
	return prv
}

type PaymentProvider struct {
	api      sdkapi.IPluginApi
	provider sdkapi.IPaymentProvider
}

func (self *PaymentProvider) IProvider() sdkapi.IPaymentProvider {
	return self.provider
}

func (self *PaymentProvider) PaymentOpts(clnt sdkapi.IClientDevice) []sdkapi.PaymentOpt {
	return self.provider.PaymentOpts(clnt)
}
