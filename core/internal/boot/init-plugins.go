//go:build !mono

package boot

import (
	"fmt"
	"log"
	"time"

	"core/internal/plugins"
	"core/internal/utils/pkg"
)

type InstallStatus struct {
	bp *plugins.BootProgress
}

func (is *InstallStatus) Write(p []byte) (n int, err error) {
	status := string(p)
	is.bp.AppendLog(status)
	return len(p), nil
}

func InitPlugins(g *plugins.CoreGlobals) {
	bp := g.BootProgress
	inst := &InstallStatus{bp: bp}

	for _, def := range pkg.AllPluginDef() {
		recompile := pkg.NeedsRecompile(def)

		// Skip when plugin is already installed and no recompilation is needed
		if !recompile {
			bp.AppendLog("Plugin is already installed: " + def.String())
			continue
		}

		switch def.Src {
		case pkg.PluginSrcLocal, pkg.PluginSrcSystem:
			info, err := pkg.BuildFromLocal(inst, def)
			if err != nil {
				bp.AppendLog(fmt.Sprintf("Error buidling plugin from local path %s: %s", def.LocalPath, err.Error()))
			} else {
				bp.AppendLog("Successfully installed plugin: " + info.Package)
			}

		default:
			log.Println("Unknown plugin source:", def.Src)
		}

		// if def.Src == pkg.PluginSrcGit {
		// 	info, err := pkg.BuildFromGit(inst, def)
		// 	if err != nil {
		// 		log.Println("buildFromGit error:", err)
		// 		msg = fmt.Sprintf("Error building plugin from git source %s: %s", def.GitURL, err.Error())
		// 	} else {
		// 		msg = "Installed plugin: " + info.Package
		// 	}
		// }

		// if def.Src == config.PluginSrcStore {
		// 	log.Printf("TODO: build from store")
		// }

		time.Sleep(1000 * 3 * time.Millisecond)
	}

	pluginDirs := pkg.InstalledDirList()
	log.Println("Installed plugin directories:", pluginDirs)

	for _, dir := range pluginDirs {
		log.Println("Loading plugin from :", dir)
		p := plugins.NewPluginApi(dir, g.PluginMgr, g.TrafficMgr)
		g.PluginMgr.RegisterPlugin(p)
	}
}
