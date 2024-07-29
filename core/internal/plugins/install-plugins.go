package plugins

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"core/build/tools"
	"core/internal/config"
	"core/internal/config/plugincfg"
	"core/internal/utils/encdisk"
	"core/internal/utils/git"
	sdkplugin "sdk/api/plugin"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
	sdkstr "sdk/utils/strings"
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

	go func() {
		for _, def := range config.AllPluginSrc() {
			log.Printf("Plugin Def: %+v\n", def)
			if def.Src == config.PluginSrcGit {
				info, err := buildFromGit(out, def)
				if err != nil {
					log.Println("buildFromGit error:", err)
					out.Msg <- fmt.Sprintf("Error building plugin from git source %s: %s", def.GitURL, err.Error())
				} else {
					out.Msg <- "Installed plugin: " + info.Package
				}
			}

			// if def.Src == config.PluginSrcStore {
			// 	log.Printf("TODO: build from store")
			// }

			if def.Src == config.PluginSrcSystem || def.Src == config.PluginSrcLocal {
				info, err := buildFromLocal(out, def)
				if err != nil {
					out.Msg <- fmt.Sprintf("Error buidling plugin from local path %s: %s", def.LocalPath, err.Error())
				} else {
					out.Msg <- "Installed plugin: " + info.Package
				}
			}

		}

		out.Done <- nil
	}()

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

	// TODO: remove
	time.Sleep(3 * time.Second)

	for _, file := range tools.PLUGIN_FILES {
		if err := sdkfs.Copy(filepath.Join(src, file), filepath.Join(installPath, file)); err != nil {
			return err
		}
	}

	if opts.RemoveSrc {
		os.RemoveAll(src)
	}

	return mnt.Unmount()
}

func buildFromLocal(w io.Writer, src *config.PluginSrcDef) (*sdkplugin.PluginInfo, error) {
	w.Write([]byte("Building plugin from local path: " + src.LocalPath))
	info, err := plugincfg.GetPluginInfo(src.LocalPath)
	if err != nil {
		return nil, err
	}
	err = installPlugin(src.LocalPath, info, InstallOpts{RemoveSrc: false})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func buildFromGit(w io.Writer, src *config.PluginSrcDef) (*sdkplugin.PluginInfo, error) {
	repo := git.RepoSource{URL: src.GitURL, Ref: src.GitRef}
	clonePath := filepath.Join(sdkpaths.TmpDir, "plugins", sdkstr.Rand(16))

	w.Write([]byte("Cloning plugin from git: " + src.GitURL))
	err := git.Clone(w, repo, clonePath)
	if err != nil {
		return nil, err
	}

	info, err := plugincfg.GetPluginInfo(clonePath)
	if err != nil {
		return nil, err
	}

	err = installPlugin(clonePath, info, InstallOpts{})
	if err != nil {
		return nil, err
	}

	os.RemoveAll(clonePath)
	return plugincfg.GetPluginInfo(clonePath)
}
