package tasks

import (
	"context"
	"fmt"
	"github.com/Noah-Huppert/media-recyclarr/emby"
	"github.com/Noah-Huppert/media-recyclarr/jelly"
	"github.com/Noah-Huppert/media-recyclarr/trasher"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

// taskQueueSize is the number of tasks which can be in the queue
const taskQueueSize = 10

// submitTaskPayload records a request to run a task
type submitTaskPayload struct {
	// taskType is the type of task to run
	taskType TaskType
}

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

	// taskQueue receives messages when tasks should be run
	taskQueue chan submitTaskPayload
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
		taskQueue:   make(chan submitTaskPayload, taskQueueSize),
	}
}

// Run blocks and executes tasks as they are needed
func (mgr *TaskManager) Run(graceful context.Context, harsh context.Context) error {
	mgr.logger.Info("started")

	refreshTicker := time.NewTicker(time.Hour * 1)
	mgr.submitRefreshTasks()

	for {
		select {
		case <-graceful.Done():
			mgr.logger.Info("graceful shutdown complete")
			return nil

		case <-harsh.Done():
			mgr.logger.Info("harsh shutdown complete")
			return nil

		case req := <-mgr.taskQueue:
			mgr.logger.Debug("received request to execute task", zap.String("task_type", string(req.taskType)))
			if err := mgr.executeTask(req); err != nil {
				mgr.logger.Error("failed to execute task", zap.String("task_type", string(req.taskType)), zap.Error(err))
			} else {
				mgr.logger.Debug("successfully executed task", zap.String("task_type", string(req.taskType)))
			}

			break

		case <-refreshTicker.C:
			mgr.submitRefreshTasks()
		}
	}
}

// executeTask runs the logic of the requested task
func (mgr *TaskManager) executeTask(req submitTaskPayload) error {
	switch req.taskType {
	case TaskUpdateUsers:
		return nil

	default:
		return fmt.Errorf("unknown task type: %s", req.taskType)
	}

	return nil
}

// TaskType indicates the task logic to run
type TaskType string

const (
	// TaskUpdateUsers indicates the update users task
	TaskUpdateUsers TaskType = "update_users"
)

// Submit adds a request to execute a task to the queue
func (mgr *TaskManager) Submit(taskType TaskType) {
	mgr.taskQueue <- submitTaskPayload{
		taskType: taskType,
	}
}

// submitRefreshTasks submits tasks which are needed for a full refresh of data
func (mgr *TaskManager) submitRefreshTasks() {
	mgr.Submit(TaskUpdateUsers)
}
