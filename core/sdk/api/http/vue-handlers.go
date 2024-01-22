package sdkhttp

import "net/http"

type VueHandlerFn func(w IVueResponse, r *http.Request) (err error)
type VueAdminNavsFunc func(r *http.Request) []VueAdminNav
type VuePortalItemsFunc func(r *http.Request) []VuePortalItem
