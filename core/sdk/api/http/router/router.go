package router

import "net/http"

// IRouter is used to set application routes specific to your plugin.
type IRouter interface {

	// Register a handler for a GET request to the given pattern.
	Get(pattern string, h http.HandlerFunc) IRoute

	// Register a handler for a POST request to the given pattern.
	Post(pattern string, h http.HandlerFunc) IRoute

	// Register a handler for a PUT request to the given pattern.
	Put(pattern string, h http.HandlerFunc) IRoute

	// Register a handler for a PATCH request to the given pattern.
	Delete(pattern string, h http.HandlerFunc) IRoute

	// Register a handler for a OPTIONS request to the given pattern.
	Options(pattern string, h http.HandlerFunc) IRoute

	// Register a handler for a HEAD request to the given pattern.
	Group(pattern string, fn func(subrouter IRouter))

	// Register a middleware to be used on all routes in this router instance.
	Use(...func(next http.Handler) http.Handler)
}
