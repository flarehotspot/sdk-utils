package assets

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	stdstr "strings"

	"github.com/flarehotspot/core/utils/crypt"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

type CacheData struct {
	Sum        string `json:"sum"`
	PublicPath string `json:"pub_path"`
	FilePath   string `json:"file_path"`
}

func cacheFile(k string) (f string, err error) {
	data, err := readManifest(k)
	if err != nil {
		return "", err
	}
	return data.FilePath, nil
}

func cacheKey(files []string) string {
	hash, _ := crypt.FastHashFiles(files...)
	return hash
}

func isValidCache(files []string, data *CacheData) bool {
	result, err := crypt.FastHashFiles(files...)
	if err != nil {
		return false
	}
	return result == data.Sum
}

func cacheExists(files []string) (*CacheData, bool) {
	key := cacheKey(files)
	data, err := readManifest(key)

	if err != nil {
		return nil, false
	}

	if _, err := os.Stat(data.FilePath); errors.Is(err, os.ErrNotExist) {
		return nil, false
	}

	if !isValidCache(files, data) {
		return nil, false
	}

	return data, true
}

func writeCache(outfile string, concat string, files []string) (string, error) {
	key := cacheKey(files)
	ext := filepath.Ext(outfile)
	hash, err := crypt.SHA1Files(files)
	if err != nil {
		return "", err
	}

	o := fmt.Sprintf("%s-%s", stdstr.Replace(outfile, ext, "", 1), hash) + ext
	s := filepath.Join(paths.AppDir, o)
	d := filepath.Dir(s)
	sum, err := crypt.FastHashFiles(files...)
	if err != nil {
		return "", err
	}

	cache := CacheData{
		Sum:        sum,
		PublicPath: o,
		FilePath:   filepath.Join(paths.AppDir, o),
	}

	prevFile, cacheErr := cacheFile(key)
	defer func() {
		if cacheErr != nil {
			return
		}
		if prevFile == cache.FilePath {
			return
		}
		err := os.Remove(prevFile)
		if err != nil {
			log.Println(err)
		}
	}()

	err = writeManifest(key, &cache)
	if err != nil {
		return "", nil
	}

	os.MkdirAll(d, os.ModePerm)
	err = ioutil.WriteFile(s, []byte(concat), 0644)
	if err != nil {
		log.Println("Error writing to file: ", s, err)
		return "", err
	}

	return o, nil
}

func filePathComment(f string) string {
	stars := "**************************************************"
	return fmt.Sprintf("\n/%s\nFile: %s\n%s/\n", stars, paths.Strip(f), stars)
}
