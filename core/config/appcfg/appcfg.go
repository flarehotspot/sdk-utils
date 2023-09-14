package appcfg

import (
	"os"
	"path/filepath"

	"github.com/flarehotspot/core/sdk/libs/yaml-3"
	"github.com/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/core/sdk/utils/strings"
)

var configPath = filepath.Join(paths.ConfigDir, "application.yml")

type AppConfig struct {
	Lang     string `yaml:"lang"`
	Currency string `yaml:"currency"`
	Secret   string `yaml:"secret"`
}

func genDefaults() (*AppConfig, error) {
	cfg := &AppConfig{
		Lang:     "en",
		Currency: "USD",
		Secret:   strings.Rand(16),
	}

	b, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(configPath, b, 0644)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func ReadConfig() (*AppConfig, error) {
	cbytes, err := os.ReadFile(configPath)
	if err != nil {
		return genDefaults()
	}

	var cfg AppConfig
	err = yaml.Unmarshal(cbytes, &cfg)
	if err != nil {
		return genDefaults()
	}

	return &cfg, nil
}

func WriteConfig(cfg *AppConfig) error {
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, b, 0644)
}
