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
		info, err := pkg.InstallSrcDef(inst, def)
		if err != nil {
			msg := fmt.Sprintf("Error installing plugin %s: %s", def.String(), err.Error())
			bp.AppendLog(msg)
		} else {
			msg := "Installed plugin: " + info.Package
			bp.AppendLog(msg)
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
