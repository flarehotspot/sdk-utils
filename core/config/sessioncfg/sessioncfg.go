package sessioncfg

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"encoding/json"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

type SessionExp struct {
	Minutes    uint `json:"minutes"`
	Mbytes     uint `json:"mbytes"`
	PauseLimit uint `json:"pause_limit"`
	ExpDays    uint `json:"exp_days"`
}

type SessCfgData struct {
	StartOnBoot       bool          `json:"start_on_boot"`
	PauseInternetDown bool          `json:"pause_on_net_down"`
	ResumeInterUp     bool          `json:"resume_on_net_up"`
	ResumeWifiConnect bool          `json:"resume_on_wifi_conn"`
	Expirations       []*SessionExp `json:"expirations"`
}

type SessionLimitsCfg struct{}

var configPath = filepath.Join(paths.ConfigDir, "sessions.json")

func Defaults() (*SessCfgData, error) {
	cfgPath := filepath.Join(paths.DefaultsDir, "sessions.json")
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
	b, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, b, 0644)
}

func ComputeExp(exps []*SessionExp, minutes uint, mbytes uint) time.Time {
	log.Println("TODO: Compute expiration")
	return time.Now().Add(time.Duration(minutes) * time.Minute)
}

func readConfig(configPath string) (*SessCfgData, error) {
	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var cfg SessCfgData
	err = json.Unmarshal(b, &cfg)
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
