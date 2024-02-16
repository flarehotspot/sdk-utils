//go:build !dev

package assets

// import (
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"strings"

// 	"github.com/flarehotspot/flarehotspot/core/env"
// 	sdkfs "github.com/flarehotspot/flarehotspot/core/sdk/utils/fs"
// 	sdkpaths "github.com/flarehotspot/flarehotspot/core/sdk/utils/paths"
// 	"github.com/flarehotspot/flarehotspot/core/utils/crypt"
// 	jobque "github.com/flarehotspot/flarehotspot/core/utils/job-que"
// 	minifyv2 "github.com/tdewolff/minify/v2"
// 	"github.com/tdewolff/minify/v2/css"
// 	"github.com/tdewolff/minify/v2/js"
// )

// var bundleQue = jobque.NewJobQues()

// func Bundle(files []string) (CacheData, error) {
// 	result, err := bundleQue.Exec(func() (interface{}, error) {
// 		if len(files) == 0 {
// 			return "", ErrNoAssets
// 		}

// 		useCache := env.GoEnv != env.ENV_DEV
// 		if cache, ok := cacheExists(files); ok && useCache {
// 			return cache, nil
// 		}

// 		concat, err := minifyFiles(files)
// 		if err != nil {
// 			return CacheData{}, err
// 		}

// 		return writeCache(concat, files)
// 	})

// 	return result.(CacheData), err
// }

// func minifyFiles(files []string) (concat string, err error) {
// 	if len(files) == 0 {
// 		return "", nil
// 	}

// 	var sb strings.Builder

// 	mtjs := "application/javascript"
// 	mtcss := "text/css"
// 	m := minifyv2.New()
// 	m.AddFunc(mtcss, css.Minify)
// 	m.AddFunc(mtjs, js.Minify)

// 	allconcat, err := concatFiles(files)
// 	if err != nil {
// 		return "", err
// 	}

// 	ext := filepath.Ext(files[0])
// 	hash, _ := crypt.SHA1Files(files...)
// 	tmpdir := filepath.Join(sdkpaths.TmpDir, "assets-concat")
// 	tmpfile := filepath.Join(tmpdir, hash+ext)

// 	err = sdkfs.EnsureDir(tmpdir)
// 	if err != nil {
// 		log.Println(err)
// 		return "", err
// 	}

// 	err = os.WriteFile(tmpfile, []byte(allconcat), 0644)
// 	if err != nil {
// 		log.Println(err)
// 		return "", err
// 	}

// 	// css or js file
// 	var mimetype string
// 	if ext == ".css" {
// 		mimetype = mtcss
// 	}
// 	if ext == ".js" {
// 		mimetype = mtjs
// 	}
// 	r, err := os.Open(tmpfile)
// 	if err != nil {
// 		log.Println(err)
// 		return "", nil
// 	}
// 	if ext == ".js" {
// 		sb.WriteString(";")
// 	}
// 	if err = m.Minify(mimetype, &sb, r); err != nil {
// 		log.Println("Cannot minify asset file "+tmpfile+":", err)
//         return "", err
// 	}

// 	return sb.String(), nil
// }
