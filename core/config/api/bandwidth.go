package cfgapi

import (
	"github.com/flarehotspot/flarehotspot/core/config"
	"github.com/flarehotspot/sdk/api/config"
)

func NewBandwdCfgApi() *BandwdCfgApi {
	return &BandwdCfgApi{}
}

type BandwdCfgApi struct{}

func (c *BandwdCfgApi) GetConfig(ifname string) (sdkcfg.BandwdData, bool) {
	cfg, err := config.ReadBandwidthConfig()
	if err != nil {
		return sdkcfg.BandwdData{}, false
	}

	bcfg, ok := cfg.Lans[ifname]
	if !ok {
		return sdkcfg.BandwdData{}, false
	}

	return sdkcfg.BandwdData{
		UseGlobal:       bcfg.UseGlobal,
		GlobalDownMbits: bcfg.GlobalDownMbits,
		GlobalUpMbits:   bcfg.GlobalUpMbits,
		UserDownMbits:   bcfg.UserDownMbits,
		UserUpMbits:     bcfg.UserUpMbits,
	}, true
}

func (c *BandwdCfgApi) SetConfig(ifname string, cfg sdkcfg.BandwdData) error {
	oldCfg, err := config.ReadBandwidthConfig()
	if err != nil {
		return err
	}

	oldCfg.Lans[ifname] = config.IfCfg{
		UseGlobal:       cfg.UseGlobal,
		GlobalDownMbits: cfg.GlobalDownMbits,
		GlobalUpMbits:   cfg.GlobalUpMbits,
		UserDownMbits:   cfg.UserDownMbits,
		UserUpMbits:     cfg.UserUpMbits,
	}

	return config.WriteBandwidthConfig(oldCfg)
}
