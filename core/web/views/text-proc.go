package views

import (
	"strings"
	"text/template"

	"github.com/flarehotspot/core/sdk/api/http"
)

func TextProc(file string, helpers http.IHelpers, data any) (text string, err error) {
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
