package plugincfg

// import (
// "os"
// "path/filepath"
// "testing"

// "github.com/flarehotspot/go-utils/fs"
// "github.com/flarehotspot/go-utils/paths"
// "github.com/stretchr/testify/assert"
// )

// func Test_PluginDlr_Download_Git_Src(t *testing.T) {
// src := PluginSrcDef{
// Src:    PluginSrcGit,
// Id:     "flarehotspot/default-theme",
// Branch: "main",
// }

// dler := NewDownloader(&src)
// path, err := dler.Download(nil)
// assert.Nil(t, err)
// assert.True(t, fs.Exists(path))
// assert.Equal(t, filepath.Join(paths.TmpDir, "plugins/downloads", string(src.Id)), path)
// err = os.RemoveAll(paths.TmpDir)
// assert.Nil(t, err)
// }
