package bandwdcfg

import (
	"os"
	"path/filepath"

	"github.com/flarehotspot/core/sdk/libs/yaml-3"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

var (
	cfgFile = filepath.Join(paths.ConfigDir, "bandwidth.yml")
)

type BandwdCfg struct {
	Lans map[string]*IfCfg `yaml:"lans"`
}

type IfCfg struct {
	UseGlobal       bool `yaml:"use_global"`
	GlobalDownMbits int  `yaml:"global_down_mbits"`
	GlobalUpMbits   int  `yaml:"global_up_mbits"`
	UserDownMbits   int  `yaml:"user_down_mbits"`
	UserUpMbits     int  `yaml:"user_up_mbits"`
}

func Read() (*BandwdCfg, error) {
	cfg, err := readFile(cfgFile)
	if err != nil {
		cfg, err = readDefaults()
	}

	return cfg, err
}

func Save(cfg *BandwdCfg) error {
	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(cfgFile, bytes, 0644)
}

func readFile(f string) (*BandwdCfg, error) {
	bytes, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var cfg BandwdCfg
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func readDefaults() (*BandwdCfg, error) {
	f := filepath.Join(paths.ConfigDir, ".defaults", "bandwidth.yml")
	return readFile(f)
}
