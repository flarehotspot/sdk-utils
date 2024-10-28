//go:build staging

package env

const (
	GO_ENV    int8   = ENV_STAGING
	HTTP_PORT int    = 80
	BASE_URL  string = "http://api.adopisoft.xyz"
	RPC_TOKEN        = "xxxxxxxxxx"
)

var (
	BuildTags string = "staging"
)
