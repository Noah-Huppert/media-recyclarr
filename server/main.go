package main

import (
	"context"
	stdLog "log"
	"strings"
	"time"

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
	defer func() {
		if err := log.Sync(); err != nil {
			if strings.Contains(err.Error(), "invalid argument") {
				// For some reason Sync() returns "invalid argument" when it successfully runs
				return
			}
			stdLog.Fatalf("failed to sync logs: %s", err)
		}
	}()

	// Load configuration
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal("failed to load configuration", zap.Error(err))
	}

	jellyClient, err := jelly.NewJellyClient(jelly.NewJellyClientOpts{
		Logger:           log.Named("jellyseerr"),
		JellyseerrURL:    cfg.JellyseerrURL,
		JellyseerrAPIKey: cfg.JellyseerrAPIKey,
	})
	if err != nil {
		log.Fatal("failed to create jellyseerr client", zap.Error(err))
	}

	embyClient, err := emby.NewEmbyClient(emby.NewEmbyClientOpts{
		Logger:     log.Named("emby"),
		EmbyURL:    cfg.EmbyURL,
		EmbyAPIKey: cfg.EmbyAPIKey,
	})
	if err != nil {
		log.Fatal("failed to create emby client", zap.Error(err))
	}

	trasher := NewTrasher(NewTrasherOpts{
		Logger:      log.Named("trasher"),
		EmbyClient:  embyClient,
		JellyClient: jellyClient,
		ExpireAfter: time.Hour * 24 * 30, // 30 days
	})

	reqMedia, err := trasher.GetRequestedMedia(ctxPair.Graceful())
	if err != nil {
		log.Fatal("failed to get requested media", zap.Error(err))
	}

	expiredMedia, err := trasher.GetExpiredMedia(ctxPair.Graceful(), reqMedia)
	if err != nil {
		log.Fatal("failed to get expired media", zap.Error(err))
	}

	for _, line := range expiredMedia.FormatTree(0) {
		log.Debug(line)
	}
}
