package sdkhttp

import "net/http"

type VueAdminRoute struct {
	RouteName   string
	RoutePath   string
	Component   string
	HandlerFunc http.HandlerFunc
	Middlewares []func(http.Handler) http.Handler
	PermitFn    func(perms []string) bool
}
