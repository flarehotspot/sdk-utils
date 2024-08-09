package pkg

import (
	"os"
	"path/filepath"

	"core/internal/utils/encdisk"
	sdkplugin "sdk/api/plugin"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
	sdkstr "sdk/utils/strings"
)

type InstallOpts struct {
	RemoveSrc bool
}

func installPlugin(src string, info sdkplugin.PluginInfo, opts InstallOpts) error {
	diskfile := filepath.Join(sdkpaths.TmpDir, "plugin-build", "disk", info.Package)
	buildPath := filepath.Join(sdkpaths.TmpDir, "plugin-build", "mount", info.Package)
	dev := sdkstr.Slugify(info.Package, "_")
	mnt := encdisk.NewEncrypedDisk(buildPath, diskfile, dev)
	if err := mnt.Mount(); err != nil {
		return err
	}

	if err := BuildPlugin(src, buildPath); err != nil {
		return err
	}

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
