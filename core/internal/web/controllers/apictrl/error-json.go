package apictrl

import (
	"net/http"

	"github.com/flarehotspot/core/internal/web/response"
)

func ErrJson(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	response.Json(w, map[string]any{"error": err.Error()}, 500)
}
