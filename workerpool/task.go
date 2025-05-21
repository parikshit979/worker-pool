package workerpool

import (
	"context"
)

// Task represents a unit of work to be processed by the worker pool.
// Each task must implement the Execute method, which contains the logic to be executed.
type Task interface {
	Execute(ctx context.Context) error
}
