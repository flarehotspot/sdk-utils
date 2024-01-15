package response

import (
	"log"
	"net/http"

	sdkviews "github.com/flarehotspot/core/sdk/api/http/views"
	"github.com/flarehotspot/core/web/views"
)

func ViewWithLayout(w http.ResponseWriter, layout *string, viewpath string, helpers sdkviews.IViewHelpers, data any) {
	html, err := views.ViewProc(layout, viewpath, helpers, data)
	if err != nil {
		Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func View(w http.ResponseWriter, viewpath string, helpers sdkviews.IViewHelpers, data any) {
	html, err := views.ViewProc(nil, viewpath, helpers, data)
	if err != nil {
		log.Printf("View error: %+v", err)
		Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func Text(w http.ResponseWriter, file string, helpers sdkviews.IViewHelpers, data any) {
	text, err := views.TextProc(file, helpers, data)
	if err != nil {
		log.Printf("Text response error: %+v", err)
		Error(w, err)
		return
	}

	w.Write([]byte(text))
}
