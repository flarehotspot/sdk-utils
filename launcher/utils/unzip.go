package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func Unzip(zippath string, extractTo string) error {
    if err := EmptyDir(extractTo); err != nil {
        return err
    }

	fmt.Printf("Extracting %s to %s\n", zippath, extractTo)
	cmd := exec.Command("unzip", zippath, "-d", extractTo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
