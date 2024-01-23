package sdkfs

import "os"

func IsFile(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false // Path does not exist or there was an error accessing it
	}

	return !info.IsDir() && (info.Mode()&os.ModeType == 0) // Check if it's not a directory and is a regular file
}
