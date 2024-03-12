package boot

import (
	"log"
	"path/filepath"

	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/utils/migrate"
	"github.com/flarehotspot/sdk/utils/paths"
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

	g.PluginMgr.MigrateAll()
}
