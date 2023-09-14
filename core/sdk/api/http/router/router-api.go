package router

type IRouterApi interface {
	AdminRouter() IRouter
	PluginRouter() IRouter
	MuxRouteName(name PluginRouteName) (muxname MuxRouteName)
	UrlForMuxRoute(name MuxRouteName, pairs ...string) (url string)
	UrlForRoute(name PluginRouteName, pairs ...string) (url string)
}
