//go:build dev

package env

const (
	GO_ENV   int8   = ENV_DEV
	BaseURL  string = "http://api.adopisoft.local"
	HttpPort int    = 3000
)
