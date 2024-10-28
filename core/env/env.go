//go:build !dev && !staging

package env

const (
	GO_ENV    int8   = ENV_PRODUCTION
	HTTP_PORT int    = 80
	BASE_URL  string = "http://api.adopisoft.com"
	RPC_TOKEN        = "xxxxxxxxxx"
)

var (
	BuildTags string = "prod"
)
