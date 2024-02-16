//go:build !dev

package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/flarehotspot/flarehotspot/core/env"
	"github.com/gorilla/mux"
)

func StartServer(r *mux.Router, forever bool) *http.Server {
	addr := fmt.Sprintf(":%d", env.HttpPort)
	log.Println("Listening on port", addr)
	// log.Fatal(http.ListenAndServe(port, router.RootRouter()))

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if !forever {
		go func() {
			err := srv.ListenAndServe()
			if err != nil {
				log.Printf("Error starting server: %v\n", err)
			}
		}()
	} else {
		log.Fatal(srv.ListenAndServe())
	}

	return srv
}
