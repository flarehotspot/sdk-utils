package views

var (
	GlobalFuncMap = map[string]func(){
		// "htmlsafe": func(html string) template.HTML {
		// 	return template.HTML(html)
		// },
		// "htmlattr": func(attr string) template.HTMLAttr {
		// 	return template.HTMLAttr(attr)
		// },
		// "uppercase": func(s string) string {
		// 	return strings.ToUpper(s)
		// },
		// "lowercase": func(s string) string {
		// 	return strings.ToLower(s)
		// },
	}
)

func MergeFuncMaps(maps ...map[string]func()) map[string]func() {
	merged := map[string]func(){}
	for _, m := range maps {
		if m != nil {
			for k, v := range m {
				merged[k] = v
			}
		}
	}
	return merged
}
