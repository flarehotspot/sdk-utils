//go:build !mono

package boot

import (
	"log"

	"core/internal/config"
	"core/internal/plugins"
)

func InitPlugins(g *plugins.CoreGlobals) {
	bp := g.BootProgress
	install := plugins.InstallPlugins()
	done := false

	for !done {
		select {
		case msg := <-install.Msg:
			bp.AppendLog(msg)
		case err := <-install.Done:
			if err != nil {
				bp.Done(err)
				return
			}
			done = true
		}
	}

	pluginDirs := config.InstalledDirList()
	log.Println("Installed plugin directories:", pluginDirs)

	for _, dir := range pluginDirs {
		log.Println("Loading plugin from :", dir)
		p := plugins.NewPluginApi(dir, g.PluginMgr, g.TrafficMgr)
		g.PluginMgr.RegisterPlugin(p)
	}
}
