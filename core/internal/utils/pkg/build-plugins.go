package pkg

import (
	"os"
	"path/filepath"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkpkg "github.com/flarehotspot/go-utils/pkg"
)

func BuildLocalPlugins() error {
	pluginPaths := LocalPluginPaths()
	for _, searchPath := range pluginPaths {
		pluginPath, err := sdkpkg.FindPluginSrc(searchPath)
		if err != nil {
			return err
		}

		workdir := filepath.Join(sdkpaths.TmpDir, "builds", filepath.Base(pluginPath))
		defer os.RemoveAll(workdir)

		if err := BuildTemplates(pluginPath); err != nil {
			return err
		}

		if err := BuildPluginSo(pluginPath, workdir); err != nil {
			return err
		}

		info, err := sdkpkg.GetInfoFromPath(pluginPath)
		if err != nil {
			return err
		}

		pluginInstallDir := filepath.Join(sdkpaths.PluginsDir, "installed", info.Package)

		if err := os.RemoveAll(pluginInstallDir); err != nil {
			return err
		}

		for _, f := range PLuginFiles {
			if err := sdkfs.Copy(filepath.Join(pluginPath, f.File), filepath.Join(pluginInstallDir, f.File)); err != nil && !f.Optional {
				return err
			}
		}

	}
	return nil
}
