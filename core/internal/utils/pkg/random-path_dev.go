//go:build dev

package pkg

import (
	"path/filepath"

	sdkutils "github.com/flarehotspot/sdk-utils"
)

func RandomPluginPath() string {
	return filepath.Join(sdkutils.PathTmpDir, "plugins", sdkutils.RandomStr(16))
}
