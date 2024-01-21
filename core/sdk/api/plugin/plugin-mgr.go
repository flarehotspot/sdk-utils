package sdkplugin

// IPluginMgr is used to get data of installed plugins in the system.
type IPluginMgr interface {

	// Find a plugin by name as defined in package.yml "name" field.
	FindByName(name string) (IPluginApi, bool)

	// Find a plugin by path as defined in package.yml "package" field.
	FindByPkg(pkg string) (IPluginApi, bool)

	// Returns all plugins installed in the system.
	All() []IPluginApi
}
