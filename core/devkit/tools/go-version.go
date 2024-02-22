package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	sdkpaths "github.com/flarehotspot/sdk/utils/paths"
)

func GoVersion() (string, error) {
	goVersionPath := filepath.Join(sdkpaths.CoreDir, "go-version")
	b, err := os.ReadFile(goVersionPath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	v := string(b)
	return strings.TrimSpace(v), nil
}

// GoShortVersion returns the short version of the core go version. For example if go version is "1.19.12" it returns "1.19"
func GoShortVersion() (string, error) {
	v, err := GoVersion()
	if err != nil {
		return "", err
	}
	varr := strings.Split(v, ".")
	return varr[0] + "." + varr[1], nil
}
