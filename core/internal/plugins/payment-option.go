package plugins

import (
	sdkapi "sdk/api"
)

func NewPaymentOpt(api sdkapi.IPluginApi, opt sdkapi.PaymentOpt) PaymentOption {
	uuid := api.Info().Package + "::" + opt.OptName
	return PaymentOption{api, opt, uuid}
}

type PaymentOption struct {
	api  sdkapi.IPluginApi
	Opt  sdkapi.PaymentOpt
	UUID string
}
