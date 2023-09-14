package log

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "app: ", log.LstdFlags|log.Lshortfile)
}

func Println(args ...any) {
	logger.Println(args...)
}

func Printf(format string, args ...any) {
	logger.Printf(format, args...)
}
