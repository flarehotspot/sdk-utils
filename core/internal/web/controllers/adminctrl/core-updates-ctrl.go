package adminctrl

import (
	"core/internal/plugins"
	rpc "core/internal/rpc"
	"core/internal/utils/sysup"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	sdkdownloader "github.com/flarehotspot/go-utils/downloader"
	sdkextract "github.com/flarehotspot/go-utils/extract"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

func FetchUpdatesCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		latestCoreRelease, err := fetchLatestCoreRelease()
		if err != nil {
			log.Println("Error fetching latest core release:", err)
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, latestCoreRelease, http.StatusOK)
	}
}

func fetchLatestCoreRelease() (sysup.Version, error) {
	srv, ctx := rpc.GetCoreMachineTwirpServiceAndCtx()
	latestCoreRelease, err := srv.FetchLatestCoreRelease(ctx, &rpc.FetchLatestCoreReleaseRequest{})
	if err != nil {
		log.Println("Error: ", err)
		return sysup.Version{}, err
	}

	return sysup.Version{
		Major:          int(latestCoreRelease.Major),
		Minor:          int(latestCoreRelease.Minor),
		Patch:          int(latestCoreRelease.Patch),
		CoreZipFileUrl: latestCoreRelease.CoreZipFileUrl,
		ArchBinFileUrl: latestCoreRelease.ArchBinFileUrl,
	}, nil
}

func GetCurrentCoreVersionCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		coreVersion, err := getCurrentCoreVersion()
		if err != nil {
			log.Println("Error:", err)
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, coreVersion, http.StatusOK)
	}
}

// returns the installed core release version
func getCurrentCoreVersion() (sysup.Version, error) {
	// get file content
	var meta struct {
		Name        string `json:"Name"`
		Package     string `json:"Package"`
		Description string `json:"Description"`
		Version     string `json:"Version"`
	}
	pluginJsonFilePath := filepath.Join(sdkpaths.CoreDir, "plugin.json")
	if err := readPluginReleaseData(&meta, pluginJsonFilePath); err != nil {
		log.Printf("Error reading %v: %v", pluginJsonFilePath, err)
		return sysup.Version{}, err
	}

	coreVersion, err := parseVersion(meta.Version)
	if err != nil {
		log.Println("Error parsing plugin version:", err)
		return sysup.Version{}, err
	}

	return coreVersion, nil
}

// reads the plugin.json from the specified path and populates the meta interface
func readPluginReleaseData(meta interface{}, pluginJsonFilePath string) error {
	b, err := os.ReadFile(pluginJsonFilePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, meta); err != nil {
		log.Println("Error unmarshaling the json: ", err)
		return err
	}

	return nil
}

// parses the string versions into a Version struct
func parseVersion(rawVersion string) (sysup.Version, error) {
	prVersion := strings.Split(rawVersion, ".")
	majorVersion, err := strconv.Atoi(prVersion[0])
	if err != nil {
		log.Println("Error parsing major version: ", err)
		return sysup.Version{}, err
	}
	minorVersion, err := strconv.Atoi(prVersion[1])
	if err != nil {
		log.Println("Error parsing minor version: ", err)
		return sysup.Version{}, err
	}
	patchVersion, err := strconv.Atoi(strings.Split(prVersion[2], "-")[0])
	if err != nil {
		log.Println("Error parsing patch version: ", err)
		return sysup.Version{}, err
	}

	return sysup.Version{
		Major: majorVersion,
		Minor: minorVersion,
		Patch: patchVersion,
	}, nil
}

// web controller for /core/download
func DownloadUpdatesCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		var data sysup.Version
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("Error reading the request body:", err)
			return
		}

		// TODO: remove logs
		fmt.Printf("data: %v\n", data)

		stringedVersion := sysup.StringifyVersion(data)
		updatesPath := filepath.Join(sdkpaths.CacheDir, "updates", "core", stringedVersion)
		coreFilesPath := filepath.Join(updatesPath, "core-files")
		archBinFilesPath := filepath.Join(updatesPath, "arch-bin-files")

		// download core files
		err = downloadFiles(data.CoreZipFileUrl, coreFilesPath)
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("Error downloading core files:", err)
			return
		}
		// download arch bin files
		err = downloadFiles(data.ArchBinFileUrl, archBinFilesPath)
		if err != nil {
			log.Println("Error downloading arch bin files:", err)
			res.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO: verify downloaded files

		// return downloaded local file paths
		localUpdateFiles := sysup.UpdateFiles{
			LocalCoreFilesPath:    coreFilesPath,
			LocalArchBinFilesPath: archBinFilesPath,
		}

		res.Json(w, localUpdateFiles, http.StatusOK)
	}
}

// download the files specified from the src and puts on to the dest
func downloadFiles(src string, dest string) error {
	downloader := sdkdownloader.NewDownloader(src, dest)
	err := downloader.Download()
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	return nil
}

// web controller for /core/update
func UpdateCoreCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		// TODO: read version from request body

		var data sysup.UpdateFiles
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("Error reading the request body:", err)
			return
		}

		// TODO: remove logs
		fmt.Printf("data: %v\n", data)
		fmt.Println("data versions: ", data.Major, data.Minor, data.Patch)

		if err := updateCore(data); err != nil {
			log.Println("Error:", err)
			res.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, "updated (testing)", http.StatusOK)
	}
}

// actual implementation of update of the downloaded latest core release,
// by simply extracting the downloaded files and running the updater
func updateCore(localUpdateFiles sysup.UpdateFiles) error {
	// extract downloaded arch bin files
	// extract path convention .tmp/updates/core/<version>/extracted
	// .tmp/updates/core/<version>/core-files
	// .tmp/updates/core/<version>/archbin-files
	extractPath := filepath.Join(sdkpaths.TmpDir, "updates", "core", sysup.StringifyVersion(localUpdateFiles.Version), "extracted")
	fmt.Println("Extracting downloaded latest release to: ", extractPath)

	// TODO: extract downloaded update files to extractPath
	sdkextract.Extract(localUpdateFiles.LocalCoreFilesPath, extractPath)
	sdkextract.Extract(localUpdateFiles.LocalArchBinFilesPath, extractPath)

	if err := sysup.ExecuteUpdater(localUpdateFiles.Version); err != nil {
		log.Println("Error executing updater: ", err)
		return err
	}

	return nil
}
