//go:build dev

package env

import (
	"log"
)

func Print() {
  log.Println(lineComment)
	log.Println("Base API: ", BaseURL)
	log.Println("Mode: ", "Development")
	log.Println("Http Port: ", HttpPort)
  log.Println(lineComment)
}
