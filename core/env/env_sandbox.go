//go:build sandbox

package env

const (
	GO_ENV    int8   = ENV_SANDBOX
	HTTP_PORT int    = 80
	BASE_URL  string = "http://api.adopisoft.xyz"
	RPC_TOKEN        = "xxxxxxxxxx"
)

var (
	BuildTags string = "sandbox"
)
