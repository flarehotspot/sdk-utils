package response

import (
	sdkviews "github.com/flarehotspot/core/sdk/api/http/views"
	"github.com/flarehotspot/core/web/views"
	"net/http"
)

func ViewWithLayout(w http.ResponseWriter, layout *views.ViewInput, viewpath string, helpers sdkviews.IViewHelpers, data any) {
	content := views.ViewInput{File: viewpath}
	html, err := views.ViewProc(layout, content, helpers, data)
	if err != nil {
		Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func View(w http.ResponseWriter, view *views.ViewInput, helpers sdkviews.IViewHelpers, data any) {
	html, err := views.ViewProc(nil, *view, helpers, data)
	if err != nil {
		Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
