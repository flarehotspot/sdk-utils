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
				Title:       "General",
				Description: "Some general settings",
				Fields: []sdkfields.ConfigField{
					sdkfields.TextField{
						Name:    "site_title",
						Label:   "Site Title",
						Default: "My Site",
					},
				},
			},
		}

		pcfg := cfgfields.NewPluginConfig(g.CoreAPI, c)
		pcfg.LoadConfig()

		result, err := pcfg.GetStringValue("General", "site_title")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("Result:" + result))
	}))

	router.RootRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Warning: unknown route requested: ", r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
	})
}
