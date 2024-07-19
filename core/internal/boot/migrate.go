package boot

import (
	"log"
	"path/filepath"

	"core/internal/plugins"
	"core/internal/utils/migrate"
	sdkpaths "sdk/utils/paths"
)

func RunMigrations(g *plugins.CoreGlobals) {
	db := g.Db.SqlDB()

	err := migrate.Init(db)
	if err != nil {
		log.Println(err)
		return
	}

	err = migrate.MigrateUp(db, filepath.Join(sdkpaths.CoreDir, "resources/migrations"))
	if err != nil {
		log.Printf("Core migrations error: %s", err.Error())
	} else {
		log.Println("Core migrations success!")
	}

	// pluginDirs := config.PluginDirList()
	// for _, pdir := range pluginDirs {
	// 	migdir := filepath.Join(pdir, "resources/migrations")
	// 	err := migrate.MigrateUp(g.Db.SqlDB(), migdir)
	// 	if err != nil {
	// 		log.Println("Error in plugin migration "+pdir, ":", err.Error())
	// 	} else {
	// 		log.Println("Done migrating plugin:", pdir)
	// 	}
	// }
}
