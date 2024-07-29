package plugins

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"os/exec"
// 	"path/filepath"

// 	"core/internal/config/plugincfg"
// 	"sdk/utils/fs"
// 	"sdk/utils/paths"
// )

// type InstPrgrs struct {
// 	Done bool   `json:"done"`
// 	Msg  string `json:"msg"`
// 	Err  bool   `json:"err"`
// }

// func Build(w io.Writer, pluginSrc string, buildOpts ...string) error {

// 	pluginSrc, err := plugincfg.FindPluginSrc(pluginSrc)
// 	if err != nil {
// 		return err
// 	}

// 	pluginInfo, err := plugincfg.GetPluginInfo(pluginSrc)
// 	if err != nil {
// 		return err
// 	}

// 	w.Write([]byte(fmt.Sprintf("Preparing to build package '%s'...", pluginInfo.Package)))

// 	installPath := filepath.Join(sdkpaths.TmpDir, "build", pluginInfo.Package)
// 	err = os.RemoveAll(installPath)
// 	if err != nil {
// 		return err
// 	}

// 	err = sdkfs.MoveDir(pluginSrc, installPath)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("Done moving files: ", pluginSrc, installPath)

// 	gowork := fmt.Sprintf(`
// use (
//   %s
//  %s
// )

// go 1.19
//     `, filepath.Join(sdkpaths.AppDir, "core"), installPath)

// 	err = os.WriteFile(filepath.Join(installPath, "go.work"), []byte(gowork), 0644)
// 	if err != nil {
// 		return err
// 	}
// 	log.Println("done writing go.work")

// 	soPath := filepath.Join(installPath, "plugin.so")
// 	buildargs := []string{"build", "-buildmode=plugin", "-trimpath", "-ldflags", "-s -w"}
// 	buildargs = append(buildargs, buildOpts...)
// 	buildargs = append(buildargs, "-o", soPath)

// 	buildCmd := exec.Command("go", buildargs...)
// 	buildCmd.Dir = installPath

// 	log.Println("Building plugin " + pluginInfo.Name)
// 	log.Println("go + ", buildargs)

// 	w.Write([]byte(fmt.Sprintf("Building plugin %s...", pluginInfo.Name)))
// 	err = buildCmd.Run()
// 	if err != nil {
// 		return err
// 	}

// 	pluginPath := filepath.Join(sdkpaths.PluginsDir, pluginInfo.Package)
// 	os.RemoveAll(filepath.Join(pluginPath, ".git"))
// 	log.Println("Moving plugin files to", pluginPath)

// 	w.Write([]byte(fmt.Sprintf("Cleaning up plugin %s...", pluginInfo.Name)))

// 	err = sdkfs.MoveDir(installPath, pluginPath)
// 	if err != nil {
// 		return err
// 	}

// 	patterns := []string{"*.go", "*.mod", "*.work", "*.md"}
// 	for _, pattern := range patterns {
// 		log.Println("Removing pattern", pattern)
// 		sdkfs.RmPattern(pluginPath, pattern)
// 	}

// 	err = sdkfs.RmEmpty(pluginPath)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
