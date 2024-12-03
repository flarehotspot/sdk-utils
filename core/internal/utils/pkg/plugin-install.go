package pkg

import (
	"core/internal/utils/download"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"

	sdkextract "github.com/flarehotspot/go-utils/extract"
	sdkgit "github.com/flarehotspot/go-utils/git"
	sdkpkg "github.com/flarehotspot/go-utils/pkg"

	"core/internal/utils/encdisk"
	sdkplugin "sdk/api/plugin"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkstr "github.com/flarehotspot/go-utils/strings"
)

type PluginMetadata struct {
	Def sdkpkg.PluginSrcDef
}

type PluginFile struct {
	File     string
	Optional bool
}

var PLuginFiles = []PluginFile{
	{File: "LICENSE.txt", Optional: false},
	{File: "go.mod", Optional: false},
	{File: "plugin.json", Optional: false},
	{File: "plugin.so", Optional: false},
	{File: "resources/assets/dist", Optional: true},
	{File: "resources/migrations", Optional: true},
	{File: "resources/translations", Optional: true},
}

func InstallSrcDef(w io.Writer, def sdkpkg.PluginSrcDef) (info sdkplugin.PluginInfo, err error) {
	switch def.Src {
	case sdkpkg.PluginSrcGit:
		info, err = InstallFromGitSrc(w, def)
	case sdkpkg.PluginSrcLocal, sdkpkg.PluginSrcSystem:
		info, err = InstallFromLocalPath(w, def)
	case sdkpkg.PluginSrcStore:
		info, err = InstallFromPluginStore(w, def)
	default:
		return sdkplugin.PluginInfo{}, errors.New("Invalid plugin source: " + def.Src)
	}

	return info, err
}

func InstallFromLocalPath(w io.Writer, def sdkpkg.PluginSrcDef) (info sdkplugin.PluginInfo, err error) {
	w.Write([]byte("Installing plugin from local path: " + def.LocalPath))

	info, err = GetInfoFromPath(def.LocalPath)
	if err != nil {
		return
	}

	err = InstallPlugin(def.LocalPath, InstallOpts{Def: def, RemoveSrc: false})
	if err != nil {
		return
	}

	return
}

// func InstallFromZipFile(w io.Writer, def config.PluginSrcDef) (info sdkplugin.PluginInfo, err error) {
// 	w.Write([]byte("Installing zipped plugin from local path: " + def.LocalPath))

// 	// prepare path
// 	randomPath := RandomPluginPath()
// 	workPath := filepath.Join(randomPath, "workpath")

// 	// extract compressed plugin release
// 	sdkextract.Extract(def.LocalZipFile, workPath)

// 	if err = os.RemoveAll(filepath.Dir(def.LocalZipFile)); err != nil {
// 		return
// 	}

// 	// gets the plugin release source path
// 	newWorkPath, err := FindPluginSrc(workPath)
// 	if err != nil {
// 		err = errors.New("Unable to find plugin source in: " + workPath)
// 		log.Println("Error: ", err)
// 		return sdkplugin.PluginInfo{}, err
// 	}

// 	// read the plugin.json
// 	info, err = GetInfoFromPath(newWorkPath)
// 	if err != nil {
// 		log.Println("Error getting plugin info: ", err)
// 		return
// 	}

// 	def.LocalPath = filepath.Join(GetInstallPath(info.Package))

// 	if err := InstallPlugin(newWorkPath, InstallOpts{Def: def, RemoveSrc: false}); err != nil {
// 		return sdkplugin.PluginInfo{}, err
// 	}

// 	return info, nil
// }

func InstallFromPluginStore(w io.Writer, def sdkpkg.PluginSrcDef) (sdkplugin.PluginInfo, error) {
	w.Write([]byte("Installing plugin from store: " + def.StorePackage))

	// prepare path
	randomPath := RandomPluginPath()
	diskfile := filepath.Join(randomPath, "disk")
	mountpath := filepath.Join(randomPath, "mount")
	clonePath := filepath.Join(mountpath, "clone", "0") // need extra sub dir
	workPath := filepath.Join(mountpath, "clone", "1")  // need extra sub dir

	// prepare encrypted virtual disk path
	dev := sdkstr.Rand(8)
	mnt := encdisk.NewEncrypedDisk(diskfile, mountpath, dev)
	if err := mnt.Mount(); err != nil {
		log.Println("Error mounting disk: ", err)
		return sdkplugin.PluginInfo{}, err
	}
	defer mnt.Unmount()

	// download plugin release zip file
	log.Println("downloading plugin release: ", def.StoreZipUrl)
	downloader := download.NewDownloader(def.StoreZipUrl, clonePath)
	if err := downloader.Download(); err != nil {
		log.Println("Error: ", err)
		return sdkplugin.PluginInfo{}, err
	}

	// extract compressed plugin release
	sdkextract.Extract(clonePath, workPath)

	// clear StoreZipUrl def
	def.StoreZipUrl = ""

	newWorkPath, err := sdkpkg.FindPluginSrc(workPath)
	if err != nil {
		err = errors.New("Unable to find plugin source in: " + workPath)
		log.Println("Error: ", err)
		return sdkplugin.PluginInfo{}, err
	}
	info, err := GetInfoFromPath(newWorkPath)
	if err != nil {
		log.Println("Error getting plugin info: ", err)
		return sdkplugin.PluginInfo{}, err
	}

	if err := InstallPlugin(newWorkPath, InstallOpts{Def: def, RemoveSrc: false}); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	return info, nil
}

