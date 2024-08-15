package pkg

import (
	"core/internal/utils/cmd"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	sdkplugin "sdk/api/plugin"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
	sdkruntime "sdk/utils/runtime"
)

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
    `, sdkruntime.GO_VERSION, info.Package)

	goworkFile := filepath.Join(workdir, "go.work")
	if err := os.WriteFile(goworkFile, []byte(goWork), sdkfs.PermFile); err != nil {
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
		log.Printf("Error copying '%s' to '%s': %+v\n", pluginSoOut, pluginSoPath, err)
		return err
	}

	return nil
}

type GoBuildArgs struct {
	WorkDir   string
	Env       []string
	ExtraArgs []string
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

	cmdstr := goBin
	for _, arg := range buildCmd {
		cmdstr += " " + arg
	}

	// cmdfile := filepath.Join(params.WorkDir, sdkstr.Rand(16)+".sh")
	// if err := os.WriteFile(cmdfile, []byte(cmdstr), sdkfs.PermFile); err != nil {
	// 	return err
	// }

	fmt.Printf(`Build working directory: %s`+"\n", sdkpaths.StripRoot(params.WorkDir))
	// cmd := exec.Command(shell, cmdfile)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Env = append(os.Environ(), params.Env...)
	// cmd.Dir = params.WorkDir

	err := cmd.Exec(cmdstr, &cmd.ExecOpts{
		Stdout: os.Stdout,
		Env:    append(os.Environ(), params.Env...),
		Dir:    params.WorkDir,
	})
	if err != nil {
		return err
	}

	fmt.Println("Module built successfully: " + sdkpaths.StripRoot(filepath.Join(params.WorkDir, outfile)))
	return nil
}

type InstallOpts struct {
	RemoveSrc bool
}
