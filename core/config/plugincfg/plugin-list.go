package plugincfg

import (
	"github.com/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

func ListDirs() []string {
	dirs, err := fs.LsDirs(paths.VendorDir, false)
	if err != nil {
		panic("Unable to list plugin directories.")
	}
	return dirs
}
