package cfgapi

import (
	"github.com/flarehotspot/core/config/appcfg"
	"github.com/flarehotspot/core/sdk/api/config"
)

type AppCfgApi struct{}

func (c *AppCfgApi) Read() (*config.AppCfg, error) {
	cfg, err := appcfg.ReadConfig()
	if err != nil {
		return nil, err
	}

	return &config.AppCfg{
		Lang:     cfg.Lang,
		Currency: cfg.Currency,
		Secret:   cfg.Secret,
	}, nil
}

func (c *AppCfgApi) Write(cfg *config.AppCfg) error {
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
