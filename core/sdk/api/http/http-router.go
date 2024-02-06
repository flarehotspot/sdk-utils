package sdkhttp

import "net/http"

type MuxRouteName string
type PluginRouteName string

type HttpRouter interface {
	AdminRouter() RouterInstance
	PluginRouter() RouterInstance
	MuxRouteName(name PluginRouteName) (muxname MuxRouteName)
	UrlForMuxRoute(name MuxRouteName, pairs ...string) (url string)
	UrlForRoute(name PluginRouteName, pairs ...string) (url string)
}

type RouterInstance interface {

	// Register a handler for a GET request to the given pattern.
	Get(pattern string, handler http.HandlerFunc, middlewares ...func(next http.Handler) http.Handler) (route HttpRoute)

	// Register a handler for a POST request to the given pattern.
	Post(pattern string, handler http.HandlerFunc, middlewares ...func(next http.Handler) http.Handler) (route HttpRoute)

	// Register a subrouter for a given path
	Group(pattern string, fn func(subrouter RouterInstance))

	// Register a middleware to be used on all routes in this router instance.
	Use(middlewares ...func(next http.Handler) http.Handler)
}
