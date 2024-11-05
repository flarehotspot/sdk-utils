package plugins

import (
	cfgapi "core/internal/config/api"
	sdkcfg "sdk/api/config"
)

const (
	DEFAULT_CONFIG_KEY = "default"
)

func NewConfigApi(api *PluginApi) {
	cfgApi := &ConfigApi{api}
	api.ConfigAPI = cfgApi
}

type ConfigApi struct {
	api *PluginApi
}

func (self *ConfigApi) Custom(key string) sdkcfg.CustomCfgApi {
	if key == "" {
		key = DEFAULT_CONFIG_KEY
	}
	return cfgapi.NewCustomCfgApi(key, self.api.Pkg())
}

func (self *ConfigApi) Application() sdkcfg.AppCfgApi {
	return cfgapi.NewAppCfgApi()
}

func (self *ConfigApi) Bandwidth(ifname string) sdkcfg.BandwidthCfgApi {
	return cfgapi.NewBandwdCfgApi(ifname)
}
