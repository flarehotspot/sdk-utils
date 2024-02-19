package tools

import (
	"path/filepath"

	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
	sdkfs "github.com/flarehotspot/sdk/utils/fs"
	sdkpaths "github.com/flarehotspot/sdk/utils/paths"
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
