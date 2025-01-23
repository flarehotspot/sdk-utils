package main

import (
	"core/build/tools"
	"core/internal/utils/pkg"
	"os"
	"path/filepath"

	sdkutils "github.com/flarehotspot/sdk-utils"
)

func main() {
	tools.SyncCoreVersion()
	tools.SyncGoVersion()
	version := pkg.GetCoreInfo().Version
	releaseNotePath := filepath.Join(sdkutils.PathCoreDir, "build", "release-notes", version+".md")
	if !sdkutils.FsExists(releaseNotePath) {
		if err := os.WriteFile(releaseNotePath, []byte("## "+version+"\n\n"), 0644); err != nil {
			panic(err)
		}
		return
	}
}
