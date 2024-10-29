package pkg

import (
	"path/filepath"

	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

func GetInstallPath(pkg string) string {
	return filepath.Join(sdkpaths.PluginsDir, "installed", pkg)
}

func GetPendingUpdatePath(pkg string) string {
	return filepath.Join(sdkpaths.PluginsDir, "update", pkg)
}

func GetBackupPath(pkg string) string {
	return filepath.Join(sdkpaths.PluginsDir, "backup", pkg)
}

func FindDefInstallPath(def PluginSrcDef) (path string, ok bool) {
	installedPlugins := InstalledPluginsList()
	for _, p := range installedPlugins {
		if (def.Src == PluginSrcLocal || def.Src == PluginSrcSystem) && p.Def.LocalPath == def.LocalPath {
			return p.InstallPath, true
		}
		if def.Src == PluginSrcGit && p.Def.GitURL == def.GitURL {
			return p.InstallPath, true
		}
		if def.Src == PluginSrcStore && p.Def.StorePluginReleaseId == def.StorePluginReleaseId {
			return p.InstallPath, true
		}
		if def.Src == PluginSrcZip && p.Def.LocalPath == def.LocalPath {
			return p.InstallPath, true
		}
	}
	return "", false
}
