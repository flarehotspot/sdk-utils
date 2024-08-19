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
		// // Skip when plugin is already installed and no recompilation is needed
		// installed := pkg.IsPluginInstalled(def)
		// recompile := pkg.NeedsRecompile(def)
		// info, err := pkg.GetPluginInfo(def)

		// if installed && err == nil {
		// 	// Install pending update
		// 	if pkg.HasPendingUpdate(info.Package) {
		// 		bp.AppendLog(fmt.Sprintf("Plugin %s has pending update, moving to installed folder.", info.Package))
		// 		if err := pkg.MovePendingUpdate(info.Package); err != nil {
		// 			// Restore previous version if error occurs
		// 			bp.AppendLog(fmt.Sprintf("Error moving plugin %s to installed folder, restoring previously installed version.", info.Package))
		// 			if err := pkg.RestoreBackup(info.Package); err != nil {
		// 				log.Println("RestoreBackup error:", err)
		// 				msg := fmt.Sprintf("Error restoring previous version of %s: %s", info.Package, err.Error())
		// 				bp.AppendLog(msg)
		// 				continue
		// 			}
		// 		}
		// 	}
		// }

		// if !recompile {
		// 	msg := fmt.Sprintf("Plugin %s is already installed", info.Package)
		// 	bp.AppendLog(msg)
		// 	continue
		// }

		// if installed {
		// 	if err := pkg.CreateBackup(info.Package); err != nil {
		// 		bp.AppendLog(fmt.Sprintf("Error creating backup for plugin %s: %s", info.Package, err.Error()))
		// 		continue
		// 	}
		// 	if err := os.RemoveAll(pkg.GetInstallPath(info.Package)); err != nil {
		// 		bp.AppendLog(fmt.Sprintf("Error removing plugin %s: %s", info.Package, err.Error()))
		// 		continue
		// 	}
		// }

		info, err := pkg.InstallSrcDef(inst, def)
		if err != nil {
			msg := fmt.Sprintf("Error installing plugin %s: %s", def.String(), err.Error())
			bp.AppendLog(msg)

			if pkg.HasBackup(info.Package) {
				if err := pkg.RestoreBackup(info.Package); err != nil {
					log.Println("RestoreBackup error:", err)
					msg := fmt.Sprintf("Error restoring previous version of %s: %s", info.Package, err.Error())
					bp.AppendLog(msg)
				}
			}
		} else {
			msg := "Installed plugin: " + info.Package
			bp.AppendLog(msg)

			if pkg.HasBackup(info.Package) {
				if err := pkg.RemoveBackup(info.Package); err != nil {
					log.Println("RemoveBackup error:", err)
					msg := fmt.Sprintf("Error removing backup for %s: %s", info.Package, err.Error())
					bp.AppendLog(msg)
				}
			}
		}

		time.Sleep(1000 * 3 * time.Millisecond)
	}

	// Load plugins
	pluginDirs := pkg.InstalledDirList()
	log.Println("Installed plugin directories:", pluginDirs)
	for _, dir := range pluginDirs {
		log.Println("Loading plugin from :", dir)
		p := plugins.NewPluginApi(dir, g.PluginMgr, g.TrafficMgr)
		g.PluginMgr.RegisterPlugin(p)
	}
}
