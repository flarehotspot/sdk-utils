package config

import (
	"net"
	"sort"

	sdkslices "github.com/flarehotspot/flarehotspot/core/sdk/utils/slices"
	networkutil "github.com/flarehotspot/flarehotspot/core/utils/network"
)

const sessionRatesJsonFile = "session-rates.json"

type SessionRate struct {
	Uuid       string  `json:"uuid"`
	Network    string  `json:"network"`
	Amount     float64 `json:"amount"`
	TimeMins   uint    `json:"time_mins"`
	DataMbytes uint    `json:"data_mbytes"`
}

func (rate *SessionRate) validateRate() error {
	_, _, err := net.ParseCIDR(rate.Network)
	return err
}

func ReadWifiRatesConfig() ([]*SessionRate, error) {
	var cfg []*SessionRate
	if err := readConfigFile(&cfg, sessionRatesJsonFile); err != nil {
		return cfg, err
	}

	return SortSessionRates(cfg), nil
}

func WriteWifiRatesConfig(cfg []*SessionRate) error {
	for _, r := range cfg {
		err := r.validateRate()
		if err != nil {
			return err
		}
	}

	return writeConfigFile(sessionRatesJsonFile, cfg)
}

func SortSessionRates(rates []*SessionRate) []*SessionRate {
	sort.Slice(rates, func(i, j int) bool {
		r1 := rates[i]
		r2 := rates[j]
		return r1.Amount < r2.Amount
	})
	return rates
}

func FilterSessionRatesByNet(ip string, rates []*SessionRate) ([]*SessionRate, error) {
	rates = sdkslices.Filter(rates, func(r *SessionRate) bool {
		ok, err := networkutil.IpInSubnet(ip, r.Network)
		if err != nil {
			return false
		}
		return ok
	})

	return SortSessionRates(rates), nil
}

// func validateRate(rate WifiRate) error {
// 	_, _, err := net.ParseCIDR(rate.Network)
// 	return err
// }

// func ipInSubnet(ip string, networkIp string) (bool, error) {
// 	testIP := net.ParseIP(ip)
// 	_, subnet, err := net.ParseCIDR(networkIp)
// 	if err != nil {
// 		return false, err
// 	}
// 	return subnet.Contains(testIP), nil
// }
