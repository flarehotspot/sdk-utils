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
	// Clean up dist folder
	distPath := filepath.Join(pluginDir, "resources/assets/dist")
	if err = os.RemoveAll(distPath); err != nil {
		return
	}

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
	jsDistPath := filepath.Join(pluginDir, JsDist, subdir)
	cssDistPath := filepath.Join(pluginDir, CssDist, subdir)
	manifestFile := filepath.Join(pluginDir, ManifestJson)

	if !sdkfs.Exists(manifestFile) {
		return
	}

	if err = sdkfs.EnsureDir(jsDistPath, cssDistPath); err != nil {
		return
	}

	results = CompileResults{
		Scripts: make(map[string]string),
		Styles:  make(map[string]string),
	}

	for k, files := range manifest.Scripts {
		// TODO: check if scripts is directory and loadd all files inside it
		files = append(files, manifest.GlobalScripts...)
		outname := strings.TrimSuffix(k, ".js")

		fmt.Println("PluginDir: ", pluginDir)
		indexFile := filepath.Join(jsDistPath, outname+"-index.js")

		indexContent := ""
		for _, f := range files {
			f = filepath.Join(pluginDir, AssetsDir, f)
			rel, err := sdkpaths.RelativeFromTo(indexFile, f)
			if err != nil {
				return results, err
			}
			indexContent += fmt.Sprintf("require('%s');\n", rel)
		}

		if err = sdkfs.EnsureDir(filepath.Dir(indexFile)); err != nil {
			return
		}
		if err = os.WriteFile(indexFile, []byte(indexContent), sdkfs.PermFile); err != nil {
			return
		}
		defer os.Remove(indexFile)

		fmt.Printf("Compiling file: %s: %s\n", indexFile, indexContent)

		outfile := filepath.Join(jsDistPath, outname+".js")
		result := api.Build(api.BuildOptions{
			EntryPoints:       []string{indexFile},
			Outfile:           outfile,
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
				results.Scripts[k] = outpath
			}
			fmt.Printf("Outputfile written to: %s\n", out.Path)
		}
	}

	for k, files := range manifest.Styles {
		// TODO: check if scripts is directory and loadd all files inside it
		files = append(files, manifest.GlobalStyles...)
		outname := strings.TrimSuffix(k, ".css")

		indexFile := filepath.Join(cssDistPath, outname+"-index.css")

		indexContent := ""
		for _, f := range files {
			f = filepath.Join(pluginDir, AssetsDir, f)
			rel, err := sdkpaths.RelativeFromTo(indexFile, f)
			if err != nil {
				return results, err
			}
			indexContent += fmt.Sprintf("@import '%s';\n", rel)
		}

		if err = sdkfs.EnsureDir(filepath.Dir(indexFile)); err != nil {
			return
		}
		if err = os.WriteFile(indexFile, []byte(indexContent), sdkfs.PermFile); err != nil {
			return
		}
		defer os.Remove(indexFile)

		fmt.Printf("Compiling file: %s: %s\n", indexFile, indexContent)

		outfile := filepath.Join(cssDistPath, outname+".css")
		result := api.Build(api.BuildOptions{
			EntryPoints:       []string{indexFile},
			Outfile:           outfile,
			Loader:            map[string]api.Loader{".css": api.LoaderCSS},
			EntryNames:        "[name]-[hash]",
			Sourcemap:         api.SourceMapLinked,
			Bundle:            true,
			AllowOverwrite:    true,
			MinifyWhitespace:  true,
			MinifyIdentifiers: true,
			Write:             false,
		})

		if len(result.Errors) > 0 {
			err = fmt.Errorf("failed to compile CSS: %v", result.Errors)
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
				results.Styles[k] = outpath
			}
			fmt.Printf("Outputfile written to: %s\n", out.Path)
		}
	}

	return
}
