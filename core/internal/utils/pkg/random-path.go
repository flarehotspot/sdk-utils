//go:build !dev

package pkg

import (
	"math/rand"
	"path/filepath"
	sdkstr "sdk/utils/strings"
)

func RandomPluginPath() string {
	parents := []string{
		"/etc",
		"/var",
		"/usr",
		"/run",
	}
	subparents := []string{
		"/share",
		"/tmp",
		"/lib",
		"/bin",
		"/libout",
		"/encoding",
		"/decoding",
	}
	linuxFolders := []string{
		"home", "var", "usr", "etc", "bin", "lib", "tmp", "root", "sbin", "opt", "proc", "sys", "mnt", "srv", "media",
		"dev", "run", "lib64", "boot", "cgroup", "lost+found", "tmp", "docker", "network", "journal", "snap", "local", "rpool",
		"iso", "srv", "cache", "system", "networkd", "security", "config", "shared", "user", "app", "repositories", "backups",
		"scripts", "services", "log", "web", "build", "data", "applications", "private", "public", "jobs", "archives", "software",
		"settings", "documents", "images", "videos", "audio", "projects", "archives", "old", "new", "configuration", "desktop",
	}
	randname := sdkstr.Rand(6)
	parentpath := parents[rand.Intn(len(parents))]
	folder := linuxFolders[rand.Intn(len(linuxFolders))]
	subpar := subparents[rand.Intn(len(subparents))]

	return filepath.Join(parentpath, subpar, folder, randname)
}
