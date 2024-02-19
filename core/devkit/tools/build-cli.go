package tools

import (
	"fmt"
	"os"
)

func BuildFlareCLI() {
	fmt.Println("Building flare CLI...")
	// goBin := GoBin()
	// sdkfs.EnsureDir("bin")
	// buildArgs := BuildArgs()

	// os.RemoveAll("bin")

	// cliPath := "bin/flare"
	// if runtime.GOOS == "windows" {
	// 	cliPath += ".exe"
	// }

	// buildCmd := []string{"build"}
	// buildCmd = append(buildCmd, buildArgs...)
	// buildCmd = append(buildCmd, "-o", cliPath, "core/devkit/cli/flare.go")

	//    cliFile := "core/devkit/cli/flare.go"
	// cliPath := "bin/flare"
	// cmd := exec.Command(goBin, buildCmd...)

	// err := cmd.Run()

	cliFile := "core/devkit/cli/flare.go"
	cliPath := "bin/flare"
	workDir, _ := os.Getwd()
	err := BuildGoModule(cliFile, cliPath, workDir)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Flare CLI built at: %s\n", cliPath)
}
