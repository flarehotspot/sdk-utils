package router

import "net/http"

type VueHandlerFn func(w VueResponseWriter, r *http.Request, vars map[string]string) (err error)
type VueAdminNavsHandler func(r *http.Request, vars map[string]string) []VueAdminNav
type VuePortalItemsHandler func(r *http.Request, vars map[string]string) []VuePortalItem

type VueResponseWriter interface {
	JsonData(data any)
	Redirect(routename string, pairs ...string)
}
