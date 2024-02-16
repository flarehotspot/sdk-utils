package response

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	httpI "github.com/flarehotspot/flarehotspot/core/sdk/api/http"
	tmplcache "github.com/flarehotspot/flarehotspot/core/utils/flaretmpl"
)

func File(w http.ResponseWriter, file string, helpers httpI.HttpHelpers, data any) {
	tmpl, err := tmplcache.GetTextTemplate(file)
	if err != nil {
		ErrorHtml(w, err.Error())
		return
	}

	vdata := &ViewData{
		ViewHelpers: helpers,
		ViewData:    data,
	}

	var buff strings.Builder
	err = tmpl.Execute(&buff, vdata)

	if err != nil {
		ErrorHtml(w, err.Error())
		return
	}

	info, err := os.Stat(file)
	if err != nil {
		ErrorHtml(w, err.Error())
		return
	}

	lastModified := info.ModTime().UTC().Format(http.TimeFormat)
	contentType := mime.TypeByExtension(filepath.Ext(file))

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Last-Modified", lastModified)
	w.Write([]byte(buff.String()))
}
