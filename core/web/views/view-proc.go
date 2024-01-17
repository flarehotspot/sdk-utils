package views

import (
	"html/template"
	"os"
	"strings"

	"github.com/flarehotspot/core/sdk/api/http"
)

func ViewProc(layout string, contentHtml *template.HTML, helpers http.IHelpers, data any) (html template.HTML, err error) {
	content, err := os.ReadFile(layout)
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("").Parse(string(content))
	if err != nil {
		return "", err
	}

	vdata := &ViewData{
		helpers:     helpers,
		contentData: data,
	}

	if contentHtml != nil {
		vdata.contentHtml = *contentHtml
	}

	var output strings.Builder
	err = tmpl.Execute(&output, vdata)
	if err != nil {
		return "", err
	}

	return template.HTML(output.String()), nil
}
