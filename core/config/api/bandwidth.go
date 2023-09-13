package cfgapi

import (
	"github.com/flarehotspot/core/config/bandwdcfg"
	"github.com/flarehotspot/core/sdk/api/config"
)

type BandwdCfgApi struct{}

func NewBandwdCfgApi() *BandwdCfgApi {
	return &BandwdCfgApi{}
}

func (c *BandwdCfgApi) GetConfig(ifname string) (*config.BandwdData, bool) {
	cfg, err := bandwdcfg.Read()
	if err != nil {
		return nil, false
	}

	bcfg, ok := cfg.Lans[ifname]
	if !ok {
		return nil, false
	}

	return &config.BandwdData{
		UseGlobal:       bcfg.UseGlobal,
		GlobalDownMbits: bcfg.GlobalDownMbits,
		GlobalUpMbits:   bcfg.GlobalUpMbits,
		UserDownMbits:   bcfg.UserDownMbits,
		UserUpMbits:     bcfg.UserUpMbits,
	}, true
}

func (c *BandwdCfgApi) SetConfig(ifname string, cfg *config.BandwdData) error {
	oldCfg, err := bandwdcfg.Read()
	if err != nil {
		return err
	}

	oldCfg.Lans[ifname] = &bandwdcfg.IfCfg{
		UseGlobal:       cfg.UseGlobal,
		GlobalDownMbits: cfg.GlobalDownMbits,
		GlobalUpMbits:   cfg.GlobalUpMbits,
		UserDownMbits:   cfg.UserDownMbits,
		UserUpMbits:     cfg.UserUpMbits,
	}

	return bandwdcfg.Save(oldCfg)
}
