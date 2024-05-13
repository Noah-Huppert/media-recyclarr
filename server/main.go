package main

import (
	"context"

	"github.com/Noah-Huppert/gointerrupt"
	"github.com/Noah-Huppert/golog"
)

func main() {
	ctxPair := gointerrupt.NewCtxPair(context.Background())
	log := golog.NewLogger("main")

	// Load configuration
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %s", err)
	}

	embyMgr, err := NewEmbyManager(NewEmbyManagerOpts{
		EmbyURL:    cfg.EmbyURL,
		EmbyAPIKey: cfg.EmbyAPIKey,
	})
	if err != nil {
		log.Fatalf("failed to create emby manager: %s", err)
	}

	jellyMgr, err := NewJellyseerrManager(NewJellyseerrManagerOpts{
		JellyseerrURL:    cfg.JellyseerrURL,
		JellyseerrAPIKey: cfg.JellyseerrAPIKey,
	})
}
