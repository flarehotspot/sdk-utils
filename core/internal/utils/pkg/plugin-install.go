package pkg

import (
	"bytes"
	"core/internal/utils/cmd"
	"core/internal/utils/download"
	"core/internal/utils/encdisk"
	"core/internal/utils/migrate"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	sdkutils "github.com/flarehotspot/sdk-utils"
)

type PluginMetadata struct {
	Def sdkutils.PluginSrcDef
}

func InstallSrcDef(w io.Writer, db *pgxpool.Pool, def sdkutils.PluginSrcDef) (info sdkutils.PluginInfo, err error) {
	switch def.Src {
	case sdkutils.PluginSrcGit:
		info, err = InstallFromGitSrc(w, db, def)
	case sdkutils.PluginSrcLocal, sdkutils.PluginSrcSystem:
		info, err = InstallFromLocalPath(w, db, def)
	case sdkutils.PluginSrcStore:
		info, err = InstallFromPluginStore(w, db, def)
	default:
		return sdkutils.PluginInfo{}, errors.New("Invalid plugin source: " + def.Src)
	}

	return info, err
}

func InstallFromLocalPath(w io.Writer, db *pgxpool.Pool, def sdkutils.PluginSrcDef) (info sdkutils.PluginInfo, err error) {
	w.Write([]byte("Installing plugin from local path: " + def.LocalPath))

	info, err = sdkutils.GetPluginInfoFromPath(def.LocalPath)
	if err != nil {
		return
	}

	err = InstallPlugin(def.LocalPath, db, InstallOpts{Def: def, RemoveSrc: false})
	if err != nil {
		return
	}

	return
}

// func InstallFromZipFile(w io.Writer, def config.PluginSrcDef) (info sdkutils.PluginInfo, err error) {
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
// 		return sdkutils.PluginInfo{}, err
// 	}

// 	// read the plugin.json
// 	info, err = GetInfoFromPath(newWorkPath)
// 	if err != nil {
// 		log.Println("Error getting plugin info: ", err)
// 		return
// 	}

// 	def.LocalPath = filepath.Join(GetInstallPath(info.Package))

// 	if err := InstallPlugin(newWorkPath, InstallOpts{Def: def, RemoveSrc: false}); err != nil {
// 		return sdkutils.PluginInfo{}, err
// 	}

// 	return info, nil
// }

func InstallFromPluginStore(w io.Writer, db *pgxpool.Pool, def sdkutils.PluginSrcDef) (sdkutils.PluginInfo, error) {
	w.Write([]byte("Installing plugin from store: " + def.StorePackage))

	// prepare path
	randomPath := RandomPluginPath()
	diskfile := filepath.Join(randomPath, "disk")
	mountpath := filepath.Join(randomPath, "mount")
	clonePath := filepath.Join(mountpath, "clone", "0") // need extra sub dir
	workPath := filepath.Join(mountpath, "clone", "1")  // need extra sub dir

	// prepare encrypted virtual disk path
	dev := sdkutils.RandomStr(8)
	mnt := encdisk.NewEncrypedDisk(diskfile, mountpath, dev)
	if err := mnt.Mount(); err != nil {
		log.Println("Error mounting disk: ", err)
		return sdkutils.PluginInfo{}, err
	}
	defer mnt.Unmount()

	// download plugin release zip file
	log.Println("downloading plugin release: ", def.StoreZipUrl)
	downloader := download.NewDownloader(def.StoreZipUrl, clonePath)
	if err := downloader.Download(); err != nil {
		log.Println("Error: ", err)
		return sdkutils.PluginInfo{}, err
	}

	// extract compressed plugin release
	sdkutils.FsExtract(clonePath, workPath)

	// clear StoreZipUrl def
	def.StoreZipUrl = ""

	newWorkPath, err := sdkutils.FindPluginSrc(workPath)
	if err != nil {
		err = errors.New("Unable to find plugin source in: " + workPath)
		log.Println("Error: ", err)
		return sdkutils.PluginInfo{}, err
	}
	info, err := sdkutils.GetPluginInfoFromPath(newWorkPath)
	if err != nil {
		log.Println("Error getting plugin info: ", err)
		return sdkutils.PluginInfo{}, err
	}

	if err := InstallPlugin(newWorkPath, db, InstallOpts{Def: def, RemoveSrc: false}); err != nil {
		return sdkutils.PluginInfo{}, err
	}

	return info, nil
}

