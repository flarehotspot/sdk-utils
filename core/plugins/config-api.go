package plugins

import (
	cfgapi "github.com/flarehotspot/core/config/api"
	Icfg "github.com/flarehotspot/core/sdk/api/config"
)

type ConfigApi struct {
	api *PluginApi
}

func (self *ConfigApi) Plugin() Icfg.IPluginCfg {
	return NewPLuginConfig(self.api)
}

func (self *ConfigApi) Application() Icfg.IApplicationCfg {
	return cfgapi.NewAppCfgApi()
}

func (self *ConfigApi) Database() Icfg.IDatabaseCfg {
	return cfgapi.NewDbCfgApi()
}

func (self *ConfigApi) WifiRates() Icfg.ISessionRatesCfg {
	return cfgapi.NewWifiRatesCfgApi()
}

func (self *ConfigApi) Sessions() Icfg.ISessionLimitsCfg {
	return cfgapi.NewSessionsCfg()
}

func (self *ConfigApi) Bandwidth() Icfg.IBandwdCfg {
	return cfgapi.NewBandwdCfgApi()
}

func NewConfigApi(api *PluginApi) *ConfigApi {
	return &ConfigApi{api}
}
