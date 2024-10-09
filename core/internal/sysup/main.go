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
	fromFlareEnv := os.Getenv("FROM_FLARE")
	if strings.ToLower(fromFlareEnv) == "true" {
		fmt.Println("Updater spawned by flare cli")

		// get flare cli pid
		ppid := os.Getppid()
		pproc, err := os.FindProcess(ppid)
		if err != nil {
			log.Println("Error:", err)
			return
		}

		// stop the flare cli, if running
		if isProcRunning(pproc) {
			err := pproc.Kill()
			if err != nil {
				log.Println("Error:", err)
				return
			}

			time.Sleep(1 * time.Second)
		}
	}

	// TODO: ensure core and arch bin files exist
	sdkfs.Exists("")

	// TODO: replace old files with the latest ones

	// TODO: start the new flare CLI server
	newFlareCliCmd := exec.Command("./bin/flare", "server")
	if err := newFlareCliCmd.Start(); err != nil {
		log.Println("Error:", err)
		return
	}

	fmt.Println("Flare System Update Successful")
}

func isProcRunning(proc *os.Process) bool {
	if err := proc.Signal(syscall.Signal(0)); err != nil {
		log.Println("Error:", err)
		return false
	}

	return true
}
