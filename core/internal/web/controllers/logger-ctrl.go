package controllers

import (
	"net/http"
	"time"

	"github.com/flarehotspot/core/internal/plugins"
)

func GetLogs(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		logs := map[string]string{
			"title":    "testing out",
			"datetime": time.Now().String(),
		}

		res.Json(w, logs, http.StatusOK)
	}
}
