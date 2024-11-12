package web

// import (
// 	"core/internal/plugins"
// 	"core/internal/web/router"
// 	"fmt"
// 	"net/http"
// )

// func TestParseForm(g *plugins.CoreGlobals) {

// 	router.RootRouter.Handle("/form", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		g.CoreAPI.HttpAPI.Forms().NewHttpForm()

// 		csrfTag := g.CoreAPI.HttpAPI.Helpers().CsrfHtmlTag(r)
// 		form := form.HtmlForm(csrfTag)
// 		form.Render(r.Context(), w)
// 	})).Methods("GET")

// 	router.RootRouter.Handle("/save-form", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if err := r.ParseForm(); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		if err := p.LoadConfig(); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		if err := p.SaveForm(r); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		siteTitle, err := p.GetStringValue("general", "site_title")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		html := fmt.Sprintf("site_title: %s", siteTitle)

// 		rates, err := p.GetMultiValue("general", "wifi_rates")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		// for i := 0; i < rates.NumRows(); i++ {
// 		// 	amount, err := rates.GetIntValue(i, "amount")
// 		// 	if err != nil {
// 		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		// 		return
// 		// 	}

// 		// 	sessionTime, err := rates.GetIntValue(i, "session_time")
// 		// 	if err != nil {
// 		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		// 		return
// 		// 	}

// 		// 	sessionData, err := rates.GetIntValue(i, "session_data")
// 		// 	if err != nil {
// 		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		// 		return
// 		// 	}

// 		// 	fmt.Fprintf(w, "---\nrow: %d\namount: %d\nsession_time: %d\nsession_data: %d\n", i, amount, sessionTime, sessionData)
// 		// }

// 		w.Write([]byte(rates.Json()))

// 		w.Write([]byte("<br />" + html))

// 	})).Methods("POST")
// }
