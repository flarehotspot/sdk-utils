package plugincfg

// import (
// 	"log"
// 	"path/filepath"

// 	fs "github.com/flarehotspot/core/sdk/utils/fs"
// 	paths "github.com/flarehotspot/core/sdk/utils/paths"
// )

// // ListDirs returns a list of plugins (aboslute path to plugin directory) from "vendor" directory.
// // If same directory name exists in "plugins" directory, the absolute path from "plugins" directory is returned instead.
// func ListDirs() []string {
// 	vendorDirs := []string{}
// 	if err := fs.LsDirs(paths.VendorDir, &vendorDirs, false); err != nil {
// 		panic("Unable to list plugin directories.")
// 	}

// 	pluginDirs := []string{}
// 	if err := fs.LsDirs(paths.PluginsDir, &pluginDirs, false); err != nil {
// 		return vendorDirs
// 	}

// 	list := []string{}

// 	for _, vendorDir := range vendorDirs {
// 		var pluginDir *string

// 		for _, pdir := range pluginDirs {
// 			vname := filepath.Base(vendorDir)
// 			pname := filepath.Base(pdir)

// 			if pname == vname {
// 				pluginDir = &pdir
// 				break
// 			}
// 		}

// 		if pluginDir != nil {
// 			list = append(list, *pluginDir)
// 		} else {
// 			list = append(list, vendorDir)
// 		}
// 	}

// 	log.Println("Plugin List: ")
// 	for _, p := range list {
// 		log.Println("\t" + p)
// 	}

// 	return list
// }
