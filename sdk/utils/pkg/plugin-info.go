package sdkpkg

import (
	"path/filepath"

	sdkfs "github.com/flarehotspot/go-utils/fs"
)

type PluginInfo struct {
	Name        string
	Package     string
	Description string
	Version     string
}

func GetInfoFromPath(src string) (PluginInfo, error) {
	var info PluginInfo
	if err := sdkfs.ReadJson(filepath.Join(src, "plugin.json"), &info); err != nil {
		return PluginInfo{}, err
	}

	return info, nil
}
