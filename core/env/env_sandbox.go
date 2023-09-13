//go:build sandbox

package env

const (
	GO_ENV   int8   = ENV_SANDBOX
	HttpPort int    = 80
	BaseURL  string = "http://api.adopisoft.xyz"
)
