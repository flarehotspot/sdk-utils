package boot

import (
	"core/internal/plugins"
	"core/internal/utils/cmd"
	"fmt"
	"os"
	"path/filepath"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

func InitOpkg(bp *plugins.BootProgress) {
	var files []string

	packagesDir := filepath.Join(sdkpaths.AppDir, "packages")
	if err := sdkfs.LsFiles(packagesDir, &files, true); err != nil {
		bp.AppendLog(fmt.Sprintf("Error listing files in packages in %s: %v", packagesDir, err.Error()))
		return
	}

	for _, f := range files {
		if filepath.Ext(f) == ".ipk" {
			bp.AppendLog("Installing ipk file: " + f)

			if err := cmd.Exec("opkg install "+f, &cmd.ExecOpts{
				Stdout: os.Stdout,
				Stderr: os.Stderr,
			}); err != nil {
				bp.AppendLog(fmt.Sprintf("Error installing ipk file %s: %v", f, err.Error()))
				return
			}

			// remove file if installed successfully
			os.Remove(f)
		}
	}
}
