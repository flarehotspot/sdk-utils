package plugincfg

// import (
// "log"
// "path/filepath"
// "sync/atomic"

// "sdk/libs/slug"
// "github.com/flarehotspot/go-utils/fs"
// "github.com/flarehotspot/go-utils/paths"
// )

// type DownloadError struct {
// errors []error
// }

// func (self *DownloadError) Error() string {
// msg := ""
// for _, err := range self.errors {
// msg += ", " + err.Error()
// }
// return msg
// }

// func DownloadPlugins() error {
// sources := AllPluginSrc()
// errCh := make(chan error)
// errors := []error{}
// var count uint32 = 0

// for _, src := range sources {
// go func(src *PluginSrcDef) {
// err := DlPlugin(src)
// errCh <- err
// atomic.AddUint32(&count, 1)
// if atomic.LoadUint32(&count) >= uint32(len(sources)) {
// close(errCh)
// }
// }(src)
// }

// go func() {
// for err := range errCh {
// if err != nil {
// log.Println(err)
// errors = append(errors, err)
// continue
// }
// }
// }()

// if len(errors) > 0 {
// return &DownloadError{errors: errors}
// }
// return nil
// }

// func DlPlugin(src *PluginSrcDef) (err error) {
// downloader := NewDownloader(src)
// dlpath, err := downloader.Download(nil)
// if err != nil {
// log.Println(err)
// log.Println("unable to clone repo: ", err)
// return err
// }

// info, err := GetPluginInfo(dlpath)
// if err != nil {
// return err
// }

// override := false
// pluginDir := filepath.Join(paths.VendorDir, slug.Make(info.Name))
// if fs.Exists(pluginDir) {
// oldInfo, err := GetPluginInfo(pluginDir)
// if err != nil {
// return err
// }
// if oldInfo.Version != info.Version {
// override = true
// }
// }
// err = fs.CopyDir(dlpath, pluginDir, fs.CopyOpts{Override: override})
// return err
// }
