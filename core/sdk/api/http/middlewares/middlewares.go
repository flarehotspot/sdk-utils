package middlewares

import "net/http"

type HttpMiddleware func(next http.Handler) http.Handler

// Middlewares contains http middlewares for admin authentication, mobile device details, etc.
type Middlewares interface {
	// Returns the middleware for admin authentication.
	AdminAuth() HttpMiddleware

	// Returns the middleware for mobile device details.
	Device() HttpMiddleware
}
