package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"core/internal/utils/encdisk"
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
	// TODO: update disk file path to randomly select either /etc /var /usr
	diskfileParentPath := filepath.Join(sdkpaths.TmpDir, "plugin-clone", "disk", info.Package)
	// ensure to create the virt disk parent file path exists
	fmt.Printf("creating virtual disk file parent path at: %s", diskfileParentPath)
	if err := os.MkdirAll(diskfileParentPath, 0755); err != nil {
		return err
	}

	diskfile := filepath.Join(diskfileParentPath, info.Package)
	buildPath := filepath.Join(sdkpaths.TmpDir, "plugin-build", "mount", info.Package)
	dev := sdkstr.Slugify(info.Package, "_")
	mnt := encdisk.NewEncrypedDisk(buildPath, diskfile, dev)
	if err := mnt.Mount(); err != nil {
		return err
	}

	if err := BuildPlugin(src, buildPath); err != nil {
		return err
	}

	// TODO: remove
	time.Sleep(3 * time.Second)

	installPath := PluginInstallPath(info)
	for _, file := range PLUGIN_FILES {
		if err := sdkfs.Copy(filepath.Join(src, file), filepath.Join(installPath, file)); err != nil {
			return err
		}
	}

	if opts.RemoveSrc {
		os.RemoveAll(src)
	}

	return mnt.Unmount()
}
