package cfgapi

import (
	"core/internal/config"
	"sdk/api/config"
)

func NewBandwdCfgApi(ifname string) *BandwdCfgApi {
	return &BandwdCfgApi{
		ifname: ifname,
	}
}

type BandwdCfgApi struct {
	ifname string
}

func (c *BandwdCfgApi) Get() (sdkcfg.BandwdCfg, bool) {
	cfg, err := config.ReadBandwidthConfig()
	if err != nil {
		return sdkcfg.BandwdCfg{}, false
	}

	bcfg, ok := cfg.Lans[c.ifname]
	if !ok {
		return sdkcfg.BandwdCfg{}, false
	}

	return sdkcfg.BandwdCfg{
		UseGlobal:       bcfg.UseGlobal,
		GlobalDownMbits: bcfg.GlobalDownMbits,
		GlobalUpMbits:   bcfg.GlobalUpMbits,
		UserDownMbits:   bcfg.UserDownMbits,
		UserUpMbits:     bcfg.UserUpMbits,
	}, true
}

func (c *BandwdCfgApi) Save(cfg sdkcfg.BandwdCfg) error {
	oldCfg, err := config.ReadBandwidthConfig()
	if err != nil {
		return err
	}

	oldCfg.Lans[c.ifname] = config.IfCfg{
		UseGlobal:       cfg.UseGlobal,
		GlobalDownMbits: cfg.GlobalDownMbits,
		GlobalUpMbits:   cfg.GlobalUpMbits,
		UserDownMbits:   cfg.UserDownMbits,
		UserUpMbits:     cfg.UserUpMbits,
	}

	return config.WriteBandwidthConfig(oldCfg)
}
