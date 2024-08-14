package pkg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"core/internal/utils/encdisk"
	"core/internal/utils/git"
	sdkplugin "sdk/api/plugin"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
	sdkruntime "sdk/utils/runtime"
	sdkstr "sdk/utils/strings"
)

func BuildFromLocal(w io.Writer, def PluginSrcDef) (sdkplugin.PluginInfo, error) {
	if ok, path := IsPluginInstalled(def); ok {
		info, err := PluginInfo(path)
		if err != nil {
			return sdkplugin.PluginInfo{}, err
		}
		w.Write([]byte("Plugin already installed: " + info.Package))
		return PluginInfo(path)
	}

	w.Write([]byte("Building plugin from local path: " + def.LocalPath))

	info, err := PluginInfo(def.LocalPath)
	if err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	err = installPlugin(def.LocalPath, info, InstallOpts{RemoveSrc: false})
	if err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	if err := MarkPluginAsInstalled(def, PluginInstallPath(info)); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	return info, nil
}

func BuildFromGit(w io.Writer, def PluginSrcDef) (sdkplugin.PluginInfo, error) {
	ok, path := IsPluginInstalled(def)
	info, err := PluginInfo(path)
	if err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	if ok {
		w.Write([]byte("Plugin already installed: " + info.Package))
		return PluginInfo(path)
	}

	// TODO: update disk file path to randomly select either /etc /var /usr
	diskfileParentPath := filepath.Join(sdkpaths.TmpDir, "plugin-clone", "disk", info.Package)
	// ensure to create the virt disk parent file path exists
	fmt.Printf("creating virtual disk file parent path at: %s", diskfileParentPath)
	if err := os.MkdirAll(diskfileParentPath, 0755); err != nil {
		return sdkplugin.PluginInfo{}, err
	}
	diskfile := filepath.Join(diskfileParentPath, info.Package)

	clonePath := filepath.Join(sdkpaths.TmpDir, "plugin-clone", "mount", info.Package)
	dev := sdkstr.Slugify(info.Package, "_")
	mnt := encdisk.NewEncrypedDisk(clonePath, diskfile, dev)
	if err := mnt.Mount(); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	w.Write([]byte("Cloning plugin from git: " + def.GitURL))
	repo := git.RepoSource{URL: def.GitURL, Ref: def.GitRef}

	if err := git.Clone(w, repo, clonePath); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	if err := installPlugin(clonePath, info, InstallOpts{}); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	if err := mnt.Unmount(); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	installPath := PluginInstallPath(info)
	if err := MarkPluginAsInstalled(def, installPath); err != nil {
		return sdkplugin.PluginInfo{}, err
	}

	return PluginInfo(clonePath)
}

type GoBuildArgs struct {
	WorkDir   string
	Env       []string
	ExtraArgs []string
}

func BuildPlugin(pluginSrcDir string, workdir string) error {
	if pluginSrcDir == "" {
		return errors.New("Build plugin error: no plugin source path")
	}

	if workdir == "" {
		return errors.New("Build plugin error: no build path")
	}

	var info sdkplugin.PluginInfo

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
	defer os.RemoveAll(workdir)

	if err := sdkfs.EnsureDir(filepath.Join(workdir, "plugins")); err != nil {
		return err
	}

	if err := sdkfs.CopyDir(pluginSrcDir, buildpath, nil); err != nil {
		return err
	}

	if err := sdkfs.CopyDir(filepath.Join(sdkpaths.AppDir, "sdk"), filepath.Join(workdir, "sdk"), nil); err != nil {
		return err
	}

	goWork := fmt.Sprintf(`
go %s

use (
    ./sdk
    ./plugins/%s
)
    `, sdkruntime.GO_SHORT_VERSION, info.Package)

	if err := os.WriteFile(filepath.Join(workdir, "go.work"), []byte(goWork), sdkfs.PermFile); err != nil {
		return err
	}

	gofile := "main.go"
	outfile := "plugin.so"
	args := &GoBuildArgs{WorkDir: buildpath, ExtraArgs: []string{"-buildmode=plugin"}}
	if err := BuildGoModule(gofile, outfile, args); err != nil {
		return err
	}

	pluginSoOut := filepath.Join(buildpath, "plugin.so")
	fmt.Printf("Copying '%s' to '%s'\n", sdkpaths.StripRoot(pluginSoOut), sdkpaths.StripRoot(pluginSoPath))
	if err := sdkfs.CopyFile(pluginSoOut, pluginSoPath); err != nil {
		return err
	}

	return nil
}

func BuildGoModule(gofile string, outfile string, params *GoBuildArgs) error {
	if params == nil {
		params = &GoBuildArgs{}
	}

	fmt.Println("Building go module: " + sdkpaths.StripRoot(filepath.Join(params.WorkDir, gofile)))

	goBin := GoBin()
	buildArgs := BuildArgs()
	buildArgs = append(buildArgs, params.ExtraArgs...)

	buildCmd := []string{"build"}
	buildCmd = append(buildCmd, buildArgs...)
	buildCmd = append(buildCmd, "-o", outfile, gofile)

	fmt.Printf(`Build working directory: %s`+"\n", sdkpaths.StripRoot(params.WorkDir))
	cmd := exec.Command(goBin, buildCmd...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), params.Env...)
	cmd.Dir = params.WorkDir

	fmt.Printf("Executing: %s %s\n", goBin, strings.Join(buildCmd, " "))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error building go module " + sdkpaths.StripRoot(params.WorkDir) + ":" + err.Error())
		return err
	}

	fmt.Println("Module built successfully: " + sdkpaths.StripRoot(filepath.Join(params.WorkDir, outfile)))
	return nil
}
