package response

import (
	"log"
	"net/http"

	"github.com/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/core/web/views"
)

func ViewWithLayout(w http.ResponseWriter, layout *views.ViewInput, viewpath string, fmap map[string]func(), data any) {
	fmap = views.MergeFuncMaps(views.GlobalFuncMap, layout.FuncMap, fmap)
    contentIncludePath := paths.RelativeFromTo(layout.File, viewpath)

	tmpl, err := views.LayoutViewProc(fmap, layout, viewpath)
	if err != nil {
		Error(w, err)
		return
	}

	err = tmpl.Execute(w, nil, data)
	if err != nil {
		log.Println("Error executing "+viewpath+":", err)
	}
}

func View(w http.ResponseWriter, view *views.ViewInput, fmap map[string]func(), data any) {
	fmap = views.MergeFuncMaps(views.GlobalFuncMap, view.FuncMap, fmap)
	tmpl, err := views.ViewProc(fmap, view)

	if err != nil {
		Error(w, err)
		return
	}

	err = tmpl.Execute(w, nil, data)
	if err != nil {
		log.Println("Error executing "+view.File+":", err)
	}
}
