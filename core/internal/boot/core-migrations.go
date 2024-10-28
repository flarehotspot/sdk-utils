package boot

import (
	"log"
	"path/filepath"

	"core/internal/plugins"
	"core/internal/utils/migrate"

	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

func RunCoreMigrations(g *plugins.CoreGlobals) {
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
}
