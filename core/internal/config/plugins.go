package config

import jobque "core/internal/utils/job-que"

const pluginsJsonFile = "plugins.json"

var q = jobque.NewJobQues()

type PluginsConfig struct {
	Recompile []string `json:"recompile"`
}

func ReadPluginsConfig() (PluginsConfig, error) {
	cfg, err := q.Exec(func() (interface{}, error) {
		var cfg PluginsConfig
		err := readConfigFile(pluginsJsonFile, &cfg)
		if err != nil {
			return PluginsConfig{}, err
		}
		return cfg, nil
	})

	if err != nil {
		return PluginsConfig{}, err
	}

	return cfg.(PluginsConfig), nil
}

func WritePluginsConfig(cfg PluginsConfig) error {
	_, err := q.Exec(func() (interface{}, error) {
		return nil, writeConfigFile(pluginsJsonFile, cfg)
	})

	return err
}
