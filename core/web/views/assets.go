package views

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/flarehotspot/core/utils/assets"
	"github.com/flarehotspot/core/sdk/libs/slug"
	"github.com/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/core/sdk/utils/slices"
)

type AssetBundle struct {
	mu        *sync.Mutex
	ScriptSrc string
	StyleSrc  string
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

var emptyManifest = AssetsManifest{
	PublicPrefix: "",
	Scripts:      []string{},
	Styles:       []string{},
}

func ViewManifest(v string) (va AssetsManifest) {
	f := v + ".assets.json"
	fbytes, err := ioutil.ReadFile(f)
	if err != nil {
		return emptyManifest
	}

	var a AssetsManifest
	err = json.Unmarshal(fbytes, &a)
	if err != nil {
		return emptyManifest
	}

	return a
}

func AssetBundles(view string, opts *BundleExtras) (bundle AssetBundle, err error) {
	m := ViewManifest(view)
	ext := filepath.Ext(view)
	fname := strings.ReplaceAll(strings.ReplaceAll(view, paths.AppDir, ""), "/", "-")
	fname = slug.Make(strings.ReplaceAll(fname, ext, "")) + ext
	if strings.HasPrefix(fname, "-") {
		fname = strings.TrimPrefix(fname, "-")
	}
	dstDir := filepath.Join(PublicPrefix, m.PublicPrefix)
	viewDir := filepath.Dir(view)

	jsSrcs := []string{}
	cssSrcs := []string{}
	dirSrcs := []AssetFolder{}

	if opts != nil {
		if opts.ExtraJS != nil {
			jsSrcs = *opts.ExtraJS
		}
		if opts.ExtraCSS != nil {
			cssSrcs = *opts.ExtraCSS
		}
		if opts.ExtraDirs != nil {
			dirSrcs = *opts.ExtraDirs
		}
	}

	// process js files
	jsFile := filepath.Join(dstDir, "js", fname+".js")
	if len(m.Scripts) > 0 {
		srcs := slices.MapString(m.Scripts, func(s string) string {
			return filepath.Join(viewDir, s)
		})
		jsSrcs = append(jsSrcs, srcs...)
	}

	scriptSrc, err := assets.Bundle(jsFile, jsSrcs)
	if err != nil {
		if !errors.Is(err, assets.ErrNoAssets) {
			return bundle, err
		}
		bundle.ScriptSrc = ""
	} else {
		bundle.ScriptSrc = scriptSrc
	}

	// process css files
	cssFile := filepath.Join(dstDir, "css", fname+".css")
	if len(m.Styles) > 0 {
		srcs := slices.MapString(m.Styles, func(s string) string {
			return filepath.Join(viewDir, s)
		})
		cssSrcs = append(cssSrcs, srcs...)
	}

	styleSrc, err := assets.Bundle(cssFile, cssSrcs)
	if err != nil {
		if !errors.Is(err, assets.ErrNoAssets) {
			return bundle, err
		}
		bundle.StyleSrc = ""
	} else {
		bundle.StyleSrc = styleSrc
	}

	// process folders
	if m.Folders != nil {
		dirSrcs = append(dirSrcs, m.Folders...)
	}

	var wg sync.WaitGroup
	wg.Add(len(dirSrcs))

	for _, f := range dirSrcs {
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
		}(f)
	}

	wg.Wait()
	return bundle, nil
}
