package plugins

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	timeFormat = "[15:04:05.000]"

	reset = "\033[0m"

	black        = 30
	red          = 31
	green        = 32
	yellow       = 33
	blue         = 34
	magenta      = 35
	cyan         = 36
	lightGray    = 37
	darkGray     = 90
	lightRed     = 91
	lightGreen   = 92
	lightYellow  = 93
	lightBlue    = 94
	lightMagenta = 95
	lightCyan    = 96
	white        = 97
)

type LoggerApi struct {
	debug *log.Logger
	info  *log.Logger
	error *log.Logger
}

func NewLoggerApi() *LoggerApi {
	return &LoggerApi{
		debug: log.New(os.Stdout, colorize(yellow, "flarelog [DEBUG]: "), log.LstdFlags|log.Lshortfile),
		info:  log.New(os.Stdout, colorize(lightBlue, "flarelog [INFO]: "), log.LstdFlags|log.Lshortfile),
		error: log.New(os.Stdout, colorize(lightRed, "flarelog [ERROR]: "), log.LstdFlags|log.Lshortfile),
	}
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return logFile, nil
}

func (self *LoggerApi) Debug(msg string) error {
	// log to console
	self.debug.Output(2, msg)

	// log to log file
	// open file
	file, err := openLogFile("./logs.log")
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	log.Println()

	// write logs to file

	return nil
}

func (self *LoggerApi) Info(msg string) error {
	// log to console
	self.info.Output(2, msg)

	// log to log file
	// open file
	// write logs to file

	return nil
}

func (self *LoggerApi) Error(msg string) error {
	// log to console
	self.error.Output(2, msg)

	// log to log file
	// open file
	// write logs to file

	return nil
}

func colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}
