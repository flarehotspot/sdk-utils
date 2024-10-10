package adminctrl

import (
	"core/internal/plugins"
	rpc "core/internal/rpc"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	sdkdownloader "github.com/flarehotspot/go-utils/downloader"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

type Version struct {
	Major          int
	Minor          int
	Patch          int
	CoreZipFileUrl string
	ArchBinFileUrl string
}

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

func fetchLatestCoreRelease() (Version, error) {
	srv, ctx := rpc.GetCoreMachineTwirpServiceAndCtx()
	latestCoreRelease, err := srv.FetchLatestCoreRelease(ctx, &rpc.FetchLatestCoreReleaseRequest{})
	if err != nil {
		log.Println("Error: ", err)
		return Version{}, err
	}

	return Version{
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
func getCurrentCoreVersion() (Version, error) {
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
		return Version{}, err
	}

	coreVersion, err := parseVersion(meta.Version)
	if err != nil {
		log.Println("Error parsing plugin version:", err)
		return Version{}, err
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
func parseVersion(rawVersion string) (Version, error) {
	prVersion := strings.Split(rawVersion, ".")
	majorVersion, err := strconv.Atoi(prVersion[0])
	if err != nil {
		log.Println("Error parsing major version: ", err)
		return Version{}, err
	}
	minorVersion, err := strconv.Atoi(prVersion[1])
	if err != nil {
		log.Println("Error parsing minor version: ", err)
		return Version{}, err
	}
	patchVersion, err := strconv.Atoi(strings.Split(prVersion[2], "-")[0])
	if err != nil {
		log.Println("Error parsing patch version: ", err)
		return Version{}, err
	}

	return Version{
		Major: majorVersion,
		Minor: minorVersion,
		Patch: patchVersion,
	}, nil
}

// web controller for /core/download
func DownloadUpdatesCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		var data Version
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("Error reading the request body:", err)
			return
		}

		// TODO: remove logs
		fmt.Printf("data: %v\n", data)

		stringedVersion := stringifyVersion(data)
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

		// return the downloaded local paths to clients
		type UpdateFiles struct {
			LocalCoreFilesPath    string
			LocalArchBinFilesPath string
		}

		res.Json(w, "downloaded (testing)", http.StatusOK)
	}
}

// retuns a stringed version
func stringifyVersion(data Version) string {
	return fmt.Sprintf("v%v.%v.%v", data.Major, data.Minor, data.Patch)
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
func UpateCoreCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		if err := updateCore(); err != nil {
			log.Println("Error:", err)
			res.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, "updated (testing)", http.StatusOK)
	}
}

// actual implementation of update of the downloeded latest core release,
// by simply extracting the downloaded files and running the updater
func updateCore() error {
	// TODO: extract downloaded arch bin files

	// TODO: run flare system updater
	sysup := exec.Command("")
	sysup.Stdout = os.Stdout
	sysup.Stderr = os.Stderr

	// add env to sysup to inform sysup that it was spawn from flare cli
	sysup.Env = append(os.Environ(), "RUN_BY_FLARE=true")

	return nil
}
