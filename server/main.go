package main

import (
	"context"
	stdLog "log"

	"github.com/Noah-Huppert/gointerrupt"
	"github.com/Noah-Huppert/media-recyclarr/emby"
	"github.com/Noah-Huppert/media-recyclarr/jelly"
	"go.uber.org/zap"
)

func main() {
	ctxPair := gointerrupt.NewCtxPair(context.Background())

	// Setup logger
	log, err := zap.NewDevelopment()
	if err != nil {
		stdLog.Fatalf("failed to setup logger: %s", err)
	}
	defer log.Sync()

	// Load configuration
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal("failed to load configuration", zap.Error(err))
	}

	jellyClient, err := jelly.NewJellyClient(jelly.NewJellyClientOpts{
		JellyseerrURL:    cfg.JellyseerrURL,
		JellyseerrAPIKey: cfg.JellyseerrAPIKey,
	})
	if err != nil {
		log.Fatal("failed to create jellyseerr client", zap.Error(err))
	}

	embyClient, err := emby.NewEmbyClient(emby.NewEmbyClientOpts{
		EmbyURL:    cfg.EmbyURL,
		EmbyAPIKey: cfg.EmbyAPIKey,
	})
	if err != nil {
		log.Fatal("failed to create emby client", zap.Error(err))
	}

	trasher := NewTrasher(NewTrasherOpts{
		EmbyClient:  embyClient,
		JellyClient: jellyClient,
	})

	log.Debug("use these vars", zap.Any("embyMgr", embyMgr), zap.Any("jellyMgr", jellyMgr), zap.Any("ctxPair", ctxPair))
}
