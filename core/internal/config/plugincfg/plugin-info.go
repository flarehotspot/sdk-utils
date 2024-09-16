package plugincfg

import (
	"core/internal/utils/pkg"
	"log"
	"os"
	"path/filepath"

	sdkplugin "sdk/api/plugin"
	"sdk/libs/go-json"
)

func GetPluginInfo(pluginPath string) (*sdkplugin.PluginInfo, error) {
	dir, err := pkg.FindPluginSrc(pluginPath)
	if err != nil {
		return nil, err
	}

	var info sdkplugin.PluginInfo
	jsonFile := filepath.Join(dir, "plugin.json")

	b, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(b, &info)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &info, nil
}
