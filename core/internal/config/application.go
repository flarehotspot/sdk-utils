package config

import (
	"github.com/flarehotspot/sdk/utils/strings"
)

const applicationJsonFile = "application.json"

type AppConfig struct {
	Lang          string `json:"lang"`
	Currency      string `json:"currency"`
	AssetsVersion string `json:"assets_version"`
	Secret        string `json:"secret"`
}

func ReadApplicationConfig() (AppConfig, error) {
	var cfg AppConfig

	err := readConfigFile(applicationJsonFile, &cfg)
	if err != nil {
		// generate defaults if not exists
		cfg := AppConfig{
			Lang:     "en",
			Currency: "USD",
			Secret:   sdkstr.Rand(16),
		}

		err = writeConfigFile(applicationJsonFile, cfg)
		if err != nil {
			return cfg, err
		}

		return cfg, nil
	}

	return cfg, nil
}

func WriteApplicationConfig(cfg AppConfig) error {
	return writeConfigFile(applicationJsonFile, cfg)
}
