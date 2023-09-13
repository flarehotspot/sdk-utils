package boot

import (
	"fmt"
	"log"
	"time"

	"github.com/flarehotspot/core/globals"
)

func Init(g *globals.CoreGlobals) {
	bp := g.BootProgress
	now := time.Now()

	InitDirs()

	go func() {
		bp.SetStatus("Initializing plugins...")
		err := InitPlugins(g)
		if err != nil {
			bp.SetDone(err)
			return
		}

		// delay boot
		time.Sleep(1000 * 3 * time.Millisecond)

		bp.SetStatus("Initializing storage...")
		InitStorage()

		// delay boot
		time.Sleep(1000 * 3 * time.Millisecond)

		bp.SetStatus("Initializing admin accounts...")
		InitAccounts()

		// delay boot
		time.Sleep(1000 * 3 * time.Millisecond)

		bp.SetStatus("Setting up network interfaces...")
		InitNetwork()

		// delay boot
		time.Sleep(1000 * 3 * time.Millisecond)

		s := fmt.Sprintf("Done booting in %v", time.Since(now))
		bp.SetStatus(s)

		time.Sleep(1000 * 1 * time.Millisecond)
		bp.SetDone(nil)

		log.Println("Done booting in", time.Since(now))
	}()

	InitHttpServer(g)
}
