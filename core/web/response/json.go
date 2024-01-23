package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any, status int) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		ErrorJson(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonBytes)
}
