//go:build !dev

package env

import (
	"log"
)

func Print() {
  log.Println(lineComment)
	log.Println("Mode: ", "Production/Staging")
	log.Println("Http Port: ", HttpPort)
  log.Println(lineComment)
}
