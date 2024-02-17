package plugincfg

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"encoding/json"
	fs "github.com/flarehotspot/sdk/utils/fs"
	paths "github.com/flarehotspot/sdk/utils/paths"
)

type PluginInfo struct {
	Name        string   `json:"name"`
	Package     string   `json:"package"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	Features    []string `json:"features"`
}

func GetPluginInfo(pluginPath string) (*PluginInfo, error) {
	dir, err := FindPluginSrc(pluginPath)
	if err != nil {
		return nil, err
	}

	var info PluginInfo
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

	return "", errors.New("Can't find plugin.json in " + paths.Strip(dir))
}

func GetInstallInfo(pkg string) (*PluginInfo, error) {
	installPath := filepath.Join(paths.PluginsDir, pkg)
	return GetPluginInfo(installPath)
}
