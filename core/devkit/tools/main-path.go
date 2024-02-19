package tools

import (
	"path/filepath"
	"runtime"

	sdkpaths "github.com/flarehotspot/core/sdk/utils/paths"
)

func MainFile() string {
	if runtime.GOOS == "windows" {
		return "main.exe"
	}
	return "main.app"
}

func MainPath() string {
	return filepath.Join(sdkpaths.AppDir, "main", MainFile())
}
