/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkfs

import (
	"io"
	"os"
)

func IsFile(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false // Path does not exist or there was an error accessing it
	}

	return !info.IsDir() && (info.Mode()&os.ModeType == 0) // Check if it's not a directory and is a regular file
}

func ReadFile(f string) (string, error) {
	b, err := os.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppendFile appends data to a file named by filename.
// If the file does not exist, AppendFile creates it with permissions perm.
func AppendFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
