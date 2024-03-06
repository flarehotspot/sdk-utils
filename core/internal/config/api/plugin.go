package cfgapi

import "github.com/flarehotspot/core/internal/config"

func NewPluginCfgApi(pkg string) *config.PluginConfig{
    return config.NewPLuginConfig(pkg)
}
