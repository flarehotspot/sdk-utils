package pkg

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"

	"core/internal/utils/encdisk"
	"core/internal/utils/git"
	sdkplugin "sdk/api/plugin"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
	sdkstr "sdk/utils/strings"
)

type PluginFile struct {
	File     string
	Optional bool
}

var PLuginFiles = []PluginFile{
	{
		File:     "plugin.json",
		Optional: false,
	},
	{
		File:     "plugin.so",
		Optional: false,
	},
	{
		File:     "resources",
		Optional: true,
	},
	{
		File:     "go.mod",
		Optional: false,
	},
	{
		File:     "LICENSE.txt",
		Optional: true,
	},
}

func InstallSrcDef(w io.Writer, def PluginSrcDef) (sdkplugin.PluginInfo, error) {
	switch def.Src {
	case PluginSrcGit:
		return InstallGitSrc(w, def)
	case PluginSrcLocal, PluginSrcSystem:
		return InstallLocalPlugin(w, def)
	default:
		return sdkplugin.PluginInfo{}, errors.New("Invalid plugin source: " + def.Src)
	}
}

func InstallLocalPlugin(w io.Writer, def PluginSrcDef) (sdkplugin.PluginInfo, error) {
	w.Write([]byte("Building plugin from local path: " + def.LocalPath))

	info, err := PluginInfo(def.LocalPath)
	if err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	err = InstallPluginPath(def.LocalPath, InstallOpts{RemoveSrc: false})
	if err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	if err := MarkPluginAsInstalled(def, GetInstallPath(info)); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	return info, nil
}

func InstallGitSrc(w io.Writer, def PluginSrcDef) (sdkplugin.PluginInfo, error) {
	rnd := sdkstr.Rand(16)
	diskfile := filepath.Join(sdkpaths.TmpDir, "plugin-clone", "disk", rnd)
	mountpath := filepath.Join(sdkpaths.TmpDir, "plugin-clone", "mount", rnd)
	clonePath := filepath.Join(mountpath, "clone")
	dev := sdkstr.Slugify(rnd, "_")
	mnt := encdisk.NewEncrypedDisk(diskfile, mountpath, dev)
	if err := mnt.Mount(); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	w.Write([]byte("Cloning plugin from git: " + def.GitURL))
	repo := git.RepoSource{URL: def.GitURL, Ref: def.GitRef}

	if err := git.Clone(w, repo, clonePath); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	info, err := PluginInfo(clonePath)
	if err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	if err := InstallPluginPath(clonePath, InstallOpts{RemoveSrc: false}); err != nil {
		os.RemoveAll(clonePath)
		return sdkplugin.PluginInfo{}, err
	}

	if err := mnt.Unmount(); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	if err := MarkPluginAsInstalled(def, GetInstallPath(info)); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	return info, nil
}

func InstallPluginPath(src string, opts InstallOpts) error {
	info, err := PluginInfo(src)
	if err != nil {
		return err
	}

	dev := sdkstr.Slugify(info.Package, "_")
	parentpath := RandomPluginPath()
	diskfile := filepath.Join(parentpath, "plugin-clone", "disk", dev)
	mountpath := filepath.Join(parentpath, "plugin-build", "mount", dev)
	buildpath := filepath.Join(mountpath, "build")
	mnt := encdisk.NewEncrypedDisk(diskfile, mountpath, dev)

	// TODO: remove logs
	log.Println("\n\n---\nMounting..")
	if err := mnt.Mount(); err != nil {
		return err
	}

	// TODO: remove logs
	log.Println("\n\n---\nBuilding plugin..")
	if err := BuildPlugin(src, buildpath); err != nil {
		return err
	}

	installPath := GetInstallPath(info)
	for _, f := range PLuginFiles {
		if err := sdkfs.Copy(filepath.Join(src, f.File), filepath.Join(installPath, f.File)); err != nil && !f.Optional {
			return err
		}
	}

	if opts.RemoveSrc {
		if err := os.RemoveAll(src); err != nil {
			return err
		}
	}

	// TODO: remove logs
	log.Println("\n\n---\nUnmounting..")
	return mnt.Unmount()
}
