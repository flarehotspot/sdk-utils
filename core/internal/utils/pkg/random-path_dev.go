//go:build dev

package pkg

import (
	"path/filepath"
	sdkpaths "sdk/utils/paths"
	sdkstr "sdk/utils/strings"
)

func RandomPluginPath() string {
	return filepath.Join(sdkpaths.TmpDir, "plugins", sdkstr.Rand(16))
}
