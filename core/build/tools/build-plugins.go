package tools

import (
	"core/internal/utils/pkg"
	"os"
	"path/filepath"

	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
)

func BuildLocalPlugins() error {
	pluginPaths := pkg.LocalPluginPaths()
	for _, pluginPath := range pluginPaths {
		workdir := filepath.Join(sdkpaths.TmpDir, "builds", filepath.Base(pluginPath))
		if err := pkg.BuildPlugin(pluginPath, workdir); err != nil {
			return err
		}

		info, err := pkg.GetSrcInfo(pluginPath)
		if err != nil {
			return err
		}

		pluginInstallDir := filepath.Join(sdkpaths.PluginsDir, "installed", info.Package)

		if err := os.RemoveAll(pluginInstallDir); err != nil {
			return err
		}

		for _, f := range pkg.PLuginFiles {
			if err := sdkfs.Copy(filepath.Join(pluginPath, f.File), filepath.Join(pluginInstallDir, f.File)); err != nil && !f.Optional {
				return err
			}
		}

	}
	return nil
}
