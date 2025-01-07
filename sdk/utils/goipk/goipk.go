package sdkgoipk

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	sdkdownloader "github.com/flarehotspot/go-utils/downloader"
	sdkfs "github.com/flarehotspot/go-utils/fs"
	pathsutil "github.com/flarehotspot/go-utils/paths"
)

type GoIpkPaths struct {
	GoIpkPath    string
	GoSrcIpkPath string
}

const (
	GO_REPO_URL = "https://github.com/flarehotspot/golang-releases"
)

func DownloadGoIPK(osArch string, goVersion string) (GoIpkPaths, error) {
	golangIpk := fmt.Sprintf("golang_%s_%s.ipk", goVersion, osArch)
	golangSrcIpk := fmt.Sprintf("golang-src_%s_%s.ipk", goVersion, osArch)
	cacheDir := filepath.Join(pathsutil.CacheDir, "go", goVersion)
	golangIpkCachePath := filepath.Join(cacheDir, golangIpk)
	golangSrcIpkCachePath := filepath.Join(cacheDir, golangSrcIpk)

	if !sdkfs.Exists(golangIpkCachePath) || !sdkfs.Exists(golangSrcIpkCachePath) {
		// cache does not exist
		fmt.Println("Go cache doesn't exist for: go:", goVersion, "os arch:", osArch)
		fmt.Println("Downloading go: OsArch:", osArch, "Go Version:", goVersion)

		if err := sdkfs.EnsureDir(cacheDir); err != nil {
			return GoIpkPaths{}, err
		}

		// https://github.com/flarehotspot/golang-releases/releases/download/v1.21.11/golang_1.21.11_x86_64.ipk
		golangUrl := fmt.Sprintf("%s/releases/download/v%s/golang_%s_%s.ipk", GO_REPO_URL, goVersion, goVersion, osArch)
		golangSrcUrl := fmt.Sprintf("%s/releases/download/v%s/golang-src_%s_%s.ipk", GO_REPO_URL, goVersion, goVersion, osArch)
		fmt.Println("Downloading", golangUrl)

		dl := sdkdownloader.NewDownloader(golangUrl, golangIpkCachePath)
		if err := dl.Download(); err != nil {
			log.Println("Error: Unable to download golang.ipk")
			return GoIpkPaths{}, err
		}

		fmt.Println("Downloading", golangSrcUrl)

		dl = sdkdownloader.NewDownloader(golangSrcUrl, golangSrcIpkCachePath)
		if err := dl.Download(); err != nil {
			log.Println("Error: Unable to download golang-src.ipk")
			return GoIpkPaths{}, err
		}
	}

	goDlSrc := GoIpkPaths{
		GoIpkPath:    golangIpkCachePath,
		GoSrcIpkPath: golangSrcIpkCachePath,
	}

	fmt.Printf("Golang Src: %+v\n", goDlSrc)

	return goDlSrc, nil
}

func InstallGoIPK(ipkPaths GoIpkPaths) error {
	cmd := exec.Command("opkg", "install", ipkPaths.GoIpkPath, ipkPaths.GoSrcIpkPath)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Unable to install go ipk: %s", err)
	}
	return nil
}
