package boot

import (
	"log"
	"path/filepath"

	"github.com/flarehotspot/core/internal/plugins"
	paths "github.com/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/core/internal/utils/migrate"
)

func RunMigrations(g *plugins.CoreGlobals) {
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
