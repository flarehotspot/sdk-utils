//go:build dev

package env

const (
	GoEnv   int8   = ENV_DEV
	BaseURL  string = "http://api.adopisoft.local"
	HttpPort int    = 3000
)
