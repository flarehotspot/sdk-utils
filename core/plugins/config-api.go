package plugins

import (
	cfgapi "github.com/flarehotspot/core/config/api"
	"github.com/flarehotspot/core/sdk/api/config"
)

type ConfigApi struct {
	api *PluginApi
}

func (self *ConfigApi) Plugin() sdkcfg.PluginCfg {
	return NewPLuginConfig(self.api)
}

func (self *ConfigApi) Application() sdkcfg.ApplicationCfg {
	return cfgapi.NewAppCfgApi()
}

func (self *ConfigApi) Database() sdkcfg.DbCfg {
	return cfgapi.NewDbCfgApi()
}

func (self *ConfigApi) SessionRates() sdkcfg.SessionRatesCfg {
	return cfgapi.NewSessionRatesCfgApi()
}

func (self *ConfigApi) Sessions() sdkcfg.SessionLimitsCfg {
	return cfgapi.NewSessionsCfg()
}

func (self *ConfigApi) Bandwidth() sdkcfg.BandwidthCfg {
	return cfgapi.NewBandwdCfgApi()
}

func NewConfigApi(api *PluginApi) *ConfigApi {
	return &ConfigApi{api}
}
