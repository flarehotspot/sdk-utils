package response

import (
	"net/http"

	"github.com/goccy/go-json"
)

func ErrorJson(w http.ResponseWriter, err string, status int) {
	data, _ := json.Marshal(map[string]any{
		"error": err,
	})
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
