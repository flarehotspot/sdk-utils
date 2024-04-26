package plugincfg

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"encoding/json"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
	fs "github.com/flarehotspot/sdk/utils/fs"
	paths "github.com/flarehotspot/sdk/utils/paths"
)

func GetPluginInfo(pluginPath string) (*sdkplugin.PluginInfo, error) {
	dir, err := FindPluginSrc(pluginPath)
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

func FindPluginSrc(dir string) (string, error) {
	files := []string{}
	err := fs.LsFiles(dir, &files, true)
	if err != nil {
		return dir, err
	}

	for _, f := range files {
		if filepath.Base(f) == "plugin.json" {
			return filepath.Dir(f), nil
		}
	}

	return "", errors.New("Can't find plugin.json in " + paths.StripRoot(dir))
}

func GetInstallInfo(pkg string) (*sdkplugin.PluginInfo, error) {
	installPath := filepath.Join(paths.PluginsDir, pkg)
	return GetPluginInfo(installPath)
}
