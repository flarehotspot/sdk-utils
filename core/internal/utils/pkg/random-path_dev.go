//go:build dev

package pkg

import (
	"path/filepath"

	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkstr "github.com/flarehotspot/go-utils/strings"
)

func RandomPluginPath() string {
	return filepath.Join(sdkpaths.TmpDir, "plugins", sdkstr.Rand(16))
}
