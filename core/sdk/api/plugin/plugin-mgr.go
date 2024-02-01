package sdkplugin

// PluginsMgrApi is used to get data of installed plugins in the system.
type PluginsMgrApi interface {

	// Find a plugin by name as defined in package.yml "name" field.
	FindByName(name string) (PluginApi, bool)

	// Find a plugin by path as defined in package.yml "package" field.
	FindByPkg(pkg string) (PluginApi, bool)

	// Returns all plugins installed in the system.
	All() []PluginApi
}
