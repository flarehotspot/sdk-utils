package assets

import (
	"core/env"
	jobque "core/internal/utils/job-que"
)

var bundleQue = jobque.NewJobQues()

func Bundle(files []string) (data CacheData, err error) {
	result, err := bundleQue.Exec(func() (interface{}, error) {
		if len(files) == 0 {
			return "", ErrNoAssets
		}

		useCache := env.GoEnv != env.ENV_DEV
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
	return concatFiles(files)
}
