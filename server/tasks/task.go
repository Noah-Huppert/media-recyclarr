package tasks

import "context"

type Task interface {
	// Run executes task logic
	Run(ctx context.Context) error
}
