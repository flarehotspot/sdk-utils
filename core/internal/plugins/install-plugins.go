package plugins

import (
	"io"
	"os"
	"path/filepath"

	"core/build/tools"
	"core/internal/config"
	"core/internal/config/plugincfg"
	"core/internal/utils/encdisk"
	"core/internal/utils/git"
	"sdk/api/plugin"
	sdkfs "sdk/utils/fs"
	"sdk/utils/paths"
	"sdk/utils/strings"
)

type InstallOpts struct {
	RemoveSrc bool
}

type InstallStatus struct {
	Msg  chan string
	Done chan error
}

func (i *InstallStatus) Write(p []byte) (n int, err error) {
	status := string(p)
	i.Msg <- status
	return len(p), nil
}

func InstallPlugins() *InstallStatus {
	out := &InstallStatus{
		Msg:  make(chan string),
		Done: make(chan error),
	}

	// go func() {
	// 	for _, def := range config.AllPluginSrc() {
	// 		if !isInstalled(def) {
	// 			if def.Src == config.PluginSrcGit {
	// 				info, err := buildFromGit(out, def)
	// 				if err != nil {
	// 					log.Println("buildFromGit error:", err)
	// 					out.Done <- err
	// 					return
	// 				}

	// 				err = plugincfg.WriteCache(def, info)
	// 				if err != nil {
	// 					log.Println("WriteCache error:", err)
	// 					out.Done <- err
	// 					return
	// 				}
	// 			}

	// 			if def.Src == config.PluginSrcStore {
	// 				log.Printf("TODO: build from store")
	// 			}

	// 		}
	// 	}

	// 	out.Done <- nil
	// }()

	return out
}

func installPlugin(src string, info *sdkplugin.PluginInfo, opts InstallOpts) error {
	diskfile := filepath.Join(sdkpaths.TmpDir, "plugin-build", "disk", info.Package)
	buildPath := filepath.Join(sdkpaths.TmpDir, "plugin-build", "mount", info.Package)
	installPath := filepath.Join(sdkpaths.PluginsDir, "installed", info.Package)
	dev := sdkstr.Slugify(info.Package, "_")
	mnt := encdisk.NewEncrypedDisk(buildPath, diskfile, dev)
	if err := mnt.Mount(); err != nil {
		return err
	}

	if err := tools.BuildPlugin(src, buildPath); err != nil {
		return err
	}

	if err := sdkfs.CopyDir(buildPath, installPath, nil); err != nil {
		return err
	}

	if opts.RemoveSrc {
		os.RemoveAll(src)
	}

	return mnt.Unmount()
}

func buildFromGit(w io.Writer, src *config.PluginSrcDef) (*sdkplugin.PluginInfo, error) {
	repo := git.RepoSource{URL: src.GitURL, Ref: src.GitRef}
	clonePath := filepath.Join(sdkpaths.TmpDir, "plugins", sdkstr.Rand(16))

	err := git.Clone(w, repo, clonePath)
	if err != nil {
		return nil, err
	}

	info, err := plugincfg.GetPluginInfo(clonePath)
	if err != nil {
		return nil, err
	}

	// if ok := UserLocalVersion(w, info.Package); ok {
	// return plugincfg.GetInstallInfo(info.Package)
	// }

	err = Build(w, clonePath)
	if err != nil {
		return nil, err
	}

	os.RemoveAll(clonePath)
	return plugincfg.GetInstallInfo(info.Package)
}

func isInstalled(def *config.PluginSrcDef) bool {
	// cacheInfo, ok := plugincfg.GetCacheInfo(def)
	// if !ok {
	// 	return false
	// }

	// installInfo, err := plugincfg.GetInstallInfo(cacheInfo.Package)
	// if err != nil {
	// 	return false
	// }

	// return installInfo.Package == cacheInfo.Package
    return false
}
