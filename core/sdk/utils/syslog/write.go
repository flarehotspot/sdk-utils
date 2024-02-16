package sdksyslog

import (
	"os"
	"path"
	"time"

	"github.com/flarehotspot/flarehotspot/core/sdk/utils/paths"
)

func LogNotice(msg string) error {
	return write(TypeNotice, msg)
}

func LogSuccess(msg string) error {
	return write(TypeSuccess, msg)
}

func LogError(msg string) error {
	return write(TypeError, msg)
}

func Log(msg string) error {
	return write(TypeLog, msg)
}

func write(t LogType, msg string) error {
	stamp := time.Now().Format("20060102150405")
	file := path.Join(sdkpaths.LogsDir, string(t)+"-"+stamp+".log")
	err := os.WriteFile(file, []byte(msg), 0644)
	return err
}
