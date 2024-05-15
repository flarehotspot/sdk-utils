//go:build !dev && !staging

package env

const (
	GoEnv     int8   = ENV_PRODUCTION
	BaseURL   string = "http://api.adopisoft.com"
	HttpPort  int    = 80
)

var (
	BuildTags string = "prod"
)
