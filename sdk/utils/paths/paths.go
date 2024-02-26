/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkpaths

import (
	"os"
	"path/filepath"
	"strings"
)

func rootDir() string {
	if dir := os.Getenv("APPDIR"); dir != "" {
		return dir
	}
	dir, err := os.Getwd()
	if err == nil {
		return dir
	}
	dir = "."
	return dir
}

var (
	AppDir      = rootDir()
	CoreDir     = filepath.Join(AppDir, "core")
	ConfigDir   = filepath.Join(AppDir, "config")
	DefaultsDir = filepath.Join(ConfigDir, ".defaults")
	PluginsDir  = filepath.Join(AppDir, "plugins")
	PublicDir   = filepath.Join(AppDir, "public")
	LogsDir     = filepath.Join(AppDir, "logs")
	SdkDir      = filepath.Join(AppDir, "sdk")
	TmpDir      = filepath.Join(AppDir, ".tmp")
	CacheDir    = filepath.Join(TmpDir, "cache")
)

// Strip removes the project root directory prefix from absolute paths
func Strip(p string) string {
	return strings.Replace(p, AppDir+"/", "", 1)
}
