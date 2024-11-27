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
)

type PluginMetadata struct {
	Package     string
	Def         PluginSrcDef
	InstallPath string
}

type PluginSrcDef struct {
	Src                string // git | store | system | local
	StorePackage       string // if src is "store"
	StorePluginVersion string // if src is "store"
	StoreZipUrl        string // if src is "store"
	GitURL             string // if src is "git"
	GitRef             string // can be a branch, tag or commit hash
	LocalPath          string // if src is "local or system"
}

func (def PluginSrcDef) String() string {
	switch def.Src {
	case PluginSrcGit:
		return def.GitURL
	case PluginSrcStore:
		return def.StorePackage + "@" + def.StorePluginVersion
	case PluginSrcSystem, PluginSrcLocal:
		return def.LocalPath
	default:
		return "unknown plugin source: " + def.Src
	}
}

func (def PluginSrcDef) Equal(compare PluginSrcDef) bool {
	if (def.Src == PluginSrcLocal || def.Src == PluginSrcSystem) && compare.Src == def.Src && def.LocalPath == compare.LocalPath {
		return true
	}
	if def.Src == PluginSrcGit && compare.Src == PluginSrcGit && git.NeutralizeUrl(def.GitURL) == git.NeutralizeUrl(compare.GitURL) {
		return true
	}
	if def.Src == PluginSrcStore && compare.Src == PluginSrcStore && def.StorePackage == compare.StorePackage {
		return true
	}
	return false
}

type PluginsConfig struct {
	Recompile []string         `json:"recompile"`
	Metadata  []PluginMetadata `json:"metadata"`
}

func ReadPluginsConfig() (PluginsConfig, error) {
	empTyCfg := PluginsConfig{Recompile: []string{}, Metadata: []PluginMetadata{}}
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
	if pluginsCfg.Metadata == nil {
		pluginsCfg.Metadata = empTyCfg.Metadata
	}

	return pluginsCfg, nil
}

func WritePluginsConfig(cfg PluginsConfig) error {
	_, err := q.Exec(func() (interface{}, error) {
		return nil, writeConfigFile(jsonFile, cfg)
	})

	return err
}
