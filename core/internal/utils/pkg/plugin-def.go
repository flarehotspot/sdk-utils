package pkg

type PluginSrcDef struct {
	Src          string `json:"src"`           // git | strore | system | local
	StorePackage string `json:"store_pacakge"` // if src is "store"
	StoreVersion string `json:"store_version"` // if src is "store"
	GitURL       string `json:"git_url"`       // if src is "git"
	GitRef       string `json:"git_ref"`       // can be a branch, tag or commit hash
	LocalPath    string `json:"local_path"`    // if src is "local or system"
}

func (def PluginSrcDef) String() string {
	switch def.Src {
	case PluginSrcGit:
		return def.GitURL
	case PluginSrcStore:
		return def.StorePackage + "@" + def.StoreVersion
	case PluginSrcSystem, PluginSrcLocal:
		return def.LocalPath
	default:
		return "unknown"
	}
}
