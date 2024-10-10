package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	sdkfs "github.com/flarehotspot/go-utils/fs"
)

func main() {
	fmt.Println("Updating flare system's core")

	// check if updater was spawned by the flare cli using env
	fromFlareEnv := os.Getenv("RUN_BY_FLARE")
	if strings.ToLower(fromFlareEnv) == "true" {
		fmt.Println("Updater spawned by flare cli")

		// get flare cli pid
		ppid := os.Getppid()
		pproc, err := os.FindProcess(ppid)
		if err != nil {
			log.Println("Error finding parent procces id:", err)
			return
		}

		// stop the flare cli, if running
		if isProcRunning(pproc) {
			err := pproc.Kill()
			if err != nil {
				log.Println("Error finding :", err)
				return
			}

			time.Sleep(1 * time.Second)
		}
	}

	// ensure core and arch bin files exist
	coreAndArchBinFiles := []string{}
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
		return
	}

	// TODO: replace old files with the latest ones
	fmt.Println("Replacing old files..")

	// TODO: start the new flare CLI server
	fmt.Println("Starting the new flare cli..")
	wd, err := os.Getwd()
	if err != nil {
		log.Println("Error getting cwd: ", err)
		return
	}
	fmt.Printf("wd: %v\n", wd)

	newFlareCliCmd := exec.Command("./bin/flare", "server")
	newFlareCliCmd.Stdout = os.Stdout
	newFlareCliCmd.Stderr = os.Stderr
	newFlareCliCmd.Env = append(os.Environ(), "FROM_SYSUP=true")

	if err := newFlareCliCmd.Start(); err != nil {
		log.Println("Error running new flare cli:", err)
		return
	}

	fmt.Println("Flare system updated successfully")
}

func isProcRunning(proc *os.Process) bool {
	if err := proc.Signal(syscall.Signal(0)); err != nil {
		log.Println("Error:", err)
		return false
	}

	return true
}
