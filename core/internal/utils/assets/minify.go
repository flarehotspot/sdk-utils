package assets

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkstr "github.com/flarehotspot/go-utils/strings"
)

func MinifyJs(files []string) (outdir string, err error) {
	outdir = filepath.Join(sdkpaths.TmpDir, "assets/build/"+sdkstr.Rand(16))
	outfile := filepath.Join(outdir, "out.js")
	indexfile := filepath.Join(outdir, "index.js")

	if err = sdkfs.EnsureDir(filepath.Dir(indexfile)); err != nil {
		return
	}

	indexjs := ""
	for _, f := range files {
		indexjs += fmt.Sprintf("\nrequire('%s');", f)
	}

	fmt.Println("indexjs: \n", indexjs)

	if err := os.WriteFile(indexfile, []byte(indexjs), sdkfs.PermFile); err != nil {
		return "", err
	}

	result := api.Build(api.BuildOptions{
		EntryPoints:       []string{indexfile},
		Target:            api.ES5,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		Bundle:            true,
		Outfile:           outfile,
		Write:             true,
	})

	defer os.Remove(outdir)
	defer os.Remove(indexfile)

	if len(result.Errors) > 0 {
		return "", errors.New(result.Errors[0].Text)
	}

	for _, warning := range result.Warnings {
		fmt.Println("Warning: " + warning.Text)
	}

	b, err := os.ReadFile(outfile)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
