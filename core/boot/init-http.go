package boot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/flarehotspot/core/plugins"
	"github.com/flarehotspot/core/web"
	"github.com/flarehotspot/core/web/router"
)

func InitHttpServer(g *plugins.CoreGlobals) {
	web.SetupBootRoutes(g)
	server := web.StartServer(router.BootingRouter, false)

	err := <-g.BootProgress.DONE_C
	if err != nil {
		log.Println("Error while booting:", err)
		// TODO: Show recovery page
		return
	}

	log.Println("Boot progress done. Need to restart server...")

	// Gracefully shut down the server to clear booting routes
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		fmt.Printf("Server shutdown error: %v\n", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}

	// Restart the server with all routes
	web.SetupAllRoutes(g)
	log.Println("Starting server...")
	web.StartServer(router.RootRouter, true)
}
