package pkg

type PluginSrcDef struct {
	Src          string // git | strore | system | local
	StorePackage string // if src is "store"
	StoreVersion string // if src is "store"
	GitURL       string // if src is "git"
	GitRef       string // can be a branch, tag or commit hash
	LocalPath    string // if src is "local or system"
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
