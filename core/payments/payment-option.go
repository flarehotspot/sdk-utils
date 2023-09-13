package payments

import (
	"github.com/flarehotspot/core/sdk/api/payments"
	"github.com/flarehotspot/core/sdk/api/plugin"
)

type PaymentOption struct {
	api  plugin.IPluginApi
	opt  payments.IPaymentOpt
}

func (self *PaymentOption) Uuid() string {
  return self.api.Name() + "::" + self.opt.Name()
}

func (self *PaymentOption) IOption() payments.IPaymentOpt {
	return self.opt
}

func NewPaymentOpt(api plugin.IPluginApi, mtd payments.IPaymentOpt) *PaymentOption {
	return &PaymentOption{api, mtd}
}
