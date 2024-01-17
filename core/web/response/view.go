package response

import (
	"log"
	nethttp "net/http"

	"github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/web/views"
)

func ViewWithLayout(w nethttp.ResponseWriter, layout *string, viewpath string, helpers http.IHelpers, data any) {
	html, err := views.ViewProc(layout, viewpath, helpers, data)
	if err != nil {
		Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func View(w nethttp.ResponseWriter, viewpath string, helpers http.IHelpers, data any) {
	html, err := views.ViewProc(nil, viewpath, helpers, data)
	if err != nil {
		log.Printf("View error: %+v", err)
		Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func Text(w nethttp.ResponseWriter, file string, helpers http.IHelpers, data any) {
	text, err := views.TextProc(file, helpers, data)
	if err != nil {
		log.Printf("Text response error: %+v", err)
		Error(w, err)
		return
	}

	w.Write([]byte(text))
}
