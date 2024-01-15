package plugincfg

const (
	PluginSrcGit   PluginSrc = "git"
	PluginSrcStore PluginSrc = "store"
)

type PluginSrc string

// A plugin can be from store or from a git repo.
type PluginSrcDef struct {
	Src          PluginSrc `json:"src"`           // git | strore
	StorePackage string    `json:"store_pacakge"` // if src is "store"
	StoreVersion string    `json:"store_version"` // if src is "store"
	GitURL       string    `json:"git_url"`       // if src is "git"
	GitRef       string    `json:"git_ref"`       // can be a branch, tag or commit hash
}
