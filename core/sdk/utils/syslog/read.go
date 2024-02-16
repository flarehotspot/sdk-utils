package sdksyslog

import (
	"github.com/flarehotspot/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/flarehotspot/core/sdk/utils/slices"
)

func ReadNotice() ([]*LogEntry, error) {
	return ReadByType(TypeNotice)
}

func ReadSuccess() ([]*LogEntry, error) {
	return ReadByType(TypeSuccess)
}

func ReadError() ([]*LogEntry, error) {
	return ReadByType(TypeError)
}

func ReadAll() ([]*LogEntry, error) {
	files := []string{}
	if err := sdkfs.LsFiles(sdkpaths.LogsDir, &files, false); err != nil {
		return nil, err
	}

	entries := []*LogEntry{}
	for _, f := range files {
		entries = append(entries, NewLogEntry(f))
	}

	return entries, nil
}

func ReadByType(t LogType) ([]*LogEntry, error) {
	entries, err := ReadAll()
	if err != nil {
		return nil, err
	}

	entries = sdkslices.Filter(entries, func(ent *LogEntry) bool {
		return ent.Type() == t
	})

	return entries, nil
}
