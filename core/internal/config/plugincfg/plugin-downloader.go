package plugincfg

// import (
	// "errors"
	// "fmt"
	// "log"
	// "os"
	// "path/filepath"

	// "github.com/flarehotspot/goutils/cmd"
	// "github.com/flarehotspot/core/sdk/utils/fs"
	// "github.com/flarehotspot/core/sdk/utils/paths"
// )

// const (
	// maxDlTry uint8 = 3
// )

// type PluginDlder struct {
	// src      *PluginSrcDef
	// tryCount uint8
	// doneCh   chan error
// }

// func (self *PluginDlder) Download(err error) (dlpath string, e error) {
	// dlpath = filepath.Join(paths.VendorDir)
	// if self.tryCount < maxDlTry {
		// tmpD := filepath.Join(paths.TmpDir, "plugins/downloads", string(self.src.Id))
		// parentDir := filepath.Dir(tmpD)

		// if !fs.Exists(parentDir) {
			// err = os.MkdirAll(parentDir, os.ModePerm)
			// if err != nil {
				// return "", err
			// }
		// }

		// if fs.Exists(tmpD) {
			// err := os.RemoveAll(tmpD)
			// if err != nil {
				// return "", err
			// }
		// }

		// // repoUrl := "https://" + filepath.Join("github.com/", self.src.Id)
		// repoUrl := fmt.Sprintf("git@github.com:%s", self.src.Id+".git")
		// tag := self.src.Branch
		// err := cmd.Exec(fmt.Sprintf("git clone --depth 1 --branch %s %s %s", tag, repoUrl, tmpD))
		// if err != nil {
			// log.Println(err)
			// self.tryCount += 1
			// return self.Download(err)
		// }
		// if !fs.Exists(tmpD) {
			// return "", errors.New("unable to clone plugin: " + repoUrl)
		// }
		// return tmpD, nil
	// }
	// return "", err
// }

// func NewDownloader(src *PluginSrcDef) *PluginDlder {
	// return &PluginDlder{
		// src:      src,
		// tryCount: 0,
		// doneCh:   make(chan error),
	// }
// }
