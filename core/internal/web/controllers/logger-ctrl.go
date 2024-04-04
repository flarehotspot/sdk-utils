package controllers

import (
	"log"
	"net/http"

	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/utils/logger"
)

func GetLogs(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		logs, err := logger.ReadLogs()
		if err != nil {
			log.Println(err)
		}

		res.Json(w, logs, http.StatusOK)
	}
}
