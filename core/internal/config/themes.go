package config

import (
	"path/filepath"
	"sync"

	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
)

const (
	themesConfigJsonFile = "themes.json"
	defaultThemePlugin   = "com.flarego.default-theme"
)

var (
	themeCfgMu    = sync.RWMutex{}
	themeCfgCache *ThemesConfig
)

type ThemesConfig struct {
	Portal string `json:"portal"`
	Admin  string `json:"admin"`
}

func ReadThemesConfig() (ThemesConfig, error) {
	themeCfgMu.RLock()
	if themeCfgCache != nil {
		defer themeCfgMu.RUnlock()
		// prevent cache modification
		return ThemesConfig{
			Portal: themeCfgCache.Portal,
			Admin:  themeCfgCache.Admin,
		}, nil
	}
	themeCfgMu.RUnlock()

	var cfg ThemesConfig
	if err := readConfigFile(themesConfigJsonFile, &cfg); err != nil {
		return cfg, err
	}
	if !isThemeValid(cfg.Portal) {
		cfg.Portal = defaultThemePlugin
	}
	if !isThemeValid(cfg.Admin) {
		cfg.Admin = defaultThemePlugin
	}

	themeCfgMu.Lock()
	themeCfgCache = &cfg
	themeCfgMu.Unlock()

	return cfg, nil
}

func WriteThemesConfig(cfg ThemesConfig) error {
	if err := writeConfigFile(themesConfigJsonFile, cfg); err != nil {
		return err
	}

	themeCfgMu.Lock()
	themeCfgCache = &cfg
	themeCfgMu.Unlock()
	return nil
}

func isThemeValid(themePkg string) bool {
	themePath := filepath.Join(sdkpaths.PluginsDir, themePkg)
	return sdkfs.Exists(themePath)
}
