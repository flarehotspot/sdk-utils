package response

import (
	"log"
	nethttp "net/http"

	"github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/web/views"
)

func ViewWithLayout(w nethttp.ResponseWriter, layout string, content string, helpers http.IHelpers, data any) {
	contentHtml, err := views.ViewProc(content, nil, helpers, data)
	if err != nil {
		log.Printf("View error: %+v", err)
		ErrorJson(w, err)
		return
	}

	html, err := views.ViewProc(layout, &contentHtml, helpers, data)
	if err != nil {
		log.Printf("View error: %+v", err)
		ErrorJson(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func View(w nethttp.ResponseWriter, viewpath string, helpers http.IHelpers, data any) {
	html, err := views.ViewProc(viewpath, nil, helpers, data)
	if err != nil {
		log.Printf("View error: %+v", err)
		ErrorJson(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func Text(w nethttp.ResponseWriter, file string, helpers http.IHelpers, data any) {
	text, err := views.TextProc(file, helpers, data)
	if err != nil {
		log.Printf("Text response error: %+v", err)
		ErrorJson(w, err)
		return
	}

	w.Write([]byte(text))
}
