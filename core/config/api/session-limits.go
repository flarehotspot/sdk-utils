package cfgapi

import (
	"log"

	"github.com/flarehotspot/core/config"
	"github.com/flarehotspot/core/sdk/api/config"
)

type SessionLimitsCfg struct{}

func (c *SessionLimitsCfg) Read() (sdkcfg.SessCfgData, error) {
	cfg, err := config.ReadSessionSettingsConfig()
	if err != nil {
		return sdkcfg.SessCfgData{}, err
	}

	return sessionCfgToOut(cfg), nil
}

func (c *SessionLimitsCfg) Write(data sdkcfg.SessCfgData) error {
	cfg := inputToSessionCfg(data)
	return config.WriteSessionSettingsConfig(cfg)
}

func (c *SessionLimitsCfg) ComputeExpDays(minutes uint, mbytes uint) uint {
	log.Println("TODO: ComputeExpDays()")
	return 1
}

func (c *SessionLimitsCfg) ComputePauseLimit(minutes uint, mbytes uint) (limit uint) {
	log.Println("TODO: ComputePauseLimit()")
	return 1
}

func NewSessionsCfg() *SessionLimitsCfg {
	return &SessionLimitsCfg{}
}

func inputToSessionCfg(input sdkcfg.SessCfgData) config.SessCfgData {
	cfg := config.SessCfgData{
		StartOnBoot:       input.StartOnBoot,
		PauseInternetDown: input.PauseInternetDown,
		ResumeInterUp:     input.ResumeInterUp,
		ResumeWifiConnect: input.ResumeWifiConnect,
	}

	for _, exp := range input.PauseLimitDenoms {
		cfg.Expirations = append(cfg.Expirations, config.SessionExp{
			Minutes:    exp.Minutes,
			Mbytes:     exp.Mbytes,
			PauseLimit: exp.PauseLimit,
			ExpDays:    exp.ExpDays,
		})
	}

	return cfg
}

func sessionCfgToOut(cfg config.SessCfgData) sdkcfg.SessCfgData {
	data := sdkcfg.SessCfgData{
		StartOnBoot:       cfg.StartOnBoot,
		PauseInternetDown: cfg.PauseInternetDown,
		ResumeInterUp:     cfg.ResumeInterUp,
		ResumeWifiConnect: cfg.ResumeWifiConnect,
	}

	for _, exp := range cfg.Expirations {
		data.PauseLimitDenoms = append(data.PauseLimitDenoms, sdkcfg.ExpPauseDenom{
			Minutes:    exp.Minutes,
			Mbytes:     exp.Mbytes,
			PauseLimit: exp.PauseLimit,
			ExpDays:    exp.ExpDays,
		})
	}

	return data
}
