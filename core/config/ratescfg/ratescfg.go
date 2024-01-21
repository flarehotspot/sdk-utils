package ratescfg

import (
	"net"
	"os"
	"path/filepath"
	"sort"

	"encoding/json"
	paths "github.com/flarehotspot/core/sdk/utils/paths"
	slices "github.com/flarehotspot/core/sdk/utils/slices"
)

type WifiRate struct {
	Uuid       string  `json:"uuid"`
	Network    string  `json:"network"`
	Amount     float64 `json:"amount"`
	TimeMins   uint    `json:"time_mins"`
	DataMbytes uint    `json:"data_mbytes"`
}

var configPath = filepath.Join(paths.ConfigDir, "session-rates.json")

func Defaults() ([]*WifiRate, error) {
	configPath := filepath.Join(paths.DefaultsDir, "session-rates.json")
	return readFile(configPath)
}

func Read() ([]*WifiRate, error) {
	cfg, err := readFile(configPath)
	if err != nil {
		return Defaults()
	}

	return sortRates(cfg), nil
}

func Write(cfg []*WifiRate) error {
	for _, r := range cfg {
		err := validateRate(r)
		if err != nil {
			return err
		}
	}

	b, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, b, 0644)
}

func FilterByNet(ip string, rates []*WifiRate) ([]*WifiRate, error) {
	rates = slices.Filter(rates, func(r *WifiRate) bool {
		ok, err := ipInSubnet(ip, r.Network)
		if err != nil {
			return false
		}
		return ok
	})
	return sortRates(rates), nil
}

func readFile(configPath string) ([]*WifiRate, error) {
	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var rates []*WifiRate

	err = json.Unmarshal(b, &rates)
	if err != nil {
		return nil, err
	}

	return rates, nil
}

func validateRate(rate *WifiRate) error {
	_, _, err := net.ParseCIDR(rate.Network)
	return err
}

func ipInSubnet(ip string, networkIp string) (bool, error) {
	testIP := net.ParseIP(ip)
	_, subnet, err := net.ParseCIDR(networkIp)
	if err != nil {
		return false, err
	}
	return subnet.Contains(testIP), nil
}

func sortRates(rates []*WifiRate) []*WifiRate {
	sort.Slice(rates, func(i, j int) bool {
		r1 := rates[i]
		r2 := rates[j]
		return r1.Amount < r2.Amount
	})
	return rates
}
