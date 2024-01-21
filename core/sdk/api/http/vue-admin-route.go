package sdkhttp

import "net/http"

type VueAdminRoute struct {
	RouteName           string
	RoutePath           string
	Component           string
	HandlerFunc           VueHandlerFn
	Middlewares         []func(http.Handler) http.Handler
	DisableCache        bool
	PermissionsRequired []string
	PermissionsAnyOf    []string
}
