package plugins

import (
	connmgr "github.com/flarehotspot/sdk/api/connmgr"
	payments "github.com/flarehotspot/sdk/api/payments"
	plugin "github.com/flarehotspot/sdk/api/plugin"
)

func NewPaymentMgr() *PaymentsMgr {
	return &PaymentsMgr{}
}

type PaymentsMgr struct {
	providers []*PaymentProvider
}

func (self *PaymentsMgr) Options(clnt connmgr.ClientDevice) []PaymentOption {
	opts := []PaymentOption{}
	for _, prvdr := range self.providers {
		for _, opt := range prvdr.PaymentOpts(clnt) {
			opts = append(opts, NewPaymentOpt(prvdr.api, opt))
		}
	}
	return opts
}

func (self *PaymentsMgr) FindByUuid(clnt connmgr.ClientDevice, uuid string) (PaymentOption, bool) {
	methods := self.Options(clnt)
	for _, opt := range methods {
		if opt.UUID == uuid {
			return opt, true
		}
	}
	return PaymentOption{}, false
}

func (self *PaymentsMgr) NewPaymentProvider(api plugin.PluginApi, provider payments.PaymentProvider) {
	prvdr := NewPaymentProvider(api, provider)
	self.providers = append(self.providers, prvdr)
}
