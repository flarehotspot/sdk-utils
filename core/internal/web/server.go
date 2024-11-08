package web

import (
	"log"
	"net/http"
	sdkfields "sdk/api/config/fields"

	cfgfields "core/internal/config/fields"
	"core/internal/plugins"
	"core/internal/web/router"
	"core/internal/web/routes"
)

func SetupBootRoutes(g *plugins.CoreGlobals) {
	routes.AssetsRoutes(g)
	routes.BootRoutes(g)
	routes.CoreAssets(g)
}

func SetupAllRoutes(g *plugins.CoreGlobals) {
	routes.AssetsRoutes(g)
	routes.PortalRoutes(g)
	routes.AdminRoutes(g)
	routes.PaymentRoutes(g)

	router.RootRouter.Handle("/config", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := []sdkfields.Section{
			{
				Name:        "general",
				Description: "Some general settings",
				Fields: []sdkfields.ConfigField{
					sdkfields.TextField{Name: "site_title", Label: "Site Title", Default: "My Site"},
					sdkfields.TextField{Name: "color", Label: "Color", Default: "red"},
				},
			},
		}

		pcfg := cfgfields.NewPluginConfig(g.CoreAPI, c)
		pcfg.LoadConfig()

		siteTitle, err := pcfg.GetStringValue("general", "site_title")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		color, err := pcfg.GetStringValue("general", "color")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("site title:" + siteTitle + "<br>"))
		w.Write([]byte("color:" + color))
	}))

	router.RootRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Warning: unknown route requested: ", r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
	})
}
