//go:build !dev

package assets

import jobque "github.com/flarehotspot/core/utils/job-que"

var bundleQue = jobque.NewJobQues()

func Bundle(files []string) (CacheData, error) {
	result, err := bundleQue.Exec(func() (interface{}, error) {
		if len(files) == 0 {
			return "", ErrNoAssets
		}

		if cache, ok := cacheExists(files); ok {
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
