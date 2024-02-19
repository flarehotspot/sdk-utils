package plugins

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/flarehotspot/core/internal/config"
	"github.com/flarehotspot/core/internal/config/plugincfg"
	"github.com/flarehotspot/core/internal/utils/git"
	paths "github.com/flarehotspot/core/sdk/utils/paths"
	strings "github.com/flarehotspot/core/sdk/utils/strings"
    sdkplugin "github.com/flarehotspot/core/sdk/api/plugin"
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

func InstallPlugins() *InstallStatus {
	out := &InstallStatus{
		Msg:  make(chan string),
		Done: make(chan error),
	}

	go func() {
		for _, def := range config.AllPluginSrc() {
			if !isInstalled(def) {
				if def.Src == config.PluginSrcGit {
					info, err := buildFromGit(out, def)
					if err != nil {
						log.Println("buildFromGit error:", err)
						out.Done <- err
						return
					}

					err = plugincfg.WriteCache(def, info)
					if err != nil {
						log.Println("WriteCache error:", err)
						out.Done <- err
						return
					}
				}

				if def.Src == config.PluginSrcStore {
					log.Printf("TODO: build from store")
				}

			}
		}

		out.Done <- nil
	}()

	return out
}

func buildFromGit(w io.Writer, src *config.PluginSrcDef) (*sdkplugin.PluginInfo, error) {
	repo := git.RepoSource{URL: src.GitURL, Ref: src.GitRef}
	clonePath := filepath.Join(paths.TmpDir, "plugins", strings.Rand(16))

	err := git.Clone(w, repo, clonePath)
	if err != nil {
		return nil, err
	}

	info, err := plugincfg.GetPluginInfo(clonePath)
	if err != nil {
		return nil, err
	}

	// if ok := UserLocalVersion(w, info.Package); ok {
	// return plugincfg.GetInstallInfo(info.Package)
	// }

	err = Build(w, clonePath)
	if err != nil {
		return nil, err
	}

	os.RemoveAll(clonePath)
	return plugincfg.GetInstallInfo(info.Package)
}

func isInstalled(def *config.PluginSrcDef) bool {
	cacheInfo, ok := plugincfg.GetCacheInfo(def)
	if !ok {
		return false
	}

	installInfo, err := plugincfg.GetInstallInfo(cacheInfo.Package)
	if err != nil {
		return false
	}

	return installInfo.Package == cacheInfo.Package
}
