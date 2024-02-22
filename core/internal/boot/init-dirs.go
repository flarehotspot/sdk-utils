package boot

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	sdkfs "github.com/flarehotspot/sdk/utils/fs"
	paths "github.com/flarehotspot/sdk/utils/paths"
)

func InitDirs() {
	dirs := []string{
		paths.ConfigDir,
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
			if err := os.MkdirAll(d, sdkfs.PermDir); err != nil {
				log.Fatal(err)
			}
		}(d)
	}
	wg.Wait()

}
