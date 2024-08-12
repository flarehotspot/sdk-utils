package config

const pluginsJsonFile = "plugins.json"

type PluginsConfig struct {
	Recompile []string `json:"recompile"`
}

func ReadPluginsConfig() (PluginsConfig, error) {
	var cfg PluginsConfig
	err := readConfigFile(pluginsJsonFile, &cfg)
	if err != nil {
		return PluginsConfig{}, err
	}
	return cfg, nil
}

func WritePluginsConfig(cfg PluginsConfig) error {
	return writeConfigFile(pluginsJsonFile, cfg)
}
