package cfgapi

import "github.com/flarehotspot/core/internal/config"

func NewPluginCfgApi(key string, pkg string) *config.PluginConfig{
    return config.NewPLuginConfig(key, pkg)
}
