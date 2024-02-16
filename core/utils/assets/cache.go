package assets

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	sdkfs "github.com/flarehotspot/flarehotspot/core/sdk/utils/fs"
	paths "github.com/flarehotspot/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/flarehotspot/core/utils/crypt"
)

var (
	stars = "**************************************************"
)

type CacheData struct {
	Sum        string `json:"sum"`
	PublicPath string `json:"pub_path"`
	AbsPath    string `json:"file_path"`
}

func cacheFile(k string) (f string, err error) {
	data, err := readManifest(k)
	if err != nil {
		return "", err
	}
	return data.AbsPath, nil
}

func cacheKey(files []string) string {
	hash, _ := crypt.FastHashFiles(files...)
	return hash
}

func isValidCache(files []string, data CacheData) bool {
	result, err := crypt.FastHashFiles(files...)
	if err != nil {
		return false
	}
	return result == data.Sum
}

func cacheExists(files []string) (CacheData, bool) {
	key := cacheKey(files)
	data, err := readManifest(key)

	if err != nil {
		return CacheData{}, false
	}

	if _, err := os.Stat(data.AbsPath); errors.Is(err, os.ErrNotExist) {
		return CacheData{}, false
	}

	if !isValidCache(files, data) {
		return CacheData{}, false
	}

	return data, true
}

func writeCache(concat string, files []string) (data CacheData, err error) {
	ext := filepath.Ext(files[0])
	pubSubDir := strings.Replace(ext, ".", "", 1)
	key := cacheKey(files)
	hash, err := crypt.SHA1Files(files...)
	if err != nil {
		return CacheData{}, err
	}

	pubUriPath := path.Join("/public", pubSubDir, hash+ext)
	absPath := filepath.Join(paths.AppDir, pubUriPath)
	sum, err := crypt.FastHashFiles(files...)
	if err != nil {
		return CacheData{}, err
	}

	cache := CacheData{
		Sum:        sum,
		PublicPath: pubUriPath,
		AbsPath:    absPath,
	}

	defer func() {
		prevFile, cacheErr := cacheFile(key)
		if cacheErr != nil {
			return
		}
		if prevFile == cache.AbsPath {
			return
		}
		err := os.Remove(prevFile)
		if err != nil {
			log.Println(err)
		}
	}()

	err = writeManifest(key, cache)
	if err != nil {
		return CacheData{}, nil
	}

	concat = filesComment(files...) + "\n" + concat
	d := filepath.Dir(absPath)
	os.MkdirAll(d, os.ModePerm)
	err = os.WriteFile(absPath, []byte(concat), 0644)
	if err != nil {
		log.Println("Error writing to file: ", absPath, err)
		return CacheData{}, err
	}

	return cache, nil
}

func filePathComment(f string) string {
	stat, _ := os.Stat(f)
	size := sdkfs.PrettyByteSize(int(stat.Size()))
	return fmt.Sprintf("\n/%s\nFile: %s(%s)\n%s/\n", stars, paths.Strip(f), size, stars)
}

func filesComment(files ...string) string {
	comment := "/"
	comment += stars
	comment += "\nFiles:\n"
	for _, f := range files {
		stat, _ := os.Stat(f)
		size := sdkfs.PrettyByteSize(int(stat.Size()))
		comment += fmt.Sprintf("%s\t\t%s\n", size, paths.Strip(f))
	}
	comment += stars + "/\n"
	return comment
}
