package sdkhttp

import "net/http"

type MuxRouteName string
type PluginRouteName string

type IHttpRouter interface {
	AdminRouter() IRouter
	PluginRouter() IRouter
	MuxRouteName(name PluginRouteName) (muxname MuxRouteName)
	UrlForMuxRoute(name MuxRouteName, pairs ...string) (url string)
	UrlForRoute(name PluginRouteName, pairs ...string) (url string)
}

type IRouter interface {

	// Register a handler for a GET request to the given pattern.
	Get(pattern string, h http.HandlerFunc) (route IHttpRoute)

	// Register a handler for a POST request to the given pattern.
	Post(pattern string, h http.HandlerFunc) (route IHttpRoute)

	// Register a subrouter for a given path
	Group(pattern string, fn func(subrouter IRouter))

	// Register a middleware to be used on all routes in this router instance.
	Use(...func(next http.Handler) http.Handler)
}
