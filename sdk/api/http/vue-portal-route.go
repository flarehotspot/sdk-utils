package sdkhttp

import "net/http"

type VuePortalRoute struct {
	RouteName    string
	RoutePath    string
	Component    string
	HandlerFunc  http.HandlerFunc
	Middlewares  []func(http.Handler) http.Handler
}
