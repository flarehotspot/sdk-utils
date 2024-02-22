package assets

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	fs "github.com/flarehotspot/sdk/utils/fs"
	paths "github.com/flarehotspot/sdk/utils/paths"
	"github.com/flarehotspot/core/internal/utils/crypt"
	jobque "github.com/flarehotspot/core/internal/utils/job-que"
	tmplcache "github.com/flarehotspot/core/internal/utils/flaretmpl"
)

var cacheWithHelpers = sync.Map{}
var bundeWithHelpersQ = jobque.NewJobQues()

type AssetWithData struct {
	File string
	Data any
}

func BundleWithData(entries ...AssetWithData) (CacheData, error) {
	cache, err := bundeWithHelpersQ.Exec(func() (interface{}, error) {
		parsedEntryPaths := []string{}

		tmpDir := filepath.Join(paths.TmpDir, "asset-bundles/entries")
		if err := fs.EnsureDir(tmpDir); err != nil {
			return CacheData{}, err
		}

		files := []string{}
		for _, entry := range entries {
			files = append(files, entry.File)
		}

		hash, err := crypt.FastHashFiles(files...)
		if err != nil {
			return CacheData{}, err
		}

		if cacheData, ok := cacheWithHelpers.Load(hash); ok {
			return cacheData, nil
		}

		for _, entry := range entries {
			var entryCont strings.Builder
			tmpl, err := tmplcache.GetTextTemplate(entry.File)
			if err != nil {
				return CacheData{}, err
			}

			if err := tmpl.Execute(&entryCont, entry.Data); err != nil {
				return CacheData{}, err
			}

			filehash, err := crypt.FastHashFiles(entry.File)
			if err != nil {
				return CacheData{}, err
			}

			filename := filepath.Base(entry.File)
			entryOutputPath := filepath.Join(tmpDir, fmt.Sprintf("%s-%s", filehash, filename))
			err = os.WriteFile(entryOutputPath, []byte(entryCont.String()), os.ModePerm)
			if err != nil {
				return CacheData{}, err
			}

			parsedEntryPaths = append(parsedEntryPaths, entryOutputPath)
		}

		cacheData, err := Bundle(parsedEntryPaths)
		if err != nil {
			return CacheData{}, err
		}

		cacheWithHelpers.Store(hash, cacheData)
		return cacheData, nil
	})

	if err != nil {
		return CacheData{}, err
	}

	return cache.(CacheData), nil
}
