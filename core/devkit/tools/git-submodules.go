package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	sdkfs "github.com/flarehotspot/core/sdk/utils/fs"
)

var (
	GITHUB_TOKEN   = os.Getenv("GITHUB_TOKEN")
	SYSTEM_PLUGINS = []string{
		"flarehotspot/com.flarego.default-theme",
	}
)

func GitCloneSystemPlugins(rootDir string) {
	workDir := filepath.Join(rootDir, "system")
	sdkfs.EnsureDir(workDir)
	fmt.Println("Cloning system plugins in " + workDir)

	for _, s := range SYSTEM_PLUGINS {
		var gitUrl string
		if GITHUB_TOKEN != "" {
			gitUrl = fmt.Sprintf("https://oauth2:%s@github.com/%s.git", GITHUB_TOKEN, s)
		} else {
			gitUrl = fmt.Sprintf("git@github.com:%s.git", s)
		}

		fmt.Println("Cloning " + gitUrl)

		cmd := exec.Command("git", "clone", gitUrl)
		cmd.Dir = workDir
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

func GitCheckoutMain() {
	dirPaths := []string{"core", "sdk"}

	var systemPaths []string
	sdkfs.LsDirs("system", &systemPaths, false)

	var pluginDirs []string
	sdkfs.LsDirs("plugins", &pluginDirs, false)

	dirPaths = append(dirPaths, systemPaths...)
	dirPaths = append(dirPaths, pluginDirs...)

	for _, dirPath := range dirPaths {
		fmt.Printf("Checking out main branch for %s...\n", dirPath)
		cmd := exec.Command("git", "checkout", "main")
		cmd.Dir = dirPath
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}
