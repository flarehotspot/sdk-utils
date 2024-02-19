package tools

import (
	"path/filepath"

	sdkplugin "github.com/flarehotspot/core/sdk/api/plugin"
	sdkfs "github.com/flarehotspot/core/sdk/utils/fs"
	sdkpaths "github.com/flarehotspot/core/sdk/utils/paths"
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
