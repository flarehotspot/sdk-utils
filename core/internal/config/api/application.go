package cfgapi

import (
	"core/internal/config"
	"sdk/api/config"
)

func NewAppCfgApi() *AppCfgApi {
	return &AppCfgApi{}
}

type AppCfgApi struct{}

func (c *AppCfgApi) Get() (sdkcfg.AppCfg, error) {
	cfg, err := config.ReadApplicationConfig()
	if err != nil {
		return sdkcfg.AppCfg{}, err
	}

	return sdkcfg.AppCfg{
		Lang:     cfg.Lang,
		Currency: cfg.Currency,
		Secret:   cfg.Secret,
	}, nil
}

func (c *AppCfgApi) Save(cfg sdkcfg.AppCfg) error {
	data := config.AppConfig{
		Lang:     cfg.Lang,
		Currency: cfg.Currency,
		Secret:   cfg.Secret,
	}

	return config.WriteApplicationConfig(data)
}
