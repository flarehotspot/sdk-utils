package plugins

import (
	"github.com/flarehotspot/core/internal/utils/logger"
)

type LoggerApi struct{}

func NewLoggerApi() *LoggerApi {
	// TODO: create log rotation
	// run a go routine for log rotation

	// TODO: create log retention
	// run a go routine for log retention

	return &LoggerApi{}
}

func (l *LoggerApi) Info(title string, body ...any) error {
	calldepth := 1
	level := 0

	file, line := logger.GetCallerFileLine(calldepth)

	logger.LogToConsole(file, line, level, title, body...)
	logger.LogToFile(file, line, level, title, body...)

	return nil
}

func (l *LoggerApi) Debug(title string, body ...any) error {
	calldepth := 1
	level := 1

	file, line := logger.GetCallerFileLine(calldepth)

	logger.LogToConsole(file, line, level, title, body...)
	logger.LogToFile(file, line, level, title, body...)

	return nil
}

func (l *LoggerApi) Error(title string, body ...any) error {
	calldepth := 1
	level := 2

	file, line := logger.GetCallerFileLine(calldepth)

	logger.LogToConsole(file, line, level, title, body...)
	logger.LogToFile(file, line, level, title, body...)

	return nil
}
