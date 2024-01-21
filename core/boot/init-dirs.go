package boot

import (
	paths "github.com/flarehotspot/core/sdk/utils/paths"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func InitDirs() {
	dirs := []string{
		paths.ConfigDir,
		paths.VendorDir,
		paths.CacheDir,
		paths.PublicDir,
		filepath.Join(paths.CacheDir, "assets"),
		filepath.Join(paths.ConfigDir, "plugins"),
		filepath.Join(paths.ConfigDir, "accounts"),
	}
	wg := sync.WaitGroup{}
	wg.Add(len(dirs))
	for _, d := range dirs {
		go func(d string) {
			defer wg.Done()
			if err := os.MkdirAll(d, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}(d)
	}
	wg.Wait()

}
