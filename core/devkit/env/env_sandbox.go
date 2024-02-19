//go:build sandbox

package env

const (
	GoEnv   int8   = ENV_SANDBOX
	HttpPort int    = 80
	BaseURL  string = "http://api.adopisoft.xyz"
)
