package plugincfg

// import (
// 	"encoding/json"
// 	paths "sdk/utils/paths"
// 	"log"
// 	"os"
// 	"path/filepath"
// )

// const (
// 	PluginSrcGit   PluginSrc = "git"
// 	PluginSrcStore PluginSrc = "store"
// )

// type PluginSrc string

// // A plugin can be from store or from a git repo.
// type PluginSrcDef struct {
// 	Src          PluginSrc `json:"src"`           // git | strore
// 	StorePackage string    `json:"store_pacakge"` // if src is "store"
// 	StoreVersion string    `json:"store_version"` // if src is "store"
// 	GitURL       string    `json:"git_url"`       // if src is "git"
// 	GitRef       string    `json:"git_ref"`       // can be a branch, tag or commit hash
// }

// type PluginList []*PluginSrcDef

// func readConfigFile(f string) PluginList {
// 	content, err := os.ReadFile(f)
// 	if err != nil {
// 		log.Println(err)
// 		return PluginList{}
// 	}

// 	var cfg PluginList
// 	if err = json.Unmarshal(content, &cfg); err != nil {
// 		log.Println(err)
// 		return PluginList{}
// 	}

// 	return cfg
// }

// func DefaultPluginSrc() PluginList {
// 	defaultsYaml := filepath.Join(paths.DefaultsDir, "plugins.json")
// 	log.Printf("%+v", defaultsYaml)
// 	return readConfigFile(defaultsYaml)
// }

// func UserPluginSrc() PluginList {
// 	cfgPath := filepath.Join(paths.ConfigDir, "plugins.json")
// 	return readConfigFile(cfgPath)
// }

// func AllPluginSrc() PluginList {
// 	defaultPlugins := DefaultPluginSrc()
// 	userPlugins := UserPluginSrc()
// 	return append(defaultPlugins, userPlugins...)
// }
