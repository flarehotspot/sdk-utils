package sysup

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	sdkfs "github.com/flarehotspot/go-utils/fs"
)

const (
	EnvSpawner     = "SPAWNER"
	EnvCoreVersion = "CORE_VERSION"
	EnvValFlare    = "flare"
	EnvValUpdater  = "updater"
)

type Version struct {
	Major          int
	Minor          int
	Patch          int
	CoreZipFileUrl string
	ArchBinFileUrl string
}

type UpdateFiles struct {
	LocalCoreFilesPath    string
	LocalArchBinFilesPath string
	Version
}

// helper function to check if the process is spawned by flare cli
func IsSpawnedFromFlare() bool {
	spawnedFromFlareEnv := os.Getenv(EnvSpawner)
	if strings.ToLower(spawnedFromFlareEnv) == EnvValFlare {
		return true
	}
	return false
}

// updates the core plugin from a the extracted latest core release
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

// executes the copied latest core release
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

// checks if the process id is running
func IsProcRunning(proc *os.Process) bool {
	if err := proc.Signal(syscall.Signal(0)); err != nil {
		log.Println("Error:", err)
		return false
	}

	return true
}

// checks if all the necessary core release files exist
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

func ExecuteUpdater(version Version) error {
	// get the latest path
	// convention -> ./tmp/udpates/core/<version>/extracted/
	cliPath := filepath.Join(".tmp", "updates", "core", StringifyVersion(version), "extracted")
	flarePath := filepath.Join(cliPath, "bin", "flare")
	flareCmd := fmt.Sprintf("./%s", flarePath)

	// run the latest cli with "update" params
	updater := exec.Command(flareCmd, "update")
	updater.Stdout = os.Stdout
	updater.Stderr = os.Stderr

	// set env vars
	updater.Env = append(updater.Env, fmt.Sprintf("%s=%s", EnvSpawner, EnvValFlare))
	updater.Env = append(updater.Env, fmt.Sprintf("CORE_VERSION=%s", StringifyVersion(version)))

	// start
	if err := updater.Start(); err != nil {
		log.Println("Error starting updater:", err)
		return err
	}

	return nil
}

func StringifyVersion(data Version) string {
	return fmt.Sprintf("v%v.%v.%v", data.Major, data.Minor, data.Patch)
}
