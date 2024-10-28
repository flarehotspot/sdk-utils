package assets

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"core/internal/utils/crypt"
	"core/internal/utils/flaretmpl"
	jobque "core/internal/utils/job-que"

	fs "github.com/flarehotspot/go-utils/fs"
	paths "github.com/flarehotspot/go-utils/paths"
)

var cacheWithHelpers = sync.Map{}
var bundleWithHelpersQue = jobque.NewJobQue()

type AssetWithData struct {
	File string
	Data any
}

func BundleWithData(entries ...AssetWithData) (CacheData, error) {
	cache, err := bundleWithHelpersQue.Exec(func() (interface{}, error) {
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
			tmpl, err := flaretmpl.GetTextTemplate(entry.File)
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
