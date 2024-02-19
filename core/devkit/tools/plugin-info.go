package tools

import (
	"encoding/json"
	"os"
	"path/filepath"

	sdkplugin "github.com/flarehotspot/core/sdk/api/plugin"
)

func PluginInfo(pluginDir string) (sdkplugin.PluginInfo, error) {
	pluginJson := filepath.Join(pluginDir, "plugin.json")
	b, err := os.ReadFile(pluginJson)
	if err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	var info sdkplugin.PluginInfo
	err = json.Unmarshal(b, &info)
	return info, err
}
