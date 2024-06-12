package sdkruntime

import (
	"runtime"
	"strings"
)

var (
	GOOS             string
	GOVERSION        string
	GO_SHORT_VERSION string
	GOARCH           string
)

func init() {
	v := runtime.Version()
	GOVERSION = strings.Replace(v, "go", "", 1)
	varr := strings.Split(GOVERSION, ".")
	GO_SHORT_VERSION = varr[0] + "." + varr[1]
	GOARCH = runtime.GOARCH
	GOOS = runtime.GOOS
}
