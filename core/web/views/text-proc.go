package views

import (
	"strings"
	"text/template"

	"github.com/flarehotspot/core/sdk/api/http/views"
)

func TextProc(file string, helpers views.IViewHelpers, data any) (text string, err error) {
	templates, err := template.ParseFiles(file)
	if err != nil {
		return "", err
	}

	vdata := &ViewData{
		helpers:     helpers,
		contentData: data,
	}

	var buff strings.Builder
	err = templates.Execute(&buff, vdata)

	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
