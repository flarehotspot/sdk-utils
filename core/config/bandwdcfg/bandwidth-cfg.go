package bandwdcfg

import (
	"os"
	"path/filepath"

	"encoding/json"
	paths "github.com/flarehotspot/core/sdk/utils/paths"
)

var (
	cfgFile = filepath.Join(paths.ConfigDir, "bandwidth.json")
)

type BandwdCfg struct {
	Lans map[string]*IfCfg `json:"lans"`
}

type IfCfg struct {
	UseGlobal       bool `json:"use_global"`
	GlobalDownMbits int  `json:"global_down_mbits"`
	GlobalUpMbits   int  `json:"global_up_mbits"`
	UserDownMbits   int  `json:"user_down_mbits"`
	UserUpMbits     int  `json:"user_up_mbits"`
}

func Read() (*BandwdCfg, error) {
	cfg, err := readFile(cfgFile)
	if err != nil {
		cfg, err = readDefaults()
	}

	return cfg, err
}

func Save(cfg *BandwdCfg) error {
	bytes, err := json.Marshal(cfg)
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
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func readDefaults() (*BandwdCfg, error) {
	f := filepath.Join(paths.ConfigDir, ".defaults", "bandwidth.json")
	return readFile(f)
}
