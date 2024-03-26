package sdklogger

type LoggerApi interface {
	// Logs info msg to console and log file,
	// adds additional info if build in dev
	Debug(msg string) error

	// Logs debug msg to console and log file,
	// adds additional info if build in dev
	Info(msg string) error

	// Logs error msg to console and log file,
	// adds additional info if build in dev
	Error(msg string) error
}
