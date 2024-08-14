package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"core/internal/utils/encdisk"
	sdkplugin "sdk/api/plugin"
	sdkfs "sdk/utils/fs"
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
		for _, def := range AllPluginSrc() {
			if def.Src == PluginSrcGit {
				info, err := BuildFromGit(out, def)
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

			if def.Src == PluginSrcSystem || def.Src == PluginSrcLocal {
				info, err := BuildFromLocal(out, def)
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

func installPlugin(src string, info sdkplugin.PluginInfo, opts InstallOpts) error {
	dev := sdkstr.Slugify(info.Package, "_")
	parentpath := RandomPluginPath()
	diskfile := filepath.Join(parentpath, "plugin-clone", "disk", dev)
	mountpath := filepath.Join(parentpath, "plugin-build", "mount", dev)
	mnt := encdisk.NewEncrypedDisk(parentpath, diskfile, mountpath, dev)

	// TODO: remove logs
	log.Println("\n\n---\nMounting..")
	if err := mnt.Mount(); err != nil {
		return err
	}

	// TODO: remove logs
	log.Println("\n\n---\nBuilding plugin..")
	if err := BuildPlugin(src, mountpath); err != nil {
		return err
	}

	installPath := PluginInstallPath(info)
	// TODO: remove logs
	log.Println("\n\n---\nCopying files..")
	for _, file := range PLUGIN_FILES {
		// TODO: remove logs
		log.Println("\n\n---\nCopying file..")
		if err := sdkfs.Copy(filepath.Join(src, file), filepath.Join(installPath, file)); err != nil {
			return err
		}
	}

	if opts.RemoveSrc {
		os.RemoveAll(src)
	}

	// TODO: remove logs
	log.Println("\n\n---\nUnmounting..")
	return mnt.Unmount()
}
