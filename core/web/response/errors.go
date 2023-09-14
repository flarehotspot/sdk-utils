package response

import (
	"net/http"

	"github.com/flarehotspot/core/web/router"
	Irouter "github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/sdk/utils/flash"
)

type ErrRedirect struct {
	route *Irouter.MuxRouteName
	pairs *[]string
	url   *string
}

func (e *ErrRedirect) Redirect(w http.ResponseWriter, r *http.Request, err error, pairs ...string) {
	if e.url != nil {
		flash.SetFlashMsg(w, flash.Error, err.Error())
		http.Redirect(w, r, *e.url, http.StatusSeeOther)
		return
	}

	if e.pairs != nil {
		pairs = append(*e.pairs, pairs...)
	}

	url, routeErr := router.UrlForRoute(*e.route, pairs...)
	if routeErr != nil {
		http.Error(w, routeErr.Error(), http.StatusInternalServerError)
		return
	}

	flash.SetFlashMsg(w, flash.Error, err.Error())
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func NewErrRoute(route Irouter.MuxRouteName, pairs ...string) *ErrRedirect {
	return &ErrRedirect{&route, &pairs, nil}
}

func NewErrUrl(url string) *ErrRedirect {
	return &ErrRedirect{nil, nil, &url}
}

func Error(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}
