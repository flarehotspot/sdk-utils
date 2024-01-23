package payments

import (
	"sync"

	connmgr "github.com/flarehotspot/core/sdk/api/connmgr"
	payments "github.com/flarehotspot/core/sdk/api/payments"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"
)

type PaymentsMgr struct {
	mu        sync.RWMutex
	providers []*PaymentProvider
}

func (self *PaymentsMgr) Options(clnt connmgr.IClientDevice) []*PaymentOption {
	self.mu.RLock()
	defer self.mu.RUnlock()
	methods := []*PaymentOption{}
	for _, prvdr := range self.providers {
		for _, opt := range prvdr.PaymentOpts(clnt) {
			methods = append(methods, NewPaymentOpt(prvdr.api, opt))
		}
	}
	return methods
}

func (self *PaymentsMgr) FindByUuid(clnt connmgr.IClientDevice, uuid string) *PaymentOption {
	self.mu.RLock()
	defer self.mu.RUnlock()
	methods := self.Options(clnt)
	for _, mtd := range methods {
		if mtd.Uuid() == uuid {
			return mtd
		}
	}
	return nil
}

func (self *PaymentsMgr) NewPaymentProvider(api plugin.IPluginApi, provider payments.IPaymentProvider) {
	self.mu.Lock()
	defer self.mu.Unlock()
	prvdr := NewPaymentProvider(api, provider)
	self.providers = append(self.providers, prvdr)
}

func NewPaymentMgr() *PaymentsMgr {
	return &PaymentsMgr{
		mu: sync.RWMutex{},
	}
}
