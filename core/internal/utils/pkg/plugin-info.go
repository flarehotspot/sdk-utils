package pkg

import (
	"path/filepath"
	sdkplugin "sdk/api/plugin"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
)

func GetPluginInfo(def PluginSrcDef) (sdkplugin.PluginInfo, error) {
	path, ok := FindDefInstallPath(def)
	if !ok {
		return sdkplugin.PluginInfo{}, ErrNotInstalled
	}

	return GetSrcInfo(path)
}

func GetSrcInfo(src string) (sdkplugin.PluginInfo, error) {
	var info sdkplugin.PluginInfo
	if err := sdkfs.ReadJson(filepath.Join(src, "plugin.json"), &info); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	return info, nil
}

func CoreInfo() sdkplugin.PluginInfo {
	pluginJsonPath := filepath.Join(sdkpaths.CoreDir, "plugin.json")
	var pluginDef sdkplugin.PluginInfo
	if err := sdkfs.ReadJson(pluginJsonPath, &pluginDef); err != nil {
		panic(err)
	}
	return pluginDef
}
