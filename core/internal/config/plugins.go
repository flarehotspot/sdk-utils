package config

import (
	"core/internal/utils/git"
	jobque "core/internal/utils/job-que"
)

var (
	q        = jobque.NewJobQue()
	jsonFile = "plugins.json"
)

const (
	PluginSrcGit    string = "git"
	PluginSrcStore  string = "store"
	PluginSrcSystem string = "system"
	PluginSrcLocal  string = "local"
	PluginSrcZip    string = "zip"
)

type PluginMetadata struct {
	Def PluginSrcDef
}

type PluginSrcDef struct {
	Src                string // git | store | system | local | zip
	StorePackage       string // if src is "store"
	StorePluginVersion string // if src is "store"
	StoreZipUrl        string // if src is "store"
	GitURL             string // if src is "git"
	GitRef             string // can be a branch, tag or commit hash
	LocalZipFile       string // if src is "zip"
	LocalPath          string // if src is "local or system"
	InstallPath        string // where the plugin is installed
}

func (def PluginSrcDef) String() string {
	switch def.Src {
	case PluginSrcGit:
		return def.GitURL
	case PluginSrcStore:
		return def.StorePackage + "@" + def.StorePluginVersion
	case PluginSrcSystem, PluginSrcLocal:
		return def.LocalPath
	case PluginSrcZip:
		return def.LocalPath
	default:
		return "unknown plugin source: " + def.Src
	}
}

func (def PluginSrcDef) Equal(compare PluginSrcDef) bool {
	if (def.Src == PluginSrcLocal || def.Src == PluginSrcSystem) && def.LocalPath == compare.LocalPath {
		return true
	}
	if def.Src == PluginSrcGit && git.NeutralizeUrl(def.GitURL) == git.NeutralizeUrl(compare.GitURL) {
		return true
	}
	return false
}

type PluginsConfig struct {
	Recompile []string                `json:"recompile"`
	Plugins   map[string]PluginSrcDef `json:"plugins"`
}

func ReadPluginsConfig() (PluginsConfig, error) {
	empTyCfg := PluginsConfig{Recompile: []string{}, Plugins: map[string]PluginSrcDef{}}
	cfg, err := q.Exec(func() (interface{}, error) {
		var cfg PluginsConfig
		err := readConfigFile(jsonFile, &cfg)
		if err != nil {
			return empTyCfg, err
		}
		return cfg, nil
	})

	if err != nil {
		return empTyCfg, err
	}

	pluginsCfg := cfg.(PluginsConfig)
	if pluginsCfg.Plugins == nil {
		pluginsCfg.Plugins = map[string]PluginSrcDef{}
	}

	return pluginsCfg, nil
}

func WritePluginsConfig(cfg PluginsConfig) error {
	_, err := q.Exec(func() (interface{}, error) {
		return nil, writeConfigFile(jsonFile, cfg)
	})

	return err
}
