package themecfg

import (
	"log"
	"os"
	"path/filepath"

	"encoding/json"
	paths "github.com/flarehotspot/core/sdk/utils/paths"
)

type ThemesConfig struct {
	Auth   string `json:"auth"`
	Portal string `json:"portal"`
	Admin  string `json:"admin"`
}

func Defaults() *ThemesConfig {
	cfgPath := filepath.Join(paths.DefaultsDir, "themes.json")
	cfg, err := readConfigFile(cfgPath)
	if err != nil {
		panic(err)
	}
	return cfg
}

func Read() *ThemesConfig {
	cfgPath := filepath.Join(paths.ConfigDir, "themes.json")
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
	err = json.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		log.Println("Error when parsing file: ", err)
		return &def, err
	}

	return &cfg, nil
}
