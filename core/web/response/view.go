package response

import (
	"html/template"
	"log"
	"net/http"

	"github.com/flarehotspot/core/web/views"
)

// ViewWithLayout renders a content inside a layout.
// Note that function map 'fmap' can only be set once for each pair of layout and content
// and will be cached for the rest of the application lifetime.
func ViewWithLayout(w http.ResponseWriter, layout *views.ViewInput, content string, fmap template.FuncMap, data any) {
	fmap = views.MergeFuncMaps(views.GlobalFuncMap, layout.FuncMap, fmap)
	view := &views.ViewInput{File: content}

	tmpl, err := views.ViewProc(fmap, layout, view)
	if err != nil {
		Error(w, err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing "+content+":", err)
	}
}

// View renders a view file.
// Note that function map 'fmap' can only be set once for each view and will be cached for the rest of the application lifetime.
func View(w http.ResponseWriter, view *views.ViewInput, fmap template.FuncMap, data any) {
	fmap = views.MergeFuncMaps(views.GlobalFuncMap, view.FuncMap, fmap)
	tmpl, err := views.ViewProc(fmap, view)

	if err != nil {
		Error(w, err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing "+view.File+":", err)
	}
}
