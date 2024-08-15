//go:build dev

package env

import (
	"log"
)

func Print() {
	log.Println(lineComment)
	log.Println("Base API: ", BASE_URL)
	log.Println("Mode: ", "Development")
	log.Println("Http Port: ", HTTP_PORT)
	log.Println(lineComment)
}
