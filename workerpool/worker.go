package workerpool

import (
	"context"
	"sync"
)

// Worker represents a worker in the worker pool.
type Worker struct {
	workerCount int
	taskQueue   chan Task
	wg          *sync.WaitGroup
	errors      chan error
	ctx         context.Context
	cancel      context.CancelFunc
}

// NewWorker creates a new Worker with the specified worker count.
func NewWorker(workerCount int) *Worker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Worker{
		workerCount: workerCount,
		taskQueue:   make(chan Task),
		wg:          &sync.WaitGroup{},
		errors:      make(chan error, 100), // buffered for error collection
		ctx:         ctx,
		cancel:      cancel,
	}
}

// Submit adds a task to the worker pool.
func (w *Worker) Submit(task Task) {
	w.wg.Add(1)
	w.taskQueue <- task
}

// worker is a goroutine that processes tasks from the task queue.
func (w *Worker) worker() {
	for {
		select {
		case <-w.ctx.Done():
			return
		case task, ok := <-w.taskQueue:
			if !ok {
				return
			}
			if err := task.Execute(w.ctx); err != nil {
				w.errors <- err
			}
			w.wg.Done()
		}
	}
}

// Run starts the worker pool.
func (w *Worker) Run() {
	for i := 0; i < w.workerCount; i++ {
		go w.worker()
	}
}

// Wait waits for all tasks to complete and closes the error channel.
func (w *Worker) Wait() []error {
	w.wg.Wait()
	w.cancel()
	close(w.errors)
	var errs []error
	for err := range w.errors {
		errs = append(errs, err)
	}
	return errs
}

// Shutdown gracefully shuts down the worker pool.
func (w *Worker) Shutdown() {
	w.cancel()
	close(w.taskQueue)
}
