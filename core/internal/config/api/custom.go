package cfgapi

import "core/internal/config"

func NewCustomCfgApi(key string, pkg string) *config.CustomConfig {
	return config.NewCustomConfig(key, pkg)
}
