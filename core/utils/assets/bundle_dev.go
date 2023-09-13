//go:build dev

package assets

import jobque "github.com/flarehotspot/core/utils/job-que"

var bundleQue = jobque.NewJobQues()

func Bundle(outfile string, files []string) (string, error) {
	result, err := bundleQue.Exec(func() (interface{}, error) {
		if len(files) == 0 {
			return "", ErrNoAssets
		}

		if cache, ok := cacheExists(files); ok {
			return cache.PublicPath, nil
		}

		concat, err := concatFiles(files)
		if err != nil {
			return "", err
		}

		return writeCache(outfile, concat, files)
	})

	return result.(string), err
}
