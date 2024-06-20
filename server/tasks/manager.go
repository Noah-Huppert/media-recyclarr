package tasks

import (
	"context"
	"github.com/Noah-Huppert/media-recyclarr/emby"
	"github.com/Noah-Huppert/media-recyclarr/jelly"
	"github.com/Noah-Huppert/media-recyclarr/trasher"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// TaskManager is responsible for running background tasks
type TaskManager struct {
	// logger is used to output runtime information
	logger *zap.Logger

	// db is the database client
	db *gorm.DB

	// jellyClient is a Jellyseerr client
	jellyClient *jelly.JellyClient

	// embyClient is a Emby client
	embyClient *emby.EmbyClient

	// trsh is the Trasher
	trsh *trasher.Trasher
}

// NewTaskManagerOpts are options for NewTaskManager()
type NewTaskManagerOpts struct {
	// Logger is used to output runtime information
	Logger *zap.Logger

	// DB is the database client
	DB *gorm.DB

	// JellyClient is a Jellyseerr client
	JellyClient *jelly.JellyClient

	// EmbyClient is a Emby client
	EmbyClient *emby.EmbyClient

	// Trsh is the Trasher
	Trsh *trasher.Trasher
}

// NewTaskManager creates a new TaskManager
func NewTaskManager(opts NewTaskManagerOpts) *TaskManager {
	return &TaskManager{
		logger:      opts.Logger,
		db:          opts.DB,
		jellyClient: opts.JellyClient,
		embyClient:  opts.EmbyClient,
		trsh:        opts.Trsh,
	}
}

func (mgr *TaskManager) Run(graceful context.Context, harsh context.Context) error {
	mgr.logger.Info("started")
	<-graceful.Done()
	mgr.logger.Info("done")
	return nil
}
