package views

import (
	"html/template"
	"strings"

	"github.com/flarehotspot/core/sdk/api/http/views"
)

func ViewProc(layout *string, content string, helpers views.IViewHelpers, data any) (html string, err error) {
	tpl := "content"
	views := []string{content}

	if layout != nil {
		views = append(views, *layout)
		tpl = "layout"
	}

	templates, err := template.New("").ParseFiles(views...)
	if err != nil {
		return "", err
	}

	vdata := &ViewData{
		helpers:     helpers,
		contentData: data,
	}

	var buff strings.Builder
	err = templates.ExecuteTemplate(&buff, tpl, vdata)

	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
