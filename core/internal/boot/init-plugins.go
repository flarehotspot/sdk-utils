//go:build !mono

package boot

import (
	"fmt"
	"log"

	"core/internal/plugins"
	"core/internal/utils/pkg"
)

type InstallStatus struct {
	Msg  chan string
	Done chan error
}

func (i *InstallStatus) Write(p []byte) (n int, err error) {
	status := string(p)
	i.Msg <- status
	return len(p), nil
}

func InitPlugins(g *plugins.CoreGlobals) {
	bp := g.BootProgress

	progress := &InstallStatus{
		Msg:  make(chan string),
		Done: make(chan error),
	}

	go func() {
		for _, def := range pkg.AllPluginSrc() {
			if def.Src == pkg.PluginSrcGit {
				info, err := pkg.BuildFromGit(progress, def)
				if err != nil {
					log.Println("buildFromGit error:", err)
					progress.Msg <- fmt.Sprintf("Error building plugin from git source %s: %s", def.GitURL, err.Error())
				} else {
					progress.Msg <- "Installed plugin: " + info.Package
				}
			}

			// if def.Src == config.PluginSrcStore {
			// 	log.Printf("TODO: build from store")
			// }

			if def.Src == pkg.PluginSrcSystem || def.Src == pkg.PluginSrcLocal {
				info, err := pkg.BuildFromLocal(progress, def)
				if err != nil {
					progress.Msg <- fmt.Sprintf("Error buidling plugin from local path %s: %s", def.LocalPath, err.Error())
				} else {
					progress.Msg <- "Installed plugin: " + info.Package
				}
			}

		}

		progress.Done <- nil
	}()

	done := false

	for !done {
		select {
		case msg := <-progress.Msg:
			bp.AppendLog(msg)
		case err := <-progress.Done:
			if err != nil {
				bp.Done(err)
				return
			}
			done = true
		}
	}

	pluginDirs := pkg.InstalledDirList()
	log.Println("Installed plugin directories:", pluginDirs)

	for _, dir := range pluginDirs {
		log.Println("Loading plugin from :", dir)
		p := plugins.NewPluginApi(dir, g.PluginMgr, g.TrafficMgr)
		g.PluginMgr.RegisterPlugin(p)
	}
}
