package main

import (
	"context"
	"github.com/Noah-Huppert/media-recyclarr/httpapi"
	"github.com/Noah-Huppert/media-recyclarr/models"
	"github.com/Noah-Huppert/media-recyclarr/trasher"
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

	// Connect to database
	db, err := models.OpenDB(models.OpenDBOpts{
		PostgresURI: cfg.PostgresURI,
	})
	if err != nil {
		log.Fatal("failed to open database connection", zap.Error(err))
	}
	log.Debug("database connected", zap.Any("db", db))

	// Setup Jellyseerr client
	jellyClient, err := jelly.NewJellyClient(jelly.NewJellyClientOpts{
		Logger:           log.Named("jellyseerr"),
		JellyseerrURL:    cfg.JellyseerrURL,
		JellyseerrAPIKey: cfg.JellyseerrAPIKey,
	})
	if err != nil {
		log.Fatal("failed to create jellyseerr client", zap.Error(err))
	}

	// Setup Emby
	embyClient, err := emby.NewEmbyClient(emby.NewEmbyClientOpts{
		Logger:     log.Named("emby"),
		EmbyURL:    cfg.EmbyURL,
		EmbyAPIKey: cfg.EmbyAPIKey,
	})
	if err != nil {
		log.Fatal("failed to create emby client", zap.Error(err))
	}

	// Setup Trasher
	trsh := trasher.NewTrasher(trasher.NewTrasherOpts{
		Logger:      log.Named("trasher"),
		EmbyClient:  embyClient,
		JellyClient: jellyClient,
		ExpireAfter: time.Hour * 24 * 30, // 30 days
	})

	// Setup and run HTTP API
	apiSrv := httpapi.NewHTTPAPI(httpapi.NewHTTPAPIOpts{
		Logger:  log.Named("http-api"),
		Address: cfg.HTTPAPIAddress,
		Trasher: trsh,
	})

	if err := apiSrv.Run(ctxPair.Graceful(), ctxPair.Harsh()); err != nil {
		log.Fatal("failed to run HTTP API server", zap.Error(err))
	}
}
