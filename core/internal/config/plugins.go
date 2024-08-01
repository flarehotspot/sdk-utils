package config

const pluginsJsnFile = "plugins.json"

type PluginsConfig struct {
	Recompile []string `json:"recompile"`
}

func ReadPluginsConfig() (*PluginsConfig, error) {
	var cfg PluginsConfig
	err := readConfigFile(pluginsJsnFile, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func WritePluginsConfig(cfg PluginsConfig) error {
	return writeConfigFile(pluginsJsnFile, cfg)
}
