package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	sdkfs "github.com/flarehotspot/go-utils/fs"
)

func GoBin() string {
	goCustomPath := os.Getenv("GO_CUSTOM_PATH")
	goCustomBin := filepath.Join(goCustomPath, "bin", "go")
	if sdkfs.Exists(goCustomBin) {
		fmt.Println("Testing go binary: ", goCustomBin)
		testGo := exec.Command(goCustomBin, "env")
		err := testGo.Run()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error running custom go binary, fallback to system...")
			return "go"
		}
		fmt.Println("Using custom go binary: ", goCustomBin)
		return goCustomBin
	}

	fmt.Println("Using system go binary...")
	return "go"
}
