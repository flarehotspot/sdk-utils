package sdkhttp

import "net/http"

type VuePortalRoute struct {
	RouteName    string
	RoutePath    string
	Component    string
	HandlerFn    VueHandlerFn
    Middlewares  []func(http.Handler) http.Handler
	DisableCache bool
}
