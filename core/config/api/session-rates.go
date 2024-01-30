package cfgapi

import (
	"log"

	"github.com/flarehotspot/core/config"
	sdkcfg "github.com/flarehotspot/core/sdk/api/config"
	sdkmdls "github.com/flarehotspot/core/sdk/api/models"
	sdkslices "github.com/flarehotspot/core/sdk/utils/slices"
	sdkstr "github.com/flarehotspot/core/sdk/utils/strings"
	networkutil "github.com/flarehotspot/core/utils/network"
)

func NewSessionRatesCfgApi() *SessionRatesApi {
	return &SessionRatesApi{}
}

type SessionRatesApi struct{}

func (c *SessionRatesApi) All() ([]sdkcfg.SessionRate, error) {
	cfg, err := config.ReadWifiRatesConfig()
	if err != nil {
		return nil, err
	}

	rates := []sdkcfg.SessionRate{}
	for _, r := range cfg {
		rate := sdkcfg.SessionRate(*r)
		rates = append(rates, rate)
	}

	return rates, nil
}

func (c *SessionRatesApi) AllByNet(network string) ([]sdkcfg.SessionRate, error) {
	cfg, err := config.ReadWifiRatesConfig()
	if err != nil {
		return nil, err
	}

	cfg, err = c.FilterByNet(network, cfg)
	if err != nil {
		return nil, err
	}

	rates := []sdkcfg.SessionRate{}
	for _, r := range cfg {
		rate := sdkcfg.SessionRate(*r)
		rates = append(rates, rate)
	}

	return rates, nil
}

func (c *SessionRatesApi) Save(rate sdkcfg.SessionRate) error {
	cfg, err := config.ReadWifiRatesConfig()
	if err != nil {
		return err
	}

	var exists bool
	var index int
	for i, r := range cfg {
		if r.Uuid != "" && r.Uuid == rate.Uuid {
			exists = true
			index = i
			break
		}
	}

	r := config.SessionRate(rate)
	if !exists {
		rate.Uuid = sdkstr.Rand(8)
		cfg = append(cfg, &r)
	} else {
		cfg[index] = &r
	}

	return config.WriteWifiRatesConfig(cfg)
}

func (c *SessionRatesApi) Write(rates []sdkcfg.SessionRate) ([]sdkcfg.SessionRate, error) {
	cfg := []*config.SessionRate{}
	for i, r := range rates {
		rate := config.SessionRate(r)
		if rate.Uuid == "" {
			rates[i].Uuid = sdkstr.Rand(8)
		}
		cfg = append(cfg, &rate)
	}

	log.Println("Config to save: ")
	for _, r := range rates {
		log.Println(r)
	}

	err := config.WriteWifiRatesConfig(cfg)
	if err != nil {
		return nil, err
	}

	return rates, nil
}

func (c *SessionRatesApi) ComputeSession(clientIP string, amount float64, t sdkmdls.SessionType) (sdkcfg.SessionResult, error) {
	log.Println("TODO: ComputeSession()")
	return sdkcfg.SessionResult{
		TimeMins:   100,
		DataMbytes: 100,
	}, nil
}

func (c *SessionRatesApi) FilterByNet(ip string, rates []*config.SessionRate) ([]*config.SessionRate, error) {
	rates = sdkslices.Filter(rates, func(r *config.SessionRate) bool {
		ok, err := networkutil.IpInSubnet(ip, r.Network)
		if err != nil {
			return false
		}
		return ok
	})
	return config.SortSessionRates(rates), nil
}