func InstallFromGitSrc(w io.Writer, def sdkpkg.PluginSrcDef) (sdkplugin.PluginInfo, error) {
	log.Println("Installing plugin from git source: " + def.String())
	randomPath := RandomPluginPath()
	diskfile := filepath.Join(randomPath, "disk")
	mountpath := filepath.Join(randomPath, "mount")
	clonePath := filepath.Join(mountpath, "clone", "0") // need extra sub dir

	dev := sdkstr.Rand(8)
	mnt := encdisk.NewEncrypedDisk(diskfile, mountpath, dev)
	if err := mnt.Mount(); err != nil {
		log.Println("Error mounting disk: ", err)
		return sdkplugin.PluginInfo{}, err
	}

	defer mnt.Unmount()

	repo := sdkgit.RepoSource{URL: def.GitURL, Ref: def.GitRef}

	log.Println("Cloning plugin from git: " + def.GitURL)
	if err := sdkgit.Clone(w, repo, clonePath); err != nil {
		log.Println("Error cloning: ", err)
		return sdkplugin.PluginInfo{}, err
	}

	info, err := GetInfoFromPath(clonePath)
	if err != nil {
		log.Println("Error getting plugin info: ", err)
		return sdkplugin.PluginInfo{}, err
	}

	if err := InstallPlugin(clonePath, InstallOpts{Def: def, RemoveSrc: false}); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	return info, nil
}

func InstallPlugin(src string, opts InstallOpts) error {
	log.Println("Installing plugin: ", src)

	var buildpath string

	if opts.Encrypt {
		dev := sdkstr.Rand(8)
		parentPath := RandomPluginPath()
		diskfile := filepath.Join(parentPath, "disk")
		mountpath := filepath.Join(parentPath, "mount")
		buildpath = filepath.Join(mountpath, "build")
		mnt := encdisk.NewEncrypedDisk(diskfile, mountpath, dev)
		if err := mnt.Mount(); err != nil {
			log.Println("Error mounting: ", err)
			return err
		}

		defer mnt.Unmount()
	} else {
		parentpath := filepath.Join(sdkpaths.TmpDir, "b", sdkstr.Rand(16))
		buildpath = filepath.Join(parentpath, "0")
		if err := sdkfs.EmptyDir(buildpath); err != nil {
			return err
		}
		defer os.RemoveAll(parentpath)
	}

	if err := BuildPluginSo(src, buildpath); err != nil {
		log.Println("Error building plugin: ", err)
		return err
	}

	info, err := GetInfoFromPath(src)
	if err != nil {
		log.Println("Error building plugin: ", err)
		return err
	}

	installPath := GetInstallPath(info.Package)
	if err := ValidateInstallPath(installPath); err == nil {
		installPath = GetPendingUpdatePath(info.Package)
	}

	if err := WriteMetadata(opts.Def, info.Package, installPath); err != nil {
		log.Println("Error building plugin: ", err)
		return err
	}

	log.Println("Copying plugin files to: ", installPath)
	for _, f := range PLuginFiles {
		err := sdkfs.Copy(filepath.Join(src, f.File), filepath.Join(installPath, f.File))
		if err != nil && !f.Optional {
			log.Println("Error building plugin: ", err)
			return err
		}
	}

	if opts.RemoveSrc {
		if err := os.RemoveAll(src); err != nil {
			log.Println("Error building plugin: ", err)
			return err
		}
	}

	log.Println("Plugin installed")

	return nil
}

func MarkToRemove(pkg string) error {
	installPath := GetInstallPath(pkg)
	if !sdkfs.Exists(installPath) {
		return errors.New("Plugin not installed: " + pkg)
	}
	uninstallFile := filepath.Join(installPath, "uninstall")
	return os.WriteFile(uninstallFile, []byte(""), sdkfs.PermFile)
}

func IsToBeRemoved(pkg string) bool {
	uninstallFile := filepath.Join(GetInstallPath(pkg), "uninstall")
	return sdkfs.Exists(uninstallFile)
}

func RemovePlugin(pkg string) error {
	meta, err := ReadMetadata(pkg)
	if err != nil {
		return err
	}
	if meta.Def.Src == sdkpkg.PluginSrcLocal || meta.Def.Src == sdkpkg.PluginSrcSystem {
		return os.RemoveAll(meta.Def.LocalPath)
	}
	if err := os.RemoveAll(GetInstallPath(pkg)); err != nil {
		return err
	}
	return nil
}
