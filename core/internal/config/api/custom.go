package cfgapi

import "github.com/flarehotspot/core/internal/config"

func NewCustomCfgApi(key string, pkg string) *config.CustomConfig {
	return config.NewCustomConfig(key, pkg)
}
