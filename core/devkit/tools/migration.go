package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	sdkfs "github.com/flarehotspot/sdk/utils/fs"
	sdkstr "github.com/flarehotspot/sdk/utils/strings"
)

func MigrationCreate(pluginDir string, name string) {
	currentTime := time.Now()
	timestamp := currentTime.Format("20060102150405.000000")
    timestamp = strings.Replace(timestamp, ".", "", 1)
	migrationsDir := filepath.Join(pluginDir, "resources/migrations")

	name = sdkstr.Slugify(name)
	migrationUpPath := filepath.Join(migrationsDir, timestamp+"_"+name+".up.sql")
	migrationDownPath := filepath.Join(migrationsDir, timestamp+"_"+name+".down.sql")

	err := sdkfs.EnsureDir(migrationsDir)
	if err != nil {
		panic(err)
	}

    contentUp := "-- Write your sql for up migration here\n"
    contentDown := "-- Write your sql for down migration here\n"

	if err := os.WriteFile(migrationUpPath, []byte(contentUp), sdkfs.PermFile); err != nil {
		panic(err)
	}

	if err := os.WriteFile(migrationDownPath, []byte(contentDown), sdkfs.PermFile); err != nil {
		panic(err)
	}

	fmt.Printf("Migration created at:\n%s\n%s\n", migrationUpPath, migrationDownPath)
}
