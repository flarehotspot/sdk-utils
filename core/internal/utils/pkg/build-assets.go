package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

const (
	AdminSubDir     = "admin"
	PortalSubDir    = "portal"
	AssetsDir       = "resources/assets"
	JsDist          = "resources/assets/dist/js"
	CssDist         = "resources/assets/dist/css"
	ManifestJson    = "resources/assets/manifest.json"
	OutManifestJson = "resources/assets/dist/manifest.json"
)

type Manifest struct {
	GlobalScripts []string            `json:"js_global"`
	GlobalStyles  []string            `json:"css_global"`
	Scripts       map[string][]string `json:"js"`
	Styles        map[string][]string `json:"css"`
}

type PluginManifests struct {
	PortalAssets Manifest `json:"portal"`
	AdminAssets  Manifest `json:"admin"`
}

type CompileResults struct {
	Scripts map[string]string
	Styles  map[string]string
}

type OutputManifest struct {
	AdminAssets  CompileResults `json:"admin"`
	PortalAssets CompileResults `json:"portal"`
}

func BuildAssets(pluginDir string) (err error) {
	manifestPath := filepath.Join(pluginDir, ManifestJson)
	if !sdkfs.Exists(manifestPath) {
		fmt.Println("No manifest file found in: " + pluginDir)
		return nil
	}

	var manifest PluginManifests
	if err = sdkfs.ReadJson(manifestPath, &manifest); err != nil {
		return err
	}

	fmt.Printf("Compiling assets manifest: %+v\n", manifest)

	outManifest := OutputManifest{}

	if results, err := compileManifest(pluginDir, manifest.PortalAssets, PortalSubDir); err != nil {
		return err
	} else {
		outManifest.PortalAssets = results
	}

	if results, err := compileManifest(pluginDir, manifest.AdminAssets, AdminSubDir); err != nil {
		return err
	} else {
		outManifest.AdminAssets = results
	}

	if err = sdkfs.WriteJson(filepath.Join(pluginDir, OutManifestJson), outManifest); err != nil {
		return err
	}

	return nil
}

func compileManifest(pluginDir string, manifest Manifest, subdir string) (results CompileResults, err error) {
	jsDist := filepath.Join(pluginDir, JsDist, subdir)
	cssDist := filepath.Join(pluginDir, CssDist, subdir)
	manifestFile := filepath.Join(pluginDir, ManifestJson)

	if !sdkfs.Exists(manifestFile) {
		return
	}

	if err = sdkfs.EnsureDir(jsDist, cssDist); err != nil {
		return
	}

	results = CompileResults{
		Scripts: make(map[string]string),
		Styles:  make(map[string]string),
	}

	for outname, jsfiles := range manifest.Scripts {
		// TODO: check if scripts is directory and loadd all files inside it
		jsfiles = append(jsfiles, manifest.GlobalScripts...)

		indexjs := ""
		for _, f := range jsfiles {
			f = filepath.Join(pluginDir, AssetsDir, f)
			indexjs += fmt.Sprintf("require('%s');\n", f)
		}

		indexoutjs := filepath.Join(sdkpaths.TmpDir, "assets/build/js", filepath.Base(pluginDir), outname)
		if err = sdkfs.EnsureDir(filepath.Dir(indexoutjs)); err != nil {
			return
		}
		if err = os.WriteFile(indexoutjs, []byte(indexjs), sdkfs.PermFile); err != nil {
			return
		}

		fmt.Printf("Compiling file: %s: %s\n", indexoutjs, indexjs)

		defer os.Remove(indexoutjs)

		// outfile := filepath.Join(jsDist, outname)
		result := api.Build(api.BuildOptions{
			EntryPoints:       []string{indexoutjs},
			Outdir:            filepath.Join(pluginDir, jsDist),
			Outbase:           pluginDir,
			Platform:          api.PlatformBrowser,
			Target:            api.ES5,
			EntryNames:        "[name]-[hash]",
			Sourcemap:         api.SourceMapLinked,
			Bundle:            true,
			AllowOverwrite:    true,
			MinifyWhitespace:  true,
			MinifyIdentifiers: true,
			Write:             false,
		})

		if len(result.Errors) > 0 {
			err = fmt.Errorf("failed to compile js: %v", result.Errors)
			return
		}

		if len(result.Warnings) > 0 {
			err = fmt.Errorf("js warnings: %v", result.Warnings)
			return
		}

		for _, out := range result.OutputFiles {
			if err = sdkfs.EnsureDir(filepath.Dir(out.Path)); err != nil {
				return
			}
			if err = os.WriteFile(out.Path, out.Contents, sdkfs.PermFile); err != nil {
				return
			}
			if filepath.Ext(out.Path) == ".js" {
				outpath := strings.Replace(out.Path, pluginDir, "", 1)
				outpath = strings.TrimPrefix(outpath, "/")
				results.Scripts[outname] = outpath
			}
			fmt.Printf("Outputfile written to: %s\n", out.Path)
		}
	}

	for outname, cssfiles := range manifest.Styles {
		// TODO: check if styles is directory and loadd all files inside it
		cssfiles := append(cssfiles, manifest.GlobalStyles...)
		indexcss := ""
		for _, f := range cssfiles {
			f = filepath.Join(pluginDir, "resources/assets", f)
			indexcss += fmt.Sprintf("\nimport '%s';", f)
		}

		indexOutCss := filepath.Join(sdkpaths.TmpDir, "assets/build", pluginDir, cssDist, outname)
		if err = sdkfs.EnsureDir(filepath.Dir(indexOutCss)); err != nil {
			return
		}
		if err = os.WriteFile(indexOutCss, []byte(indexcss), sdkfs.PermFile); err != nil {
			return
		}

		defer os.Remove(indexOutCss)

		outfile := filepath.Join(cssDist, outname)
		result := api.Build(api.BuildOptions{
			EntryPoints:       []string{indexOutCss},
			Loader:            map[string]api.Loader{"css": api.LoaderCSS},
			Outfile:           outfile,
			AssetNames:        "[name]-[hash]",
			Sourcemap:         api.SourceMapLinked,
			MinifyWhitespace:  true,
			MinifyIdentifiers: true,
			Write:             true,
		})

		if len(result.Errors) > 0 {
			err = fmt.Errorf("failed to compile css: %v", result.Errors)
			return
		}

		if len(result.Warnings) > 0 {
			err = fmt.Errorf("css warnings: %v", result.Warnings)
			return
		}

		for _, out := range result.OutputFiles {
			if err = sdkfs.EnsureDir(filepath.Dir(out.Path)); err != nil {
				return
			}
			if err = os.WriteFile(out.Path, out.Contents, sdkfs.PermFile); err != nil {
				return
			}
			if filepath.Ext(out.Path) == ".css" {
				outpath := strings.Replace(out.Path, pluginDir, "", 1)
				outpath = strings.TrimPrefix(outpath, "/")
				results.Styles[outname] = outpath
			}
			fmt.Printf("Outputfile written to: %s\n", out.Path)
		}
	}

	return
}
