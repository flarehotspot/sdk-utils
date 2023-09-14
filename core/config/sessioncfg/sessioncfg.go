package sessioncfg

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"time"

	"github.com/flarehotspot/core/sdk/libs/yaml-3"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

type SessionExp struct {
	Minutes    uint `yaml:"minutes"`
	Mbytes     uint `yaml:"mbytes"`
	PauseLimit uint `yaml:"pause_limit"`
	ExpDays    uint `yaml:"exp_days"`
}

type SessCfgData struct {
	StartOnBoot       bool          `yaml:"start_on_boot"`
	PauseInternetDown bool          `yaml:"pause_on_net_down"`
	ResumeInterUp     bool          `yaml:"resume_on_net_up"`
	ResumeWifiConnect bool          `yaml:"resume_on_wifi_conn"`
	Expirations       []*SessionExp `yaml:"expirations"`
}

type SessionLimitsCfg struct{}

var configPath = filepath.Join(paths.ConfigDir, "sessions.yml")

func Defaults() (*SessCfgData, error) {
	cfgPath := filepath.Join(paths.DefaultsDir, "sessions.yml")
	return readConfig(cfgPath)
}

func Read() (*SessCfgData, error) {
	cfg, err := readConfig(configPath)
	if err != nil {
		return Defaults()
	}

	return cfg, nil
}

func Write(cfg *SessCfgData) error {
	b, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configPath, b, 0644)
}

func ComputeExp(exps []*SessionExp, minutes uint, mbytes uint) time.Time {
	log.Println("TODO: Compute expiration")
	return time.Now().Add(time.Duration(minutes) * time.Minute)
}

func readConfig(configPath string) (*SessCfgData, error) {
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var cfg SessCfgData
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	cfg.Expirations = sortExpDESC(cfg.Expirations)

	return &cfg, nil
}

// sortExpDESC Sort expirations from greatest to least.
func sortExpDESC(exps []*SessionExp) []*SessionExp {
	sort.Slice(exps, func(i, j int) bool {
		return exps[i].Minutes > exps[j].Minutes || exps[i].Mbytes > exps[j].Mbytes
	})
	return exps
}

func sortExpASC(exps []*SessionExp) []*SessionExp {
	sort.Slice(exps, func(i, j int) bool {
		return exps[i].Minutes < exps[j].Minutes || exps[i].Mbytes < exps[j].Mbytes
	})
	return exps
}
