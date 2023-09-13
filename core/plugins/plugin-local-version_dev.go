//go:build dev
package plugins

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/core/sdk/utils/strings"
)

func UserLocalVersion(w io.Writer, pkg string) (ok bool) {
	clonePath := filepath.Join(paths.TmpDir, "plugins", strings.Rand(16))
	pluginPath := filepath.Join(paths.PluginsDir, pkg)
	if !fs.Exists(pluginPath) {
		return false
	}

	log.Printf("Using local version of plugin %s found in %s", pkg, pluginPath)
	err := fs.CopyDir(pluginPath, clonePath, fs.CopyOpts{Recursive: true})
	if err != nil {
		return false
	}

	err = Build(w, clonePath)
	if err != nil {
		return false
	}

	err = os.RemoveAll(filepath.Join(paths.VendorDir, pkg, "resources"))
	if err != nil {
		return false
	}

	os.Symlink(filepath.Join(paths.PluginsDir, pkg, "resources"), filepath.Join(paths.VendorDir, pkg, "resources"))
	if err != nil {
		return false
	}

	return true
}
