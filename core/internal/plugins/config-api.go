package plugins

import (
	cfgapi "github.com/flarehotspot/core/internal/config/api"
	"github.com/flarehotspot/sdk/api/config"
)

func NewConfigApi(api *PluginApi) *ConfigApi {
	return &ConfigApi{api}
}

type ConfigApi struct {
	api *PluginApi
}

func (self *ConfigApi) Plugin() sdkcfg.PluginCfgApi {
	return cfgapi.NewPluginCfgApi(self.api.Pkg())
}

func (self *ConfigApi) Application() sdkcfg.AppCfgApi {
	return cfgapi.NewAppCfgApi()
}

func (self *ConfigApi) Bandwidth(ifname string) sdkcfg.BandwidthCfgApi {
	return cfgapi.NewBandwdCfgApi(ifname)
}
