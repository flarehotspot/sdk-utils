package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	sdkfs "github.com/flarehotspot/sdk/utils/fs"
	sdkpaths "github.com/flarehotspot/sdk/utils/paths"
)

func InstallExactGoVersion() {
	GO_CUSTOM_PATH := os.Getenv("GO_CUSTOM_PATH")
	if GO_CUSTOM_PATH == "" {
		GO_CUSTOM_PATH = filepath.Join("go")
	}

	GOOS := runtime.GOOS
	GOARCH := runtime.GOARCH
	GO_VERSION, err := GoVersion()
	if err != nil {
		panic(err)
	}
	EXTRACT_PATH := filepath.Join(sdkpaths.CacheDir, "downloads", fmt.Sprintf("go%s-%s-%s", GO_VERSION, GOOS, GOARCH))
	err = downloadAndExtractGo(GOOS, GOARCH, GO_VERSION, EXTRACT_PATH)
	if err != nil {
		panic(err)
	}

	fmt.Println("Installing Go version to: ", GO_CUSTOM_PATH)

	err = os.RemoveAll(GO_CUSTOM_PATH)
	if err != nil {
		panic(err)
	}

	err = os.Rename(filepath.Join(EXTRACT_PATH, "go"), GO_CUSTOM_PATH)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Go version %s installed to %s\n", GO_VERSION, GO_CUSTOM_PATH)
	fmt.Printf("To use the newly installed Go version, run: \n\nexport PATH=%s/bin:$PATH\n", GO_CUSTOM_PATH)
}

func downloadAndExtractGo(goos, goarch, version, extractPath string) error {
	// Download the tar file
	fmt.Printf("Downloading Go version %s for %s-%s\n", version, goos, goarch)
	url := fmt.Sprintf("https://golang.org/dl/go%s.%s-%s.tar.gz", version, goos, goarch)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create a temporary file to store the downloaded tar.gz file
	tmpFile, err := os.CreateTemp("", "golang*.tar.gz")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Copy the downloaded content to the temporary file
	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return err
	}

	// Extract the tar.gz file to the specified path
	fmt.Println("Extracting Go to: ", extractPath)
	err = extractTarGz(tmpFile.Name(), extractPath)
	if err != nil {
		return err
	}

	return nil
}

func extractTarGz(srcFile, destPath string) error {
	sdkfs.EmptyDir(destPath)
	cmd := exec.Command("tar", "xzf", srcFile, "-C", destPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
