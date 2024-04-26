package adminctrl

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
			LogFile     string
		}

		// get queries
		rCurrentPage := r.URL.Query().Get("currentPage")
		rPerPage := r.URL.Query().Get("perPage")

		// get log files
		// logFiles := logger.GetLogFiles()

		// check queries if empty
		if rPerPage != "" {
			params.PerPage, _ = strconv.Atoi(rPerPage)
		} else {
			params.PerPage = 50
		}

		params.LogFile = "app.log"

		// get log rows
		rows := int(logger.CurrLines.Load())
		if params.LogFile != "app.log" {
			rows = logger.GetLogLines(params.LogFile)
		}

		if rCurrentPage != "" {
			params.CurrentPage, _ = strconv.Atoi(rCurrentPage)
		} else {
			params.CurrentPage = (rows + params.PerPage - 1) / params.PerPage
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
		logs, err := logger.ReadLogs(params.LogFile, start, end)
		if err != nil {
			log.Println(err)
			res.Error(w, err.Error(), 500)
			return
		}

		data := map[string]any{
			"logs":           logs,
			"rows":           rows,
			"currentPage":    params.CurrentPage,
			"perPage":        params.PerPage,
			"currentLogFile": params.LogFile,
		}

		res.Json(w, data, http.StatusOK)
	}
}
