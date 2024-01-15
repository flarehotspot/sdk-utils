package plugincfg

import (
	"encoding/json"
	"github.com/flarehotspot/core/sdk/utils/paths"
	"log"
	"os"
	"path/filepath"
)

type PluginList []*PluginSrcDef

func readConfigFile(f string) PluginList {
	content, err := os.ReadFile(f)
	if err != nil {
		log.Println(err)
		return PluginList{}
	}

	var cfg PluginList
	if err = json.Unmarshal(content, &cfg); err != nil {
		log.Println(err)
		return PluginList{}
	}

	return cfg
}

func DefaultPluginSrc() PluginList {
	defaultsYaml := filepath.Join(paths.DefaultsDir, "plugins.json")
	log.Printf("%+v", defaultsYaml)
	return readConfigFile(defaultsYaml)
}

func UserPluginSrc() PluginList {
	cfgPath := filepath.Join(paths.ConfigDir, "plugins.json")
	return readConfigFile(cfgPath)
}

func AllPluginSrc() PluginList {
	defaultPlugins := DefaultPluginSrc()
	userPlugins := UserPluginSrc()
	return append(defaultPlugins, userPlugins...)
}
