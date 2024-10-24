package updates

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	rpc "core/internal/rpc"
	sdkextract "github.com/flarehotspot/go-utils/extract"
	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdksemver "github.com/flarehotspot/go-utils/semver"
)

const (
	EnvSpawner     = "SPAWNER"
	EnvCoreVersion = "CORE_VERSION"
	EnvValFlare    = "flare"
	EnvValUpdater  = "updater"
)

type CoreReleaseUpdate struct {
	Version        sdksemver.Version
	CoreZipFileUrl string
	ArchBinFileUrl string
}

type UpdateFiles struct {
	LocalCoreFilesPath    string
	LocalArchBinFilesPath string
	CoreReleaseUpdate
}

// Helper function to check if the process is spawned by flare cli
func IsSpawnedFromFlare() bool {
	spawnedFromFlareEnv := os.Getenv(EnvSpawner)
	if strings.ToLower(spawnedFromFlareEnv) == EnvValFlare {
		return true
	}
	return false
}

// Updates the core plugin from a the extracted latest core release
func Update() error {
	// get cwd as the destination for the copying
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("Error getting cwd: ")
	}

	// get latest core release path
	crVersion := strings.ToLower(os.Getenv(EnvCoreVersion))
	fmt.Println("copying and replacing old files..")
	latestCRPath := filepath.Join(".tmp", "updates", "core", crVersion, "extracted")

	// update/copy and replace
	if err := sdkfs.CopyDir(latestCRPath, cwd, &sdkfs.CopyOpts{NoOverride: false, NonRecursive: false}); err != nil {
		log.Println("Error copying/updating the latest core release to flare path:", err)
		return err
	}

	return nil
}

// Executes the copied latest core release
func ExecuteFlare() error {
	// get the latest path
	flarePath := filepath.Join("bin", "flare")
	flareCmd := fmt.Sprintf("./%s", flarePath)

	// run the latest cli with "update" params
	flare := exec.Command(flareCmd, "server")
	flare.Stdout = os.Stdout
	flare.Stderr = os.Stderr

	// set env vars
	flare.Env = append(flare.Env, fmt.Sprintf("%s=%s", EnvSpawner, EnvValUpdater))

	// start
	if err := flare.Start(); err != nil {
		log.Println("Error starting new flare:", err)
		return err
	}

	return nil

}

// Helper function to check if the process id is running
func IsProcRunning(proc *os.Process) bool {
	if err := proc.Signal(syscall.Signal(0)); err != nil {
		log.Println("Error:", err)
		return false
	}

	return true
}

// Checks if all the necessary core release files exist
func EnsureUpdateFiles() error {
	// TODO: ensure core and arch bin files exist
	coreAndArchBinFiles := []string{
		// "",
	}
	for _, f := range coreAndArchBinFiles {
		// TODO: find out proper file path
		if sdkfs.Exists("") {
			fmt.Println(f, " exists")
			continue
		}

		// do not proceed the update
		fmt.Println(f, " does not exist")
		log.Println("Core files not complete.")
		log.Println("Aborting update..")
		return errors.New("updater: error: core files not present")
	}

	return nil
}

// Executes the new flare cli with update params
func ExecuteUpdater(version sdksemver.Version) error {
	// get the latest path
	// convention -> ./tmp/udpates/core/<version>/extracted/
	cliPath := filepath.Join(".tmp", "updates", "core", sdksemver.StringifyVersion(version), "extracted")
	flarePath := filepath.Join(cliPath, "bin", "flare")
	flareCmd := fmt.Sprintf("./%s", flarePath)

	// run the latest cli with "update" params
	updater := exec.Command(flareCmd, "update")
	updater.Stdout = os.Stdout
	updater.Stderr = os.Stderr

	// set env vars
	updater.Env = append(updater.Env, fmt.Sprintf("%s=%s", EnvSpawner, EnvValFlare))
	updater.Env = append(updater.Env, fmt.Sprintf("CORE_VERSION=%s", sdksemver.StringifyVersion(version)))

	// start
	if err := updater.Start(); err != nil {
		log.Println("Error starting updater:", err)
		return err
	}

	return nil
}

// Fetches the latest core release from flare-server
func FetchLatestCoreRelease() (sdksemver.Version, error) {
	srv, ctx := rpc.GetCoreMachineTwirpServiceAndCtx()
	latestCoreRelease, err := srv.FetchLatestCoreRelease(ctx, &rpc.FetchLatestCoreReleaseRequest{})
	if err != nil {
		log.Println("Error: ", err)
		return sdksemver.Version{}, err
	}

	return sdksemver.Version{
		Major: int(latestCoreRelease.Major),
		Minor: int(latestCoreRelease.Minor),
		Patch: int(latestCoreRelease.Patch),
	}, nil
}

// Returns the installed core release version
func GetCurrentCoreVersion() (sdksemver.Version, error) {
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
		return sdksemver.Version{}, err
	}

	coreVersion, err := sdksemver.ParseVersion(meta.Version)
	if err != nil {
		log.Println("Error parsing plugin version:", err)
		return sdksemver.Version{}, err
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

// Extracts and runs the downloaded core release, flare, with update params
func UpdateCore(localUpdateFiles UpdateFiles) error {
	// extract path convention .tmp/updates/core/<version>/extracted
	extractPath := filepath.Join(sdkpaths.TmpDir, "updates", "core", sdksemver.StringifyVersion(localUpdateFiles.Version), "extracted")
	fmt.Println("Extracting downloaded latest release to: ", extractPath)

	sdkextract.Extract(localUpdateFiles.LocalCoreFilesPath, extractPath)
	sdkextract.Extract(localUpdateFiles.LocalArchBinFilesPath, extractPath)

	if err := ExecuteUpdater(localUpdateFiles.Version); err != nil {
		log.Println("Error executing updater: ", err)
		return err
	}

	return nil
}
