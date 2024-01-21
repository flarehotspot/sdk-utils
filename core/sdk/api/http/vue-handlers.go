package sdkhttp

import "net/http"

type VueHandlerFn func(w VueResponseWriter, r *http.Request) (err error)
type VueAdminNavsHandler func(r *http.Request) []VueAdminNav
type VuePortalItemsHandler func(r *http.Request) []VuePortalItem
