package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, data any, status int) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("render.Json() error: %v\n", err)
		w.WriteHeader(status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonBytes)
}
