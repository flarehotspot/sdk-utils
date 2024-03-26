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
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

const (
	Logdir      = "./flarelogs"
	Logfilename = "app.log"
	Infoprefix  = "[INFO] "
	Debugprefix = "[DEBUG] "
	Errorprefix = "[ERROR] "
)

func NewLoggerApi() *LoggerApi {
	// run a go routine for log rotation

	// run a go routine for log retention

	return &LoggerApi{
		debugLogger: log.New(os.Stdout, "", log.LstdFlags|log.Llongfile),
		infoLogger:  log.New(os.Stdout, "", log.LstdFlags|log.Llongfile),
		errorLogger: log.New(os.Stdout, "", log.LstdFlags|log.Llongfile),
	}
}

func openLogFile(path string) (*os.File, error) {
	// ensure log file directory exists
	err := os.MkdirAll(Logdir, 0700)
	if err != nil {
		log.Fatal("Error creating log directory: ", err)
		return nil, err
	}

	// opening/creating file
	logFile, err := os.OpenFile(Logdir+"/"+path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error creating log file: ", err)
		return nil, err
	}

	return logFile, nil
}

func logToConsole(l *log.Logger, msg string) error {
	// set output to console
	l.SetOutput(os.Stdout)

	// log to console
	l.SetPrefix(colorize(yellow, Infoprefix))
	err := l.Output(2, msg)
	if err != nil {
		log.Fatal("Error logging to console", err)
		return err
	}

	return nil
}

func logToFile(l *log.Logger, msg string) error {
	// open file
	file, err := openLogFile(Logfilename)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
		return err
	}

	// set output to file
	l.SetOutput(file)

	// log to file
	l.SetPrefix(Debugprefix)
	err = l.Output(2, msg)
	if err != nil {
		log.Fatal("Error logging to file", err)
		return err
	}

	return nil
}

func (self *LoggerApi) Debug(msg string) error {
	// log to file
	err := logToFile(self.debugLogger, msg)
	if err != nil {
		log.Fatal("Error logging to file: ", err)
		return err
	}

	// log to console
	err = logToConsole(self.debugLogger, msg)
	if err != nil {
		log.Fatal("Error logging to console: ", err)
		return err
	}

	return nil
}

func (self *LoggerApi) Info(msg string) error {
	// log to file
	err := logToFile(self.infoLogger, msg)
	if err != nil {
		log.Fatal("Error logging to file: ", err)
		return err
	}

	// log to console
	err = logToConsole(self.infoLogger, msg)
	if err != nil {
		log.Fatal("Error logging to console: ", err)
		return err
	}

	return nil
}

func (self *LoggerApi) Error(msg string) error {
	// log to file
	err := logToFile(self.errorLogger, msg)
	if err != nil {
		log.Fatal("Error logging to file: ", err)
		return err
	}

	// log to console
	err = logToConsole(self.errorLogger, msg)
	if err != nil {
		log.Fatal("Error logging to console: ", err)
		return err
	}

	return nil
}

func colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

// TODO: read logs
// func readLogs() {}
