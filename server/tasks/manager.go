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

	// taskSchedules indicate when tasks should repeat
	taskSchedules TaskSchedules

	// taskQueue receives messages when tasks should be run
	taskQueue chan submitTaskPayload
}

// TaskSchedules holds the schedule for when tasks should be run on a recurring schedule
type TaskSchedules map[TaskType]time.Duration

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

	// TaskSchedules indicate when tasks should repeat
	TaskSchedules TaskSchedules
}

// RequiredTaskSchedules are task types which must have a recurring schedule when creating a task manager
var RequiredTaskSchedules []TaskType = []TaskType{
	TaskUpdateUsers,
}

// NewTaskManager creates a new TaskManager
// Requires that task schedules be provided for: RequiredTaskSchedules
func NewTaskManager(opts NewTaskManagerOpts) (*TaskManager, error) {
	for _, taskType := range RequiredTaskSchedules {
		if _, ok := opts.TaskSchedules[taskType]; !ok {
			return nil, fmt.Errorf("TaskSchedules must contain a schedule for %s", taskType)
		}
	}
	return &TaskManager{
		logger:        opts.Logger,
		db:            opts.DB,
		jellyClient:   opts.JellyClient,
		embyClient:    opts.EmbyClient,
		trsh:          opts.Trsh,
		taskSchedules: opts.TaskSchedules,
		taskQueue:     make(chan submitTaskPayload, taskQueueSize),
	}, nil
}

// Run blocks and executes tasks as they are needed
func (mgr *TaskManager) Run(graceful context.Context, harsh context.Context) error {
	mgr.logger.Info("started")

	go mgr.recurringTaskTimers(graceful)

	for {
		select {
		case <-graceful.Done():
			mgr.logger.Info("graceful shutdown complete")
			return nil

		case <-harsh.Done():
			mgr.logger.Info("harsh shutdown complete")
			return nil

		case req := <-mgr.taskQueue:
			mgr.logger.Info("received request to execute task", zap.String("task_type", string(req.taskType)))
			if err := mgr.executeTask(graceful, req); err != nil {
				mgr.logger.Error("failed to execute task", zap.String("task_type", string(req.taskType)), zap.Error(err))
			} else {
				mgr.logger.Info("successfully executed task", zap.String("task_type", string(req.taskType)))
			}

			break
		}
	}
}

// executeTask runs the logic of the requested task
func (mgr *TaskManager) executeTask(ctx context.Context, req submitTaskPayload) error {
	// Create task
	var task Task

	switch req.taskType {
	case TaskUpdateUsers:
		task = UpdateUsersTask{
			logger:     mgr.logger.Named("update-users-task"),
			db:         mgr.db,
			embyClient: mgr.embyClient,
		}
		break

	default:
		return fmt.Errorf("unknown task type: %s", req.taskType)
	}

	// Run task
	return task.Run(ctx)
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

// recurringTaskTimers starts timers for each task type in taskSchedules and submits task execution requests on the recurring schedule.
// Also runs each task once at the start.
func (mgr *TaskManager) recurringTaskTimers(ctx context.Context) {
	// Make a ticker for each task type and have them forward messages to a shared channel
	tickChan := make(chan TaskType)

	for taskType, duration := range mgr.taskSchedules {
		go func() {
			ticker := time.NewTicker(duration)

			// Run task once at the start
			tickChan <- taskType

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					tickChan <- taskType
					break
				}
			}
		}()
	}

	// Submit tasks based on shared channel
	for {
		select {
		case <-ctx.Done():
			return
		case taskType := <-tickChan:
			mgr.Submit(taskType)
		}
	}
}
