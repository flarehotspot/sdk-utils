//go:build !dev && !staging

package env

const (
	GO_ENV    int8   = ENV_PRODUCTION
	HTTP_PORT int    = 80
	BASE_URL  string = "http://api.adopisoft.com"
)

var (
	BuildTags string = "prod"
)
