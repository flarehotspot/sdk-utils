package boot

import (
	"log"

	"github.com/flarehotspot/flarehotspot/core/utils/uci"
	"github.com/flarehotspot/flarehotspot/core/utils/storage"
)

// InitStorage initializes the storage
func InitStorage() {
	info, ok := uci.GetFlareStorage()
	if !ok {
		log.Println("failed to get storage info from uci")
		return
	}

	if info.Expanded {
		log.Println("already expanded storage")
		return
	}

	err := storage.Expand(info.Storage, info.Partition, info.Partnum)
	if err != nil {
		log.Println("failed to expand storage:", err)
	}

	err = uci.SetFlareStorageExpanded(true)
	if err != nil {
		log.Println("failed to set storage expanded:", err)
	}

	err = uci.UciTree.Commit()
	if err != nil {
		log.Println("failed to commit uci:", err)
	}
}
