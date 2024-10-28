package sdkruntime

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

var (
	GOOS             string
	GO_VERSION       string
	GO_LONG_VERSION  string
	GO_SHORT_VERSION string
	GOARCH           string
)

func init() {
	b, err := os.ReadFile(filepath.Join(sdkpaths.AppDir, ".go-version"))
	if err != nil {
		panic(err)
	}
	v := string(b)
	GO_VERSION = strings.Replace(v, "go", "", 1)
	GO_VERSION = strings.TrimSpace(GO_VERSION)
	varr := strings.Split(GO_VERSION, ".")
	GO_SHORT_VERSION = varr[0] + "." + varr[1]
	GOARCH = runtime.GOARCH
	GOOS = runtime.GOOS
}
