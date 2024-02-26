/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdksyslog

import (
	"github.com/flarehotspot/sdk/utils/fs"
	"github.com/flarehotspot/sdk/utils/paths"
	"github.com/flarehotspot/sdk/utils/slices"
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
