package response

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	httpI "github.com/flarehotspot/flarehotspot/core/sdk/api/http"
	tmplcache "github.com/flarehotspot/flarehotspot/core/utils/flaretmpl"
)

type ViewData struct {
	PageContent template.HTML
	ViewData    any
	ViewHelpers httpI.HttpHelpers
}

func (vd *ViewData) ContentHtml() template.HTML {
	return vd.PageContent
}

func (vd *ViewData) Helpers() httpI.HttpHelpers {
	return vd.ViewHelpers
}

func (vd *ViewData) Data() any {
	return vd.ViewData
}

func ViewWithLayout(w http.ResponseWriter, layout string, content string, helpers httpI.HttpHelpers, data any) {
	contentHtml, err := viewProc(content, nil, helpers, data)
	if err != nil {
		log.Printf("View error: %+v", err)
		ErrorJson(w, err.Error(), 500)
		return
	}

	html, err := viewProc(layout, &contentHtml, helpers, data)
	if err != nil {
		log.Printf("View error: %+v", err)
		ErrorJson(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func View(w http.ResponseWriter, viewpath string, helpers httpI.HttpHelpers, data any) {
	html, err := viewProc(viewpath, nil, helpers, data)
	if err != nil {
		log.Printf("View error: %+v", err)
		ErrorJson(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func viewProc(layout string, contentHtml *template.HTML, helpers httpI.HttpHelpers, data any) (html template.HTML, err error) {
	tmpl, err := tmplcache.GetHtmlTemplate(layout)
	if err != nil {
		return "", err
	}

	vdata := &ViewData{
		ViewHelpers: helpers,
		ViewData:    data,
	}

	if contentHtml != nil {
		vdata.PageContent = *contentHtml
	}

	var output strings.Builder
	err = tmpl.Execute(&output, vdata)
	if err != nil {
		return "", err
	}

	return template.HTML(output.String()), nil
}
