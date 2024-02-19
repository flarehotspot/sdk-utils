package config

import (
	"log"
	"sort"
	"time"
)

const sessionSettingsJonFile = "sessions.json"

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
	Expirations       []SessionExp `json:"expirations"`
}

func ReadSessionSettingsConfig() (SessCfgData, error) {
	var cfg SessCfgData
	err := readConfigFile(&cfg, sessionSettingsJonFile)
	return cfg, err
}

func WriteSessionSettingsConfig(cfg SessCfgData) error {
	return writeConfigFile(sessionSettingsJonFile, cfg)
}

func ComputeSessionExpiration(exps []SessionExp, minutes uint, mbytes uint) time.Time {
	log.Println("TODO: Compute expiration")
	return time.Now().Add(time.Duration(minutes) * time.Minute)
}

// sortExpDESC Sort expirations from greatest to least.
func SortSessionExpirationDESC(exps []SessionExp) []SessionExp {
	sort.Slice(exps, func(i, j int) bool {
		return exps[i].Minutes > exps[j].Minutes || exps[i].Mbytes > exps[j].Mbytes
	})
	return exps
}

func SortSessionExpirationExpASC(exps []SessionExp) []SessionExp {
	sort.Slice(exps, func(i, j int) bool {
		return exps[i].Minutes < exps[j].Minutes || exps[i].Mbytes < exps[j].Mbytes
	})
	return exps
}
