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

	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

type Version struct {
	Major      int
	Minor      int
	Patch      int
	ZipFileUrl string
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
		Major:      int(latestCoreRelease.Major),
		Minor:      int(latestCoreRelease.Minor),
		Patch:      int(latestCoreRelease.Patch),
		ZipFileUrl: latestCoreRelease.ZipFileUrl,
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

func DownloadCoreReleaseCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		downloadCoreRelease()
		downloadArchBin()

		res.Json(w, "", http.StatusOK)
	}
}

func downloadCoreRelease() {
	// TODO: download core release to
}

func downloadArchBin() {

}

func UpateCoreCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()

		updateCore()

		res.Json(w, "", http.StatusOK)
	}
}

func updateCore() {
	// TODO: update core

	// TODO: extract
}
