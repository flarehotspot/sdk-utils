package sdkpkg

import (
	"path/filepath"
	sdkplugin "sdk/api/plugin"

	sdkfs "github.com/flarehotspot/go-utils/fs"
)

func GetInfoFromPath(src string) (sdkplugin.PluginInfo, error) {
	var info sdkplugin.PluginInfo
	if err := sdkfs.ReadJson(filepath.Join(src, "plugin.json"), &info); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	return info, nil
}
