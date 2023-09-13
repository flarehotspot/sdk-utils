package response

import (
	"html/template"
	"strings"
)

var (
	GlobalFuncMap = template.FuncMap{
		"htmlsafe": func(html string) template.HTML {
			return template.HTML(html)
		},
		"htmlattr": func(attr string) template.HTMLAttr {
			return template.HTMLAttr(attr)
		},
		"uppercase": func(s string) string {
			return strings.ToUpper(s)
		},
		"lowercase": func(s string) string {
			return strings.ToLower(s)
		},
	}
)

func MergeFuncMaps(maps ...template.FuncMap) template.FuncMap {
	merged := template.FuncMap{}
	for _, m := range maps {
		if m != nil {
			for k, v := range m {
				merged[k] = v
			}
		}
	}
	return merged
}
