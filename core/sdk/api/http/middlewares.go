package sdkhttp

import "net/http"

type HttpMiddleware func(next http.Handler) http.Handler

// Middlewares contains http middlewares for admin authentication, mobile device details, etc.
type Middlewares interface {
	// Returns the middleware for admin authentication.
	AdminAuth() HttpMiddleware

    // Returns middleware for caching the response. It forces browsers to cache the response for n number of days.
	CacheResponse(days int) HttpMiddleware
}
