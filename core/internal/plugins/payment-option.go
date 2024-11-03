package plugins

import (
	payments "sdk/api/payments"
	plugin "sdk/api/plugin"
)

func NewPaymentOpt(api plugin.PluginApi, opt payments.PaymentOpt) PaymentOption {
	uuid := api.Pkg() + "::" + opt.OptName
	return PaymentOption{api, opt, uuid}
}

type PaymentOption struct {
	api  plugin.PluginApi
	Opt  payments.PaymentOpt
	UUID string
}
