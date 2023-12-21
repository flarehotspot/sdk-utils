package assets

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/core/utils/crypt"
)

var (
	stars = "**************************************************"
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

func writeCache(concat string, files []string) (string, error) {
	ext := filepath.Ext(files[0])
	pubsubdir := strings.Replace(ext, ".", "", 1)
	key := cacheKey(files)
	hash, err := crypt.SHA1Files(files)
	if err != nil {
		return "", err
	}

	pubpath := filepath.Join("/public", pubsubdir, hash+ext)
	abspath := filepath.Join(paths.AppDir, pubpath)
	sum, err := crypt.FastHashFiles(files...)
	if err != nil {
		return "", err
	}

	cache := CacheData{
		Sum:        sum,
		PublicPath: pubpath,
		FilePath:   abspath,
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

	concat = filesComment(files...) + "\n" + concat
	d := filepath.Dir(abspath)
	os.MkdirAll(d, os.ModePerm)
	err = os.WriteFile(abspath, []byte(concat), 0644)
	if err != nil {
		log.Println("Error writing to file: ", abspath, err)
		return "", err
	}

	return pubpath, nil
}

func filePathComment(f string) string {
	return fmt.Sprintf("\n/%s\nFile: %s\n%s/\n", stars, paths.Strip(f), stars)
}

func filesComment(files ...string) string {
	comment := "/"
	comment += stars
	comment += "\nFiles:\n"
	for _, f := range files {
		comment += fmt.Sprintf("%s\n", paths.Strip(f))
	}
	comment += stars + "/\n"
	return comment
}
