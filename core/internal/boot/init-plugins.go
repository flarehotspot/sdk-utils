package boot

import (
	"log"

	"core/internal/config"
	"core/internal/plugins"
)

func InitPlugins(g *plugins.CoreGlobals) error {
	bp := g.BootProgress

	// out := plugins.InstallPlugins()
	// done := false

	// for !done {
	// 	select {
	// 	case msg := <-out.Msg:
	// 		g.BootProgress.SetStatus(msg)
	// 	case err := <-out.Done:
	// 		done = true

	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	bp.SetStatus("Initializing database...")
	RunMigrations(g)

	pluginDirs := config.PluginDirList()
	log.Println("Plugin dirs:", pluginDirs)

	for _, dir := range pluginDirs {
		log.Println("Loading plugin from :", dir)
		p := plugins.NewPluginApi(dir, g.PluginMgr, g.TrafficMgr)
		g.PluginMgr.RegisterPlugin(p)
	}

	return nil
}
