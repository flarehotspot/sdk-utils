//go:build !dev && !staging

package env

const (
	GO_ENV   int8   = ENV_PRODUCTION
	BaseURL  string = "http://api.adopisoft.com"
	HttpPort int    = 80
)
