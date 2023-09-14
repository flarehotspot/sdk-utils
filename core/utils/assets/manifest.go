package assets

import (
	"encoding/json"
	"github.com/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/core/sdk/utils/paths"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var manifestDir = filepath.Join(paths.CacheDir, "assets/manifest")

func init() {
	if !fs.Exists(manifestDir) {
		os.MkdirAll(manifestDir, os.ModePerm)
	}
}

func manifestFile(f string) string {
	return filepath.Join(manifestDir, f+".json")
}

func readManifest(f string) (*CacheData, error) {
	var cache CacheData
	file := manifestFile(f)
	byteManifest, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteManifest, &cache)
	if err != nil {
		return nil, err
	}

	return &cache, nil
}

func writeManifest(k string, cd *CacheData) (err error) {
	file := manifestFile(k)
	data, err := json.MarshalIndent(cd, "", "  ")
	if err != nil {
		log.Println("Error parsing asset manifest: ", err)
		return err
	}

	err = os.MkdirAll(filepath.Dir(file), os.ModePerm)
	if err != nil {
		log.Println("Error writing to file: ", file)
		return err
	}

	err = ioutil.WriteFile(file, data, 0644)
	if err != nil {
		log.Println("Error writing to file: ", file)
		return err
	}

	return nil
}
