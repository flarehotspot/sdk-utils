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
	randomPath := RandomPluginPath()
	diskfile := filepath.Join(randomPath, "disk")
	mountpath := filepath.Join(randomPath, "mount")
	clonePath := filepath.Join(mountpath, "clone")

	dev := sdkstr.Rand(8)
	mnt := encdisk.NewEncrypedDisk(diskfile, mountpath, dev)
	if err := mnt.Mount(); err != nil {
		log.Println("Error mounting disk: ", err)
		return sdkplugin.PluginInfo{}, err
	}

	defer mnt.Unmount()

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

	dev := sdkstr.Rand(8)
	parentpath := RandomPluginPath()
	diskfile := filepath.Join(parentpath, "disk")
	mountpath := filepath.Join(parentpath, "mount")
	buildpath := filepath.Join(mountpath, "build")
	mnt := encdisk.NewEncrypedDisk(diskfile, mountpath, dev)
	if err := mnt.Mount(); err != nil {
		return err
	}

	defer mnt.Unmount()

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

	return nil
}
