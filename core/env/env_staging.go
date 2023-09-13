//go:build staging

package env

const (
	GO_ENV   int8   = ENV_STAGING
	HttpPort int    = 80
	BaseURL  string = "http://api.adopisoft.xyz"
)
