package plugins

import (
	cfgapi "github.com/flarehotspot/core/config/api"
	"github.com/flarehotspot/core/sdk/api/config"
)

type ConfigApi struct {
	api *PluginApi
}

func (self *ConfigApi) Plugin() sdkcfg.IPluginCfg {
	return NewPLuginConfig(self.api)
}

func (self *ConfigApi) Application() sdkcfg.IApplicationCfg {
	return cfgapi.NewAppCfgApi()
}

func (self *ConfigApi) Database() sdkcfg.IDatabaseCfg {
	return cfgapi.NewDbCfgApi()
}

func (self *ConfigApi) SessionRates() sdkcfg.ISessionRatesCfg {
	return cfgapi.NewSessionRatesCfgApi()
}

func (self *ConfigApi) Sessions() sdkcfg.ISessionLimitsCfg {
	return cfgapi.NewSessionsCfg()
}

func (self *ConfigApi) Bandwidth() sdkcfg.IBandwdCfg {
	return cfgapi.NewBandwdCfgApi()
}

func NewConfigApi(api *PluginApi) *ConfigApi {
	return &ConfigApi{api}
}
