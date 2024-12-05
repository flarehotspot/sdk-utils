package pkg

import (
	"os"
	"path/filepath"

	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkpkg "github.com/flarehotspot/go-utils/pkg"
)

func BuildLocalPlugins() error {
	pluginDefs := LocalPluginDefs()
	for _, def := range pluginDefs {
		pluginPath, err := sdkpkg.FindPluginSrc(def.LocalPath)
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

		if err := sdkpkg.CopyPluginFiles(pluginPath, pluginInstallDir); err != nil {
			return err
		}

	}
	return nil
}
