package response

import (
	"net/http"
	json "sdk/libs/go-json"
)

func Json(w http.ResponseWriter, data any, status int) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		ErrorJson(w, err.Error(), 500)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
