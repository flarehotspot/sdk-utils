package routes

import (
	"net/http"
	"os"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/router"
)

func FaviconRoute(g *globals.CoreGlobals) {
	router.RootRouter.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		fileBytes, err := os.ReadFile(g.CoreApi.Resource("assets/images/default-favicon-32x32.png"))
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fileBytes)
	})
}
