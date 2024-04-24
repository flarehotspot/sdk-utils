package plugins

import (
	"github.com/flarehotspot/core/internal/utils/logger"
)

type LoggerApi struct{}

func NewLoggerApi() *LoggerApi {
	return &LoggerApi{}
}

func (l *LoggerApi) Info(title string, body ...any) {
	calldepth := 1
	level := 0

	file, line := logger.GetCallerFileLine(calldepth)

	logger.LogToConsole(file, line, level, title, body...)
	logger.LogToFile(file, line, level, title, body...)
}

func (l *LoggerApi) Debug(title string, body ...any) {
	calldepth := 1
	level := 1

	file, line := logger.GetCallerFileLine(calldepth)

	logger.LogToConsole(file, line, level, title, body...)
	logger.LogToFile(file, line, level, title, body...)
}

func (l *LoggerApi) Error(title string, body ...any) {
	calldepth := 1
	level := 2

	file, line := logger.GetCallerFileLine(calldepth)

	logger.LogToConsole(file, line, level, title, body...)
	logger.LogToFile(file, line, level, title, body...)
}
