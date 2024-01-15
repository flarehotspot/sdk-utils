package boot

import (
	"log"

	"github.com/flarehotspot/core/config/plugincfg"
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/plugins"
)

func InitPlugins(g *globals.CoreGlobals) error {
	bp := g.BootProgress

	out := plugins.InstallPlugins()
	done := false

	for !done {
		select {
		case msg := <-out.Msg:
			g.BootProgress.SetStatus(msg)
		case err := <-out.Done:
			done = true

			if err != nil {
				return err
			}
		}
	}

	bp.SetStatus("Initializing database...")
	RunMigrations(g)

	pluginDirs := plugincfg.ListDirs()
	log.Println("pluginDirs:", pluginDirs)

	for _, dir := range pluginDirs {
		log.Println("loading plugin from :", dir)
		p := plugins.NewPluginApi(dir, g.PluginMgr, g.TrafficMgr)
		g.PluginMgr.RegisterPlugin(p)
	}

	return nil
}
