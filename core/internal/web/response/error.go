package response

import (
	"net/http"
	"sdk/libs/go-json"
)

func ErrorJson(w http.ResponseWriter, err string, status int) {
	data, _ := json.Marshal(map[string]any{
		"error": err,
	})
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ErrorHtml(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err))
}
