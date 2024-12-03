package pkg

import (
	"path/filepath"
	sdkplugin "sdk/api/plugin"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkpkg "github.com/flarehotspot/go-utils/pkg"
)

func GetInfoFromDef(def sdkpkg.PluginSrcDef) (info sdkplugin.PluginInfo, err error) {
	path, ok := FindDefInstallPath(def)
	if !ok {
		return info, ErrNotInstalled
	}

	return sdkpkg.GetInfoFromPath(path)
}

func GetCoreInfo() sdkplugin.PluginInfo {
	pluginJsonPath := filepath.Join(sdkpaths.CoreDir, "plugin.json")
	var pluginDef sdkplugin.PluginInfo
	if err := sdkfs.ReadJson(pluginJsonPath, &pluginDef); err != nil {
		panic(err)
	}
	return pluginDef
}
