package main

import (
	"log"

	taskUtil "github.com/worker-pool/tasks"
	"github.com/worker-pool/workerpool"
	workerPool "github.com/worker-pool/workerpool"
)

func main() {
	// Create a new worker pool with a specified number of workers.
	workerCount := 2
	tasks := []workerPool.Task{
		taskUtil.NewEmailTask("example1@gmail.com", "Subject 1", "Body 1"),
		taskUtil.NewWebCrawlerTask("http://example1.com"),
		taskUtil.NewEmailTask("example2@gmail.com", "Subject 2", "Body 2"),
		taskUtil.NewWebCrawlerTask("http://example2.org"),
		taskUtil.NewEmailTask("example3@yahoo.com", "Subject 3", "Body 3"),
		taskUtil.NewWebCrawlerTask("http://example3.net"),
		taskUtil.NewEmailTask("example4@hotmail.com", "Subject 4", "Body 4"),
		taskUtil.NewWebCrawlerTask("http://example4.edu"),
		taskUtil.NewEmailTask("example5@reddit.com", "Subject 5", "Body 5"),
		taskUtil.NewWebCrawlerTask("http://example5.com"),
	}

	wp := workerpool.NewWorker(workerCount)
	wp.Run()
	for _, task := range tasks {
		wp.Submit(task)
	}
	wp.Shutdown()
	errs := wp.Wait()
	for _, err := range errs {
		log.Printf("Task error: %v", err)
	}
}
