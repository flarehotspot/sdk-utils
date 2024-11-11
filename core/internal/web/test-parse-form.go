package web

import (
	cfgfields "core/internal/config/fields"
	"core/internal/plugins"
	"core/internal/web/router"
	"core/resources/views/form"
	"fmt"
	"net/http"
)

func TestParseForm(g *plugins.CoreGlobals) {
	// flds := []sdkfields.Field{
	// 	{
	// 		Name:    "site_title",
	// 		Type:    sdkfields.FieldTypeText,
	// 		Default: "My Site",
	// 	},
	// 	{
	// 		Name: "wifi_rates",
	// 		Type: sdkfields.FieldTypeMulti,
	// 		Columns: []sdkfields.Field{
	// 			{Name: "amount", Type: sdkfields.FieldTypeNumber, Default: 1},
	// 			{Name: "session_time", Type: sdkfields.FieldTypeNumber, Default: 1},
	// 			{Name: "session_data", Type: sdkfields.FieldTypeNumber, Default: 1},
	// 		},
	// 		Default: [][]interface{}{{1, 10, 10}},
	// 	},
	// }

	// c := []sdkfields.Section{
	// 	{
	// 		Name:   "general",
	// 		Fields: flds,
	// 	},
	// }

	p := cfgfields.NewPluginConfig(g.CoreAPI)

	router.RootRouter.Handle("/form", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csrfTag := g.CoreAPI.HttpAPI.Helpers().CsrfHtmlTag(r)
		form := form.HtmlForm(csrfTag)
		form.Render(r.Context(), w)
	})).Methods("GET")

	router.RootRouter.Handle("/save-form", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := p.LoadConfig(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := p.SaveForm(r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		siteTitle, err := p.GetStringValue("general", "site_title")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		html := fmt.Sprintf("site_title: %s", siteTitle)

		rates, err := p.GetMultiValue("general", "wifi_rates")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// for i := 0; i < rates.NumRows(); i++ {
		// 	amount, err := rates.GetIntValue(i, "amount")
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}

		// 	sessionTime, err := rates.GetIntValue(i, "session_time")
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}

		// 	sessionData, err := rates.GetIntValue(i, "session_data")
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}

		// 	fmt.Fprintf(w, "---\nrow: %d\namount: %d\nsession_time: %d\nsession_data: %d\n", i, amount, sessionTime, sessionData)
		// }

		w.Write([]byte(rates.Json()))

		w.Write([]byte("<br />" + html))

	})).Methods("POST")
}
