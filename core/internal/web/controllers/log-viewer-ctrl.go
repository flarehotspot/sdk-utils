package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/utils/logger"
)

// Gets the logs based on the requested current page and
// per page queries
func GetLogs(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		var params struct {
			CurrentPage int
			PerPage     int
		}

		// get queries
		rCurrentPage := r.URL.Query().Get("currentPage")
		rPerPage := r.URL.Query().Get("perPage")

		rows := logger.GetLogLines()

		// check if the requested currentPage and perPage are empty
		if rCurrentPage != "" || rPerPage != "" {
			params.CurrentPage, _ = strconv.Atoi(rCurrentPage)
			params.PerPage, _ = strconv.Atoi(rPerPage)
		} else {
			// set default values
			params.PerPage = 50
			params.CurrentPage = (rows + params.PerPage - 1) / params.PerPage // sets the current page to last page
		}

		// set start and end lines based on the
		// currentPage and perPage query
		start := (params.PerPage * (params.CurrentPage - 1))
		if start < 0 {
			start = 0
		}
		end := start + params.PerPage - 1
		if end > rows {
			end = rows
		}

		// read logs
		logs, err := logger.ReadLogs(start, end)
		if err != nil {
			log.Println(err)
		}

		data := map[string]any{
			"logs":        logs,
			"rows":        rows,
			"currentPage": params.CurrentPage,
			"perPage":     params.PerPage,
		}

		res.Json(w, data, http.StatusOK)
	}
}
