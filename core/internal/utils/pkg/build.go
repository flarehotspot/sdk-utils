package pkg

import (
	"core/env"
	"core/internal/utils/encdisk"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkgit "github.com/flarehotspot/go-utils/git"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkpkg "github.com/flarehotspot/go-utils/pkg"
	sdkruntime "github.com/flarehotspot/go-utils/runtime"
	sdkstr "github.com/flarehotspot/go-utils/strings"
)

func BuildFromLocal(w io.Writer, def sdkpkg.PluginSrcDef) (sdkpkg.PluginInfo, error) {
	err := InstallPlugin(def.LocalPath, InstallOpts{Def: def, RemoveSrc: false})
	if err != nil {
		return sdkpkg.PluginInfo{}, err
	}

	info, err := GetInfoFromDef(def)
	if err != nil {
		return sdkpkg.PluginInfo{}, err
	}

	// TODO: remove logs
	log.Println("Marking plugins..")
	if err := WriteMetadata(def, info.Package); err != nil {
		return sdkpkg.PluginInfo{}, err
	}

	return info, nil
}

func BuildFromGit(w io.Writer, def sdkpkg.PluginSrcDef) (sdkpkg.PluginInfo, error) {
	dev := sdkstr.Slugify(sdkstr.Rand(16), "_")
	parentpath := RandomPluginPath()
	diskfile := filepath.Join(parentpath, "plugin-clone", "disk", dev)
	mountpath := filepath.Join(parentpath, "plugin-build", "mount", dev)
	clonepath := filepath.Join(mountpath, "clone")
	mnt := encdisk.NewEncrypedDisk(diskfile, mountpath, dev)
	if err := mnt.Mount(); err != nil {
		return sdkpkg.PluginInfo{}, err
	}

	w.Write([]byte("Cloning plugin from git: " + def.GitURL))
	repo := sdkgit.RepoSource{URL: def.GitURL, Ref: def.GitRef}

	if err := sdkgit.Clone(w, repo, clonepath); err != nil {
		return sdkpkg.PluginInfo{}, err
	}

	if err := InstallPlugin(clonepath, InstallOpts{Def: def}); err != nil {
		return sdkpkg.PluginInfo{}, err
	}

	if err := mnt.Unmount(); err != nil {
		return sdkpkg.PluginInfo{}, err
	}

	info, err := GetInfoFromDef(def)
	if err != nil {
		return sdkpkg.PluginInfo{}, err
	}

	if err := WriteMetadata(def, info.Package); err != nil {
		return sdkpkg.PluginInfo{}, err
	}

	return info, nil
}

func BuildPluginSo(pluginSrcDir string, workdir string) error {
	if pluginSrcDir == "" {
		return errors.New("Build plugin error: no plugin source path")
	}

	if workdir == "" {
		return errors.New("Build plugin error: no build path")
	}

	var info sdkpkg.PluginInfo

	pluginSoPath := filepath.Join(pluginSrcDir, "plugin.so")
	if err := sdkfs.ReadJson(filepath.Join(pluginSrcDir, "plugin.json"), &info); err != nil {
		return err
	}

	buildpath := filepath.Join(workdir, "plugins", info.Package)

	if sdkfs.Exists(pluginSoPath) {
		if err := os.Remove(pluginSoPath); err != nil {
			return err
		}
	}

	if err := sdkfs.EmptyDir(workdir); err != nil {
		return err
	}

	if err := sdkfs.EnsureDir(filepath.Join(workdir, "plugins")); err != nil {
		return err
	}

	if err := sdkfs.CopyDir(pluginSrcDir, buildpath, nil); err != nil {
		return err
	}

	if err := sdkfs.CopyDir(filepath.Join(sdkpaths.AppDir, "sdk"), filepath.Join(workdir, "sdk"), nil); err != nil {
		return err
	}

	libs := []string{}
	err := sdkfs.LsDirs("sdk/libs", &libs, false)
	if err != nil {
		return err
	}

	goWork := fmt.Sprintf(`
go %s

use (
    ./sdk/api
    ./sdk/utils
    `, sdkruntime.GO_VERSION)

	for _, lib := range libs {
		goWork += fmt.Sprintf("./sdk/libs/%s\n", filepath.Base(lib))
	}

	goWork += fmt.Sprintf("./plugins/%s\n)", info.Package)
	goworkFile := filepath.Join(workdir, "go.work")
	if err := os.WriteFile(goworkFile, []byte(goWork), sdkfs.PermFile); err != nil {
		return err
	}

	if err := BuildAssets(pluginSrcDir); err != nil {
		return err
	}

	// Don't build templates in development since it is already watched and built by another script.
	if env.GO_ENV != env.ENV_DEV {
		if err := BuildTemplates(buildpath); err != nil {
			return err
		}
	}

	gofile := "main.go"
	outfile := "plugin.so"
	if err := BuildGoPlugin(gofile, outfile, buildpath, nil); err != nil {
		return err
	}

	pluginSoOut := filepath.Join(buildpath, "plugin.so")
	fmt.Printf("Copying '%s' to '%s'\n", sdkpaths.StripRoot(pluginSoOut), sdkpaths.StripRoot(pluginSoPath))
	if err := sdkfs.CopyFile(pluginSoOut, pluginSoPath); err != nil {
		log.Printf("Error copying '%s' to '%s': %+v\n", pluginSoOut, pluginSoPath, err)
		return err
	}

	return nil
}

func BuildGoPlugin(gofile string, outfile string, workdir string, envs []string) error {
	goBin := GoBin()
	extraArgs := []string{"-buildmode=plugin"}

	buildOpts := sdkpkg.GoBuildOpts{
		GoBinPath: goBin,
		WorkDir:   workdir,
		Env:       envs,
		BuildTags: env.BuildTags,
		ExtraArgs: extraArgs,
	}

	if err := sdkpkg.BuildGoModule(gofile, outfile, buildOpts); err != nil {
		return err
	}

	return nil
}

type InstallOpts struct {
	Def       sdkpkg.PluginSrcDef
	RemoveSrc bool
	Encrypt   bool
}
