package appcfg

import (
	"encoding/json"
	"os"
	"path/filepath"

	paths "github.com/flarehotspot/core/sdk/utils/paths"
	strings "github.com/flarehotspot/core/sdk/utils/strings"
)

var configPath = filepath.Join(paths.ConfigDir, "application.json")

type AppConfig struct {
	Lang          string `json:"lang"`
	Currency      string `json:"currency"`
	AssetsVersion string `json:"assets_version"`
	Secret        string `json:"secret"`
}

func genDefaults() (*AppConfig, error) {
	cfg := &AppConfig{
		Lang:     "en",
		Currency: "USD",
		Secret:   strings.Rand(16),
	}

	b, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(configPath, b, 0644)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func Read() (*AppConfig, error) {
	cbytes, err := os.ReadFile(configPath)
	if err != nil {
		return genDefaults()
	}

	var cfg AppConfig
	err = json.Unmarshal(cbytes, &cfg)
	if err != nil {
		return genDefaults()
	}

	return &cfg, nil
}

func WriteConfig(cfg *AppConfig) error {
	b, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, b, 0644)
}
