package sdkpkg

import (
	"path/filepath"

	sdkfs "github.com/flarehotspot/go-utils/fs"
)

type PluginFile struct {
	File     string
	Optional bool
}

var PLuginFiles = []PluginFile{
	{File: "LICENSE.txt", Optional: false},
	{File: "go.mod", Optional: false},
	{File: "plugin.json", Optional: false},
	{File: "plugin.so", Optional: false},
	{File: "resources/assets/dist", Optional: true},
	{File: "resources/migrations", Optional: true},
	{File: "resources/translations", Optional: true},
}

func CopyPluginFiles(pluginSrc string, dest string) (err error) {
	if err := sdkfs.EnsureDir(dest); err != nil {
		return err
	}

	for _, f := range PLuginFiles {
		err := sdkfs.Copy(filepath.Join(pluginSrc, f.File), filepath.Join(dest, f.File))
		if err != nil && !f.Optional {
			return err
		}
	}
	return nil
}
