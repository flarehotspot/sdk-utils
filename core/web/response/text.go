package response

import (
	"net/http"
	"strings"

	httpI "github.com/flarehotspot/core/sdk/api/http"
	tmplcache "github.com/flarehotspot/core/utils/flaretmpl"
)

func Text(w http.ResponseWriter, file string, helpers httpI.IHelpers, data any) {
	vdata := &ViewData{
		ViewHelpers: helpers,
		ViewData:    data,
	}

	tmpl, err := tmplcache.GetTextTemplate(file)
	if err != nil {
		ErrorHtml(w, err.Error())
		return
	}

	var buff strings.Builder
	if err := tmpl.Execute(&buff, vdata); err != nil {
		ErrorHtml(w, err.Error())
		return
	}

	w.Write([]byte(buff.String()))
}
