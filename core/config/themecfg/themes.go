package themecfg

import (
	"log"
	"os"
	"path/filepath"

	"github.com/flarehotspot/core/sdk/libs/yaml-3"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

type ThemesConfig struct {
	Auth          string `yaml:"auth"`
	CaptivePortal string `yaml:"captive-portal"`
	WebAdmin      string `yaml:"web-admin"`
}

func Defaults() *ThemesConfig {
	cfgPath := filepath.Join(paths.DefaultsDir, "themes.yml")
	cfg, err := readConfigFile(cfgPath)
	if err != nil {
		panic(err)
	}
	return cfg
}

func Read() *ThemesConfig {
	cfgPath := filepath.Join(paths.ConfigDir, "themes.yml")
	cfg, err := readConfigFile(cfgPath)
	if err != nil {
		return Defaults()
	}

	return cfg
}

func readConfigFile(f string) (*ThemesConfig, error) {
	def := ThemesConfig{}
	cfgBytes, err := os.ReadFile(f)
	if err != nil {
		return &def, err
	}

	var cfg ThemesConfig
	err = yaml.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		log.Println("Error when parsing file: ", err)
		return &def, err
	}

	return &cfg, nil
}
