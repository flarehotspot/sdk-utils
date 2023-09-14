package plugincfg

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/flarehotspot/core/sdk/libs/yaml-3"
	"github.com/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

type PluginInfo struct {
	Name        string   `yaml:"name"`
	Package     string   `yaml:"package"`
	Version     string   `yaml:"version"`
	Description string   `yaml:"description"`
	Features    []string `yaml:"features"`
}

func GetPluginInfo(pluginPath string) (*PluginInfo, error) {
    log.Println("Get plugin info from: ", pluginPath)

	dir, err := FindPluginSrc(pluginPath)
	if err != nil {
		return nil, err
	}

	var info PluginInfo
	yamlfile := filepath.Join(dir, "package.yml")

	b, err := os.ReadFile(yamlfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = yaml.Unmarshal(b, &info)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &info, nil
}

func FindPluginSrc(dir string) (string, error) {
	files, err := fs.LsFiles(dir, true)
	if err != nil {
		return dir, err
	}

	for _, f := range files {
		if filepath.Base(f) == "package.yml" {
			return filepath.Dir(f), nil
		}
	}

	return "", errors.New("Can't find package.yml in " + paths.Strip(dir))
}

func GetInstallInfo(pkg string) (*PluginInfo, error) {
	installPath := filepath.Join(paths.VendorDir, pkg)
	return GetPluginInfo(installPath)
}
