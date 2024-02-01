package sdkhttp

// HttpRoute represents a single route in the router.
type HttpRoute interface {
	// Returns the name of the route.
	Name(name PluginRouteName)
}
