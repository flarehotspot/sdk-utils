package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	sdkfs "github.com/flarehotspot/core/sdk/utils/fs"
)

var (
	GITHUB_TOKEN    = os.Getenv("GITHUB_TOKEN")
	DEFAULT_PLUGINS = []string{
		"flarehotspot/com.flarego.default-theme",
	}
)

func CloneDefaultPlugins(rootDir string) {
	workDir := filepath.Join(rootDir, "plugins")
	sdkfs.EnsureDir(workDir)
	fmt.Println("Cloning system plugins in " + workDir)

	for _, repo := range DEFAULT_PLUGINS {
		GitCloneRepo(repo, workDir)
	}
}

func GitCheckoutMain() {
	dirPaths := []string{"core"}

	var pluginDirs []string
	sdkfs.LsDirs("plugins", &pluginDirs, false)
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

func GitCloneRepo(repo string, workDir string) {
	var gitUrl string
	if GITHUB_TOKEN != "" {
		gitUrl = fmt.Sprintf("https://oauth2:%s@github.com/%s.git", GITHUB_TOKEN, repo)
	} else {
		gitUrl = fmt.Sprintf("git@github.com:%s.git", repo)
	}

    dirname := filepath.Base(repo)
    os.RemoveAll(filepath.Join(workDir, dirname))

	fmt.Println("Cloning " + gitUrl + " in " + workDir)

	cmd := exec.Command("git", "clone", gitUrl)
	cmd.Dir = workDir
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
