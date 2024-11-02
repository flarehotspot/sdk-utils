package pkg

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkslices "github.com/flarehotspot/go-utils/slices"
)

const (
	AdminSubDir        = "admin"
	PortalSubDir       = "portal"
	AssetsDir          = "resources/assets"
	AdminManifestJson  = "resources/assets/manifest.admin.json"
	PortalManifestJson = "resources/assets/manifest.portal.json"
	OutManifestJson    = "resources/assets/dist/manifest.json"
)

type Manifest map[string][]string

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

	outManifest := OutputManifest{}

	adminManifestPath := filepath.Join(pluginDir, AdminManifestJson)
	if sdkfs.Exists(adminManifestPath) {
		var manifest Manifest
		if err = sdkfs.ReadJson(adminManifestPath, &manifest); err != nil {
			return err
		}
		fmt.Printf("Compiling assets manifest: %+v\n", manifest)

		if results, err := compileManifest(pluginDir, manifest, AdminSubDir); err != nil {
			return err
		} else {
			outManifest.AdminAssets = results
		}
	}

	portalManifestPath := filepath.Join(pluginDir, PortalManifestJson)
	if sdkfs.Exists(portalManifestPath) {
		var manifest Manifest
		if err = sdkfs.ReadJson(portalManifestPath, &manifest); err != nil {
			return err
		}
		fmt.Printf("Compiling assets manifest: %+v\n", manifest)

		if results, err := compileManifest(pluginDir, manifest, PortalSubDir); err != nil {
			return err
		} else {
			outManifest.PortalAssets = results
		}
	}

	outManifestFile := filepath.Join(pluginDir, OutManifestJson)
	if err = sdkfs.EnsureDir(filepath.Dir(outManifestFile)); err != nil {
		return err
	}

	if err = sdkfs.WriteJson(outManifestFile, outManifest); err != nil {
		return err
	}

	return nil
}

func compileManifest(pluginDir string, manifest Manifest, subdir string) (results CompileResults, err error) {
	results = CompileResults{
		Scripts: make(map[string]string),
		Styles:  make(map[string]string),
	}

	for k, files := range manifest {
		// TODO: check if scripts is directory and loadd all files inside it
		ext := filepath.Ext(k)

		supportedExts := []string{".js", ".css"}
		if !sdkslices.Contains(supportedExts, ext) {
			err = errors.New("Unsupported asset format: " + ext)
			return
		}

		distPath := filepath.Join(pluginDir, AssetsDir, "dist", strings.TrimPrefix(ext, "."))

		files = append(files, manifest[k]...)
		outname := strings.TrimSuffix(k, ext)
		indexFile := filepath.Join(distPath, outname+"_index"+ext)

		indexContent := ""
		for _, f := range files {
			f = filepath.Join(pluginDir, AssetsDir, f)
			rel, err := sdkpaths.RelativeFromTo(indexFile, f)
			if err != nil {
				return results, err
			}

			if ext == ".js" {
				indexContent += fmt.Sprintf("require('%s');\n", rel)
			} else if ext == ".css" {
				indexContent += fmt.Sprintf("@import '%s';\n", rel)
			}
		}

		if err = sdkfs.EnsureDir(filepath.Dir(indexFile)); err != nil {
			return
		}
		if err = os.WriteFile(indexFile, []byte(indexContent), sdkfs.PermFile); err != nil {
			return
		}
		defer os.Remove(indexFile)

		fmt.Printf("Compiling file: %s: %s\n", indexFile, indexContent)

		outfile := filepath.Join(distPath, outname+ext)

		var result api.BuildResult

		if ext == ".js" {
			result = api.Build(api.BuildOptions{
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
		} else if ext == ".css" {
			result = api.Build(api.BuildOptions{
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
		}

		if len(result.Errors) > 0 {
			err = fmt.Errorf("failed to compile %s %v", ext, result.Errors)
			return
		}

		if len(result.Warnings) > 0 {
			err = fmt.Errorf("%s warnings: %v", ext, result.Warnings)
			return
		}

		for _, out := range result.OutputFiles {
			if err = sdkfs.EnsureDir(filepath.Dir(out.Path)); err != nil {
				return
			}
			if err = os.WriteFile(out.Path, out.Contents, sdkfs.PermFile); err != nil {
				return
			}
			if filepath.Ext(out.Path) == ext {
				outpath := strings.Replace(out.Path, pluginDir, "", 1)
				outpath = strings.TrimPrefix(outpath, "/")
				results.Scripts[k] = outpath
			}
			fmt.Printf("Outputfile written to: %s\n", out.Path)
		}
	}

	return
}
