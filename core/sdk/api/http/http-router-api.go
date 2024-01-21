package sdkhttp

type IHttpRouterApi interface {
	AdminRouter() IHttpRouter
	PluginRouter() IHttpRouter
	MuxRouteName(name PluginRouteName) (muxname MuxRouteName)
	UrlForMuxRoute(name MuxRouteName, pairs ...string) (url string)
	UrlForRoute(name PluginRouteName, pairs ...string) (url string)
}
