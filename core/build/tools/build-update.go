package tools

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func BuildUpdate() {
	fmt.Println("Updating flare system")

	// TODO: stop the flare cli, if running
	flareCliPid := os.Getppid()
	fmt.Printf("pidFlareCli: %v\n", flareCliPid)
	flareCliProcess, err := os.FindProcess(flareCliPid)
	if err != nil {
		log.Println("Error", err)
		return
	}

	err = flareCliProcess.Kill()
	// add some delay

	// TODO: ensure core and arch bin files exist

	// TODO: replace old files with the latest ones

	// TODO: start the new flare CLI with an arg "server"
	newFlareCliCmd := exec.Command("./bin/flare", "server")
	if err := newFlareCliCmd.Start(); err != nil {
		log.Println("Error:", err)
		return
	}

	// Quit
	fmt.Println("Update Successful\nExiting..")
}
