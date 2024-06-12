package web

import (
	"fmt"
	"launcher/config"
	"net/http"
)

func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Starting server at port " + config.HttpPort)
	if err := http.ListenAndServe(":"+config.HttpPort, nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
