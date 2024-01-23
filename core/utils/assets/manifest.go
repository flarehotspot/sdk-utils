package assets

import (
	"encoding/json"
	fs "github.com/flarehotspot/core/sdk/utils/fs"
	paths "github.com/flarehotspot/core/sdk/utils/paths"
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

func readManifest(f string) (CacheData, error) {
	var cache CacheData
	file := manifestFile(f)
	byteManifest, err := os.ReadFile(file)
	if err != nil {
		return CacheData{}, err
	}

	err = json.Unmarshal(byteManifest, &cache)
	if err != nil {
		return CacheData{}, err
	}

	return cache, nil
}

func writeManifest(k string, cd CacheData) (err error) {
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

	err = os.WriteFile(file, data, 0644)
	if err != nil {
		log.Println("Error writing to file: ", file)
		return err
	}

	return nil
}
