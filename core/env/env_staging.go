//go:build staging

package env

const (
	GoEnv   int8   = ENV_STAGING
	HttpPort int    = 80
	BaseURL  string = "http://api.adopisoft.xyz"
)

var (
    BuildTags string = "staging"
)
