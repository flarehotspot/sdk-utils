package plugincfg

const (
	PluginSrcGit   PluginSrc = "git"
	PluginSrcStore PluginSrc = "store"
)

type PluginSrc string

// A plugin can be from store or from a git repo.
type PluginSrcDef struct {
	Src          PluginSrc `yaml:"src"`           // git | strore
	StorePackage string    `yaml:"store_pacakge"` // if src is "store"
	StoreVersion string    `yaml:"store_version"` // if src is "store"
	GitURL       string    `yaml:"git_url"`       // if src is "git"
	GitRef       string    `yaml:"git_ref"`       // can be a branch, tag or commit hash
}
