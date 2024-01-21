package cfgapi

import (
	"log"

	"github.com/flarehotspot/core/config/ratescfg"
	config "github.com/flarehotspot/core/sdk/api/config"
	models "github.com/flarehotspot/core/sdk/api/models"
	strings "github.com/flarehotspot/core/sdk/utils/strings"
)

type SessionRatesApi struct{}

func NewWifiRatesCfgApi() *SessionRatesApi {
	return &SessionRatesApi{}
}

func (c *SessionRatesApi) All() ([]*config.SessionRate, error) {
	cfg, err := ratescfg.Read()
	if err != nil {
		return nil, err
	}

	rates := []*config.SessionRate{}
	for _, r := range cfg {
		rate := config.SessionRate(*r)
		rates = append(rates, &rate)
	}

	return rates, nil
}

func (c *SessionRatesApi) AllByNet(network string) ([]*config.SessionRate, error) {
	cfg, err := ratescfg.Read()
	if err != nil {
		return nil, err
	}

	cfg, err = ratescfg.FilterByNet(network, cfg)
	if err != nil {
		return nil, err
	}

	rates := []*config.SessionRate{}
	for _, r := range cfg {
		rate := config.SessionRate(*r)
		rates = append(rates, &rate)
	}

	return rates, nil
}

func (c *SessionRatesApi) Save(rate *config.SessionRate) error {
	cfg, err := ratescfg.Read()
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

	r := ratescfg.WifiRate(*rate)
	if !exists {
		rate.Uuid = strings.Rand(8)
		cfg = append(cfg, &r)
	} else {
		cfg[index] = &r
	}

	return ratescfg.Write(cfg)
}

func (c *SessionRatesApi) Write(rates []*config.SessionRate) ([]*config.SessionRate, error) {
	cfg := []*ratescfg.WifiRate{}
	for i, r := range rates {
		rate := ratescfg.WifiRate(*r)
		if rate.Uuid == "" {
			rates[i].Uuid = strings.Rand(8)
		}
		cfg = append(cfg, &rate)
	}

	log.Println("Config to save: ")
	for _, r := range rates {
		log.Println(*r)
	}

	err := ratescfg.Write(cfg)
	if err != nil {
		return nil, err
	}

	return rates, nil
}

func (c *SessionRatesApi) ComputeSession(clientIP string, amount float64, t models.SessionType) (*config.SessionResult, error) {
	log.Println("TODO: ComputeSession()")
	return &config.SessionResult{
		TimeMins:   100,
		DataMbytes: 100,
	}, nil
}
