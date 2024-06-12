package main

import (
	"launcher/utils"
	"launcher/web"
	"os"
)

func main() {
    if utils.CoreExists() {
        os.Exit(0)
    } else {
        web.Server()
    }
}
