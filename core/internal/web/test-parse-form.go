package web

import (
	cfgfields "core/internal/config/fields"
	"core/internal/plugins"
	"core/internal/web/router"
	"core/resources/views/form"
	"fmt"
	"net/http"
	sdkfields "sdk/api/config/fields"
)

func TestParseForm(g *plugins.CoreGlobals) {
	flds := []sdkfields.ConfigField{
		sdkfields.TextField{
			Name:    "site_title",
			Default: "My Site",
		},
	}

	c := []sdkfields.Section{
		{
			Name:   "general",
			Fields: flds,
		},
	}

	p := cfgfields.NewPluginConfig(g.CoreAPI, c)

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

		w.Write([]byte(html))

	})).Methods("POST")
}
