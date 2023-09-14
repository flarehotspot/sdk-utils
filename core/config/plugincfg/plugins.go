package plugincfg

import (
	"log"
	"os"
	"path/filepath"

	"github.com/flarehotspot/core/sdk/libs/yaml-3"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

type PluginList []*PluginSrcDef

func readConfigFile(f string) PluginList {
	content, err := os.ReadFile(f)
	if err != nil {
		log.Println(err)
		return PluginList{}
	}

	var cfg PluginList
	if err = yaml.Unmarshal(content, &cfg); err != nil {
		log.Println(err)
		return PluginList{}
	}

	return cfg
}

func DefaultPluginSrc() PluginList {
	defaultsYaml := filepath.Join(paths.DefaultsDir, "plugins.yml")
	log.Printf("%+v", defaultsYaml)
	return readConfigFile(defaultsYaml)
}

func UserPluginSrc() PluginList {
	cfgPath := filepath.Join(paths.ConfigDir, "plugins.yml")
	return readConfigFile(cfgPath)
}

func AllPluginSrc() PluginList {
	defaultPlugins := DefaultPluginSrc()
	userPlugins := UserPluginSrc()
	return append(defaultPlugins, userPlugins...)
}
