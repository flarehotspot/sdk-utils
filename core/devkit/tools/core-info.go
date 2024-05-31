package tools

import (
	"path/filepath"

	sdkplugin "sdk/api/plugin"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
)

func CoreInfo() sdkplugin.PluginInfo {
	pluginJsonPath := filepath.Join(sdkpaths.CoreDir, "plugin.json")
	var pluginDef sdkplugin.PluginInfo
    err := sdkfs.ReadJson(pluginJsonPath, &pluginDef)
	if err != nil {
		panic(err)
	}
	return pluginDef
}
