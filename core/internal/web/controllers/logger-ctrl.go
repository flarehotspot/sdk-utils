package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/utils/logger"
)

func GetLogs(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		var params struct {
			Page  int
			Lines int
		}

		rows := logger.GetLogLines()

		// new approach
		rPage := r.URL.Query().Get("page")
		rLines := r.URL.Query().Get("lines")

		// check if request page and lines are empty
		if rPage != "" || rLines != "" {
			// TODO : remove after development
			g.CoreAPI.LoggerAPI.Info("Request has body", "body", r.Body)

			params.Page, _ = strconv.Atoi(rPage)
			params.Lines, _ = strconv.Atoi(rLines)
		} else {
			params.Lines = 50
			params.Page = (rows + params.Lines - 1) / params.Lines
		}

		// set starting and end lines based on page and lines
		starting := (params.Lines * (params.Page - 1)) + 1
		if starting < 0 {
			starting = 0
		}
		end := starting + params.Lines
		if end > rows {
			end = rows
		}

		logs, err := logger.ReadLogs(starting, end)
		if err != nil {
			log.Println(err)
		}

		data := map[string]any{
			"logs":        logs,
			"rows":        rows,
			"currentPage": params.Page,
			"perPage":     params.Lines,
		}

		res.Json(w, data, http.StatusOK)
	}
}