func InstallFromGitSrc(w io.Writer, db *pgxpool.Pool, def sdkutils.PluginSrcDef) (sdkutils.PluginInfo, error) {
	log.Println("Installing plugin from git source: " + def.String())
	randomPath := RandomPluginPath()
	diskfile := filepath.Join(randomPath, "disk")
	mountpath := filepath.Join(randomPath, "mount")
	clonePath := filepath.Join(mountpath, "clone", "0") // need extra sub dir

	dev := sdkutils.RandomStr(8)
	mnt := encdisk.NewEncrypedDisk(diskfile, mountpath, dev)
	if err := mnt.Mount(); err != nil {
		log.Println("Error mounting disk: ", err)
		return sdkutils.PluginInfo{}, err
	}

	defer mnt.Unmount()

	repo := sdkutils.GitRepoSource{URL: def.GitURL, Ref: def.GitRef}

	log.Println("Cloning plugin from git: " + def.GitURL)
	if err := sdkutils.GitClone(w, repo, clonePath); err != nil {
		log.Println("Error cloning: ", err)
		return sdkutils.PluginInfo{}, err
	}

	info, err := sdkutils.GetPluginInfoFromPath(clonePath)
	if err != nil {
		log.Println("Error getting plugin info: ", err)
		return sdkutils.PluginInfo{}, err
	}

	if err := InstallPlugin(clonePath, db, InstallOpts{Def: def, RemoveSrc: false}); err != nil {
		return sdkutils.PluginInfo{}, err
	}

	return info, nil
}

func InstallPlugin(src string, db *pgxpool.Pool, opts InstallOpts) error {
	log.Println("Installing plugin: ", src)

	var buildpath string

	if opts.Encrypt {
		dev := sdkutils.RandomStr(8)
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
		parentpath := filepath.Join(sdkutils.PathTmpDir, "b", sdkutils.RandomStr(16))
		buildpath = filepath.Join(parentpath, "0")
		if err := sdkutils.FsEmptyDir(buildpath); err != nil {
			return err
		}
		defer os.RemoveAll(parentpath)
	}

	if err := BuildTemplates(src); err != nil {
		log.Println("Error building plugin templates: ", err)
		return err
	}

	if err := BuildPluginSo(src, buildpath); err != nil {
		log.Println("Error building plugin: ", err)
		return err
	}

	info, err := sdkutils.GetPluginInfoFromPath(src)
	if err != nil {
		log.Println("Error building plugin: ", err)
		return err
	}

	installPath := GetInstallPath(info.Package)
	if err := ValidateInstallPath(installPath); err == nil {
		installPath = GetPendingUpdatePath(info.Package)
	}

	if err := InstallSystemPkgs(info.SystemPackages); err != nil {
		return err
	}

	if err := WriteMetadata(opts.Def, info.Package); err != nil {
		log.Println("Error building plugin: ", err)
		return err
	}

	log.Println("Copying plugin files to: ", installPath)
	if err := sdkutils.CopyPluginFiles(src, installPath); err != nil {
		return err
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

func InstallSystemPkgs(packages []string) (err error) {
	if len(packages) == 0 {
		return nil
	}

	if err := cmd.Exec("opkg update", &cmd.ExecOpts{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}); err != nil {
		return err
	}

	toBeInstalled := []string{}
	for _, pkg := range packages {
		installed, err := IsSystemPackageInstalled(pkg)
		if err != nil {
			return err
		}
		if !installed {
			toBeInstalled = append(toBeInstalled, pkg)
		}
	}

	if err := cmd.Exec("opkg install "+strings.Join(toBeInstalled, " "), &cmd.ExecOpts{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}); err != nil {
		return err
	}

	return nil
}

// IsPackageInstalled checks if a package is installed on OpenWrt.
func IsSystemPackageInstalled(opkgPackage string) (bool, error) {
	// Execute the `opkg list-installed` command
	cmd := exec.Command("opkg", "list-installed")
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("failed to execute opkg: %v, output: %s", err, output.String())
	}

	// Check if the package name exists in the output
	installedPackages := output.String()
	return strings.Contains(installedPackages, opkgPackage), nil
}

func RunMigrations(db *pgxpool.Pool, pluginPath string) (err error) {
	info, err := sdkutils.GetPluginInfoFromPath(pluginPath)
	if err != nil {
		return
	}

	name := info.Name
	migdir := filepath.Join(pluginPath, "resources/migrations")
	err = migrate.MigrateUp(db, migdir)
	if err != nil {
		log.Println("Error in plugin migration "+name, ":", err.Error())
		return err
	}

	log.Println("Done migrating plugin:", name)
	return nil
}

func MarkToRemove(pkg string) error {
	installPath := GetInstallPath(pkg)
	if !sdkutils.FsExists(installPath) {
		return errors.New("Plugin not installed: " + pkg)
	}
	uninstallFile := filepath.Join(installPath, "uninstall")
	return os.WriteFile(uninstallFile, []byte(""), sdkutils.PermFile)
}

func IsToBeRemoved(pkg string) bool {
	uninstallFile := filepath.Join(GetInstallPath(pkg), "uninstall")
	return sdkutils.FsExists(uninstallFile)
}

func RemovePlugin(pkg string) error {
	meta, err := ReadMetadata(pkg)
	if err != nil {
		return err
	}
	if meta.Def.Src == sdkutils.PluginSrcLocal || meta.Def.Src == sdkutils.PluginSrcSystem {
		return os.RemoveAll(meta.Def.LocalPath)
	}
	if err := os.RemoveAll(GetInstallPath(pkg)); err != nil {
		return err
	}
	return nil
}
