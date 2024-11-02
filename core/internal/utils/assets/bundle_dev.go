package assets

import (
	"core/env"
	jobque "core/internal/utils/job-que"
	"errors"
	"path/filepath"
)

var bundleQue = jobque.NewJobQue()

func Bundle(files []string) (data CacheData, err error) {
	result, err := bundleQue.Exec(func() (interface{}, error) {
		if len(files) == 0 {
			return data, ErrNoAssets
		}

		useCache := env.GO_ENV != env.ENV_DEV
		if cache, ok := cacheExists(files); ok && useCache {
			return cache, nil
		}

		concat, err := minifyFiles(files)
		if err != nil {
			return CacheData{}, err
		}

		return writeCache(concat, files)
	})

	return result.(CacheData), err
}

func minifyFiles(files []string) (concat string, err error) {
	// return concatFiles(files)
	f := files[0]
	ext := filepath.Ext(f)

	switch ext {
	case ".js":
		return MinifyJs(files)
	case ".css":
		return concatFiles(files)
	}

	return "", errors.New("Unsupported file type: " + ext)
}
