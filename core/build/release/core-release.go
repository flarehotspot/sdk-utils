package release

import (
	"core/build/tools"
	"fmt"
	"path/filepath"
	"runtime"
	sdkpaths "sdk/utils/paths"
)

var (
	coreInfo     = tools.CoreInfo()
	releaseDir   = filepath.Join(sdkpaths.AppDir, "core-release", fmt.Sprintf("core-%s-%s-1", coreInfo.Version, runtime.GOARCH))
	releaseFiles = []string{
        "config/.defaults",
		"core/go-version",
		"core/go.mod",
		"core/plugin.json",
		"core/plugin.so",
        "core/resources",
		"sdk",
	}
)

func CreateRelease() {
	tools.BuildCore()
}
