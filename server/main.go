package main

import (
	"context"
	stdLog "log"

	"github.com/Noah-Huppert/gointerrupt"
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

	embyMgr, err := NewEmbyManager(NewEmbyManagerOpts{
		EmbyURL:    cfg.EmbyURL,
		EmbyAPIKey: cfg.EmbyAPIKey,
	})
	if err != nil {
		log.Fatal("failed to create emby manager", zap.Error(err))
	}

	jellyMgr, err := NewJellyseerrManager(NewJellyseerrManagerOpts{
		JellyseerrURL:    cfg.JellyseerrURL,
		JellyseerrAPIKey: cfg.JellyseerrAPIKey,
	})
	if err != nil {
		log.Fatal("failed to create jellyseerr manager", zap.Error(err))
	}

	log.Debug("use these vars", zap.Any("embyMgr", embyMgr), zap.Any("jellyMgr", jellyMgr), zap.Any("ctxPair", ctxPair))
}
