package adminctrl

import (
	"core/internal/plugins"
	rpc "core/internal/rpc"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	sdkdownloader "github.com/flarehotspot/go-utils/downloader"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkstr "github.com/flarehotspot/go-utils/strings"
)

type Version struct {
	Major          int
	Minor          int
	Patch          int
	CoreZipFileUrl string
	ArchBinFileUrl string
}

func FetchLatestCoreReleaseCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		latestCoreRelease, err := fetchLatestCoreRelease()
		if err != nil {
			log.Println("Error:", err)
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

func getCurrentCoreVersion() (Version, error) {
	// get file content
	var meta struct {
		Name        string `json:"Name"`
		Package     string `json:"Package"`
		Description string `json:"Description"`
		Version     string `json:"Version"`
	}
	pluginJsonFilePath := filepath.Join(sdkpaths.CoreDir, "plugin.json")
	readPluginReleaseData(&meta, pluginJsonFilePath)

	coreVersion, err := parseVersion(meta.Version)
	if err != nil {
		log.Println("Error:", err)
		return Version{}, err
	}

	return coreVersion, nil
}

func readPluginReleaseData(meta interface{}, pluginJsonFilePath string) error {
	b, err := os.ReadFile(pluginJsonFilePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, meta); err != nil {
		log.Println("Error: ", err)
		return err
	}

	return nil
}

func parseVersion(rawVersion string) (Version, error) {
	prVersion := strings.Split(rawVersion, ".")
	majorVersion, err := strconv.Atoi(prVersion[0])
	if err != nil {
		log.Println("Error: ", err)
		return Version{}, err
	}
	minorVersion, err := strconv.Atoi(prVersion[1])
	if err != nil {
		log.Println("Error: ", err)
		return Version{}, err
	}
	patchVersion, err := strconv.Atoi(strings.Split(prVersion[2], "-")[0])
	if err != nil {
		log.Println("Error: ", err)
		return Version{}, err
	}

	return Version{
		Major: majorVersion,
		Minor: minorVersion,
		Patch: patchVersion,
	}, nil
}

func DownloadUpdatesCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		var data Version
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("Erro:", err)
			return
		}

		coreFilesPath := filepath.Join(sdkpaths.TmpDir, sdkstr.Rand(6))
		err = downloadCoreFiles(data.CoreZipFileUrl, coreFilesPath)
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("Erro:", err)
			return
		}

		archBinFilesPath := filepath.Join(sdkpaths.TmpDir, sdkstr.Rand(6))
		err = downloadArchBin(data.ArchBinFileUrl, archBinFilesPath)
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("Erro:", err)
			return
		}

		type UpdateFiles struct {
			LocalCoreFilesPath    string
			LocalArchBinFilesPath string
		}

		res.Json(w, "downloaded (testing)", http.StatusOK)
	}
}

func downloadCoreFiles(src string, dest string) error {
	downloader := sdkdownloader.NewDownloader(src, dest)
	err := downloader.Download()
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	return nil
}

func downloadArchBin(src string, dest string) error {
	downloader := sdkdownloader.NewDownloader(src, dest)
	err := downloader.Download()
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	return nil
}

func UpateCoreCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		updateCore()

		res.Json(w, "updated (testing)", http.StatusOK)
	}
}

func updateCore() {
	// TODO: update core

	// TODO: extract
}
