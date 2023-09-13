package cfgapi

import (
	"log"

	"github.com/flarehotspot/core/config/sessioncfg"
	"github.com/flarehotspot/core/sdk/api/config"
)

type SessionLimitsCfg struct{}

func (c *SessionLimitsCfg) Read() (*config.SessCfgData, error) {
	cfg, err := sessioncfg.Read()
	if err != nil {
		return nil, err
	}

	return sessionCfgToOut(cfg), nil
}

func (c *SessionLimitsCfg) Write(data *config.SessCfgData) error {
	cfg := inputToSessionCfg(data)
	return sessioncfg.Write(cfg)
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

func inputToSessionCfg(input *config.SessCfgData) *sessioncfg.SessCfgData {
	cfg := sessioncfg.SessCfgData{
		StartOnBoot:       input.StartOnBoot,
		PauseInternetDown: input.PauseInternetDown,
		ResumeInterUp:     input.ResumeInterUp,
		ResumeWifiConnect: input.ResumeWifiConnect,
	}

	for _, exp := range input.PauseLimitDenoms {
		cfg.Expirations = append(cfg.Expirations, &sessioncfg.SessionExp{
			Minutes:    exp.Minutes,
			Mbytes:     exp.Mbytes,
			PauseLimit: exp.PauseLimit,
			ExpDays:    exp.ExpDays,
		})
	}

	return &cfg
}

func sessionCfgToOut(cfg *sessioncfg.SessCfgData) *config.SessCfgData {
	data := config.SessCfgData{
		StartOnBoot:       cfg.StartOnBoot,
		PauseInternetDown: cfg.PauseInternetDown,
		ResumeInterUp:     cfg.ResumeInterUp,
		ResumeWifiConnect: cfg.ResumeWifiConnect,
	}

	for _, exp := range cfg.Expirations {
		data.PauseLimitDenoms = append(data.PauseLimitDenoms, &config.ExpPauseDenom{
			Minutes:    exp.Minutes,
			Mbytes:     exp.Mbytes,
			PauseLimit: exp.PauseLimit,
			ExpDays:    exp.ExpDays,
		})
	}

	return &data
}
