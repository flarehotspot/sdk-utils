package response

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, err error) {
	data, _ := json.Marshal(map[string]any{
		"error": err.Error(),
	})
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(data)
}
