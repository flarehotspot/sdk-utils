package tools

import (
	"os"
	"os/exec"
	"path/filepath"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

func InstallSqlc() {
	if !sdkfs.Exists(sdkpaths.SqlcBin) {
		cmd := exec.Command("go", "build", "-buildvcs=false", "-o", sdkpaths.SqlcBin, filepath.Join(sdkpaths.SdkDir, "libs/sqlc-1.26.0/cmd/sqlc"))
		cmd.Dir = sdkpaths.AppDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}
}
