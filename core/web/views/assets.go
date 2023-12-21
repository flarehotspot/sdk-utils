package views

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/core/sdk/utils/paths"
	jobque "github.com/flarehotspot/core/utils/job-que"
)

const (
	PublicPrefix  = "/public"
	TagTypeScript = "script"
	TagTypeStyle  = "style"
)

var emptyManifest = AssetsManifest{
	PublicPrefix: "",
	Scripts:      []string{},
	Styles:       []string{},
}

var (
	assetsQue = jobque.NewJobQues()
)

type AssetSources struct {
	Scripts []string
	Styles  []string
}

type AssetFolder struct {
	Src          string `json:"src"`
	PublicPrefix string `json:"prefix"`
}

type AssetsManifest struct {
	PublicPrefix string        `json:"public_prefix"`
	Scripts      []string      `json:"scripts"`
	Styles       []string      `json:"styles"`
	Folders      []AssetFolder `json:"folders"`
}

type BundleExtras struct {
	ExtraJS   *[]string
	ExtraCSS  *[]string
	ExtraDirs *[]AssetFolder
}

func ViewManifest(view string) (ma AssetsManifest) {
	f := view + ".assets.json"
	fbytes, err := os.ReadFile(f)
	if err != nil {
		return emptyManifest
	}

	err = json.Unmarshal(fbytes, &ma)
	if err != nil {
		return emptyManifest
	}

	return ma
}

func ViewAssets(view string) (sources AssetSources) {
	manifest := ViewManifest(view)

	jsSources := []string{}
	for _, s := range manifest.Scripts {
		jsSources = append(jsSources, filepath.Join(filepath.Dir(view), s))
	}

	cssSources := []string{}
	for _, s := range manifest.Styles {
		cssSources = append(cssSources, filepath.Join(filepath.Dir(view), s))
	}

	return AssetSources{
		Scripts: jsSources,
		Styles:  cssSources,
	}
}

func CopyDirsToPublic(view string) error {
	manifest := ViewManifest(view)
	dirs := manifest.Folders

	var wg sync.WaitGroup
	for _, dir := range dirs {
		go func(f AssetFolder) {
			defer wg.Done()
			src := filepath.Join(filepath.Dir(view), f.Src)
			dst := filepath.Join(paths.PublicDir, f.PublicPrefix, filepath.Base(src))
			if !fs.Exists(dst) {
				log.Println("Copy view folders: ", src, "->", dst)
				err := os.MkdirAll(filepath.Dir(dst), os.ModePerm)
				if err != nil {
					log.Println(err)
				}
				err = fs.CopyDir(src, dst, fs.CopyOpts{Override: true, Recursive: true})
				if err != nil {
					log.Println(err)
				}
			}
		}(dir)
	}
	wg.Wait()
	return nil
}
