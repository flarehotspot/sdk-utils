package cfgapi

import (
	"github.com/flarehotspot/core/config/appcfg"
	"github.com/flarehotspot/core/sdk/api/config"
)

type AppCfgApi struct{}

func (c *AppCfgApi) Read() (*sdkcfg.AppCfg, error) {
	cfg, err := appcfg.Read()
	if err != nil {
		return nil, err
	}

	return &sdkcfg.AppCfg{
		Lang:     cfg.Lang,
		Currency: cfg.Currency,
		Secret:   cfg.Secret,
	}, nil
}

func (c *AppCfgApi) Write(cfg *sdkcfg.AppCfg) error {
	data := appcfg.AppConfig{
		Lang:     cfg.Lang,
		Currency: cfg.Currency,
		Secret:   cfg.Secret,
	}

	return appcfg.WriteConfig(&data)
}

func NewAppCfgApi() *AppCfgApi {
	return &AppCfgApi{}
}
