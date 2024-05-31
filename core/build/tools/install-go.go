package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
)

func InstallGo(installPath string) {
	if installPath == "" {
		installPath = os.Getenv("GO_CUSTOM_PATH")
	}

	if installPath == "" {
		installPath = filepath.Join("go")
	}

	GOOS := runtime.GOOS
	GOARCH := runtime.GOARCH
	GO_VERSION, err := GoVersion()
	if err != nil {
		panic(err)
	}

	if GoInstallExists(installPath) {
		fmt.Printf("Go version %s already installed to %s\n", GO_VERSION, installPath)
		return
	}

	EXTRACT_PATH := filepath.Join(sdkpaths.CacheDir, "downloads", fmt.Sprintf("go%s-%s-%s", GO_VERSION, GOOS, GOARCH))
	err = downloadAndExtractGo(GOOS, GOARCH, GO_VERSION, EXTRACT_PATH)
	if err != nil {
		panic(err)
	}

	fmt.Println("Installing Go version to: ", installPath)

	err = os.RemoveAll(installPath)
	if err != nil {
		panic(err)
	}

	err = sdkfs.RenameDir(filepath.Join(EXTRACT_PATH, "go"), installPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Go version %s installed to %s\n", GO_VERSION, installPath)
	fmt.Printf("To use the newly installed Go version, run: \n\nexport PATH=%s/bin:$PATH\n", installPath)
}

func GoInstallExists(installPath string) bool {
	fmt.Println("Checking if Go is already installed...")

	GOOS := runtime.GOOS
	GOARCH := runtime.GOARCH
	GO_VERSION, err := GoVersion()
	if err != nil {
		panic(err)
	}

	goBin := filepath.Join(installPath, "bin", "go")
	cmd := exec.Command(goBin, "env")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error checking existing go install: ", err)
		return false
	}

	envValues := map[string]string{}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			envValues[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	goos := strings.Trim(envValues["GOOS"], "\"")
	goarch := strings.Trim(envValues["GOARCH"], "\"")
	goversion := strings.TrimPrefix(strings.Trim(envValues["GOVERSION"], "\""), "go")

	return goos == GOOS && goarch == GOARCH && goversion == GO_VERSION
}

func downloadAndExtractGo(goos, goarch, version, extractPath string) error {
	err := sdkfs.EnsureDir(filepath.Dir(extractPath))
	if err != nil {
		return err
	}

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
