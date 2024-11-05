package boot

import (
	"fmt"
	"path/filepath"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

func InitCoreLib() {
	coreLibSrc := filepath.Join(sdkpaths.CoreDir, "resources/assets/lib")
	coreLibDest := filepath.Join(sdkpaths.AppDir, "node_modules/@flarehotspot/lib")
	if !sdkfs.Exists(coreLibDest) {
		if err := sdkfs.CopyDir(coreLibSrc, coreLibDest, nil); err != nil {
			fmt.Println("Error copying core assets lib to node_modules ", err)
		}
	}
}
