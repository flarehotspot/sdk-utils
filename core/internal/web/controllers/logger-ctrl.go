package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/utils/logger"
)

func GetLogs(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		var params struct {
			Page  int `json:"page"`
			Lines int `json:"lines"`
		}

		// read request body
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("bad log viewer request body")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// check if request has body
		if len(reqBody) != 0 {
			g.CoreAPI.LoggerAPI.Info("Request has body", "body", r.Body)

			err := json.NewDecoder(r.Body).Decode(&params)
			if err != nil {
				log.Println("bad log viewer request")
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			params.Page = 1
			params.Lines = 50
		}

		linesCount := logger.GetLogLines()

		// set starting and end lines based on page and lines
		starting := linesCount - (params.Page * params.Lines)
		if starting < 0 {
			starting = 0
		}
		end := starting + params.Lines
		if end > linesCount {
			end = linesCount
		}

		// TODO : remove test print starting and end
		log.Println(starting, end)

		logs, err := logger.ReadLogs(starting, end)
		if err != nil {
			log.Println(err)
		}

		data := map[string]any{
			"logs":    logs,
			"logSize": linesCount,
		}

		res.Json(w, data, http.StatusOK)
	}
}
