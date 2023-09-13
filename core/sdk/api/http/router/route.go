package router

// IRoute represents a single route in the router.
type IRoute interface {
  // Returns the name of the route.
  Name(name PluginRouteName)
}
