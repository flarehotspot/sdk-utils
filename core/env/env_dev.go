//go:build dev

package env

const (
	GO_ENV    int8   = ENV_DEV
	HTTP_PORT int    = 3000
	BASE_URL  string = "http://api.adopisoft.local"
	RPC_TOKEN        = "xxxxxxxxxx"
)

var (
	BuildTags string = "dev"
)
