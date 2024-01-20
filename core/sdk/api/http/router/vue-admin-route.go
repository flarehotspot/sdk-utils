package router

import "net/http"

type VueAdminRoute struct {
	RouteName           string
	RoutePath           string
	Component           string
	HandlerFn           VueHandlerFn
	Middlewares         []func(http.Handler) http.Handler
	DisableCache        bool
	PermissionsRequired []string
	PermissionsAnyOf    []string
}
