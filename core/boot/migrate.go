package boot

import (
	"log"
	"path/filepath"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/utils/migrate"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

func RunMigrations(g *globals.CoreGlobals) {
	db := g.Db.SqlDB()

	err := migrate.Init(db)
	if err != nil {
		log.Println(err)
		return
	}

	err = migrate.MigrateUp(filepath.Join(paths.CoreDir, "resources/migrations"), db)
	if err != nil {
		log.Printf("Core migrations error: %s", err.Error())
	} else {
		log.Println("Core migrations success!")
	}

	g.PluginMgr.MigrateAll()
}
