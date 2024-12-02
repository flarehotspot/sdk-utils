package config

import (
	jobque "core/internal/utils/job-que"

	sdkpkg "github.com/flarehotspot/go-utils/pkg"
)

var (
	q        = jobque.NewJobQue()
	jsonFile = "plugins.json"
)

type PluginsConfig struct {
	Recompile []string                `json:"recompile"`
	Metadata  []sdkpkg.PluginMetadata `json:"metadata"`
}

func ReadPluginsConfig() (PluginsConfig, error) {
	empTyCfg := PluginsConfig{Recompile: []string{}, Metadata: []sdkpkg.PluginMetadata{}}
	cfg, err := q.Exec(func() (interface{}, error) {
		var cfg PluginsConfig
		err := readConfigFile(jsonFile, &cfg)
		if err != nil {
			return empTyCfg, err
		}
		return cfg, nil
	})

	if err != nil {
		return empTyCfg, err
	}

	pluginsCfg := cfg.(PluginsConfig)
	if pluginsCfg.Metadata == nil {
		pluginsCfg.Metadata = empTyCfg.Metadata
	}

	return pluginsCfg, nil
}

func WritePluginsConfig(cfg PluginsConfig) error {
	_, err := q.Exec(func() (interface{}, error) {
		return nil, writeConfigFile(jsonFile, cfg)
	})

	return err
}
