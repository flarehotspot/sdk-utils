package web

import (
	"core/internal/plugins"
	"core/internal/web/router"
	"net/http"
	sdkforms "sdk/api/forms"
	sdkhttp "sdk/api/http"
)

func TestParseForm(g *plugins.CoreGlobals) {

	c := []sdkforms.FormSection{
		{
			Name: "general",
			Fields: []sdkforms.FormField{
				sdkforms.TextField{Name: "site_title", DefaultVal: "Default Site Title"},
				sdkforms.MultiField{
					Name: "wifi_rates",
					Columns: func() []sdkforms.MultiFieldCol {
						return []sdkforms.MultiFieldCol{
							{Name: "amount", Type: sdkforms.FormFieldTypeNumber, DefaultVal: 1},
							{Name: "session_time", Type: sdkforms.FormFieldTypeNumber, DefaultVal: 1},
							{Name: "session_data", Type: sdkforms.FormFieldTypeNumber, DefaultVal: 1},
						}
					},
					DefaultVal: [][]interface{}{},
				},
			},
		},
	}

	f := sdkforms.Form{
		Name:     "default",
		Sections: c,
	}

	form, err := g.CoreAPI.HttpAPI.Forms().MakeHttpForm(f)
	if err != nil {
		panic(err)
	}

	router.RootRouter.Handle("/form", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl := form.Template(r)
		v := sdkhttp.ViewPage{
			PageContent: tpl,
		}

		g.CoreAPI.HttpAPI.HttpResponse().AdminView(w, r, v)

	})).Methods("GET")

	router.RootRouter.Handle("/save-form", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if err := r.ParseForm(); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// err := form.SaveForm(r)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// // v := sdkhttp.ViewPage{
		// // 	PageContent: form.Template(r),
		// // }

		// // g.CoreAPI.HttpAPI.HttpResponse().AdminView(w, r, v)

		// siteTitle, err := form.GetStringValue("general", "site_title")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// html := fmt.Sprintf("site_title: %s", siteTitle)

		// rates, err := form.GetMultiField("general", "wifi_rates")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// for i := 0; i < rates.NumRows(); i++ {
		// 	amount, err := rates.GetFloatValue(i, "amount")
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}

		// 	sessionTime, err := rates.GetFloatValue(i, "session_time")
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}

		// 	sessionData, err := rates.GetFloatValue(i, "session_data")
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}

		// 	fmt.Fprintf(w, "---\nrow: %d\namount: %f\nsession_time: %f\nsession_data: %f\n", i, amount, sessionTime, sessionData)
		// }

		// w.Write([]byte(rates.Json()))

		// w.Write([]byte("<br />" + html))

	})).Methods("POST")
}
