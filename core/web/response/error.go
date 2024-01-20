package response

import (
	"encoding/json"
	"net/http"
)

func ErrorJson(w http.ResponseWriter, err string) {
	data, _ := json.Marshal(map[string]any{
		"error": err,
	})
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ErrorHtml(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err))
}
