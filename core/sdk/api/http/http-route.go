package sdkhttp

// IHttpRoute represents a single route in the router.
type IHttpRoute interface {
  // Returns the name of the route.
  Name(name PluginRouteName)
}
